-- Complete Database Schema
-- This file contains the complete schema for reference
-- Run individual migration files for proper versioning

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- Users Table
-- ============================================
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_is_active ON users(is_active);

-- ============================================
-- Refresh Tokens Table
-- ============================================
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(500) NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP WITH TIME ZONE,
    is_revoked BOOLEAN DEFAULT FALSE,
    device_info VARCHAR(255),
    ip_address INET,
    CONSTRAINT fk_refresh_tokens_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
CREATE INDEX idx_refresh_tokens_is_revoked ON refresh_tokens(is_revoked);

-- ============================================
-- Password Resets Table
-- ============================================
CREATE TABLE IF NOT EXISTS password_resets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(500) NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP WITH TIME ZONE,
    is_used BOOLEAN DEFAULT FALSE,
    ip_address INET,
    CONSTRAINT fk_password_resets_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_password_resets_user_id ON password_resets(user_id);
CREATE INDEX idx_password_resets_token ON password_resets(token);
CREATE INDEX idx_password_resets_expires_at ON password_resets(expires_at);
CREATE INDEX idx_password_resets_is_used ON password_resets(is_used);

-- ============================================
-- Audit Logs Table
-- ============================================
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    action VARCHAR(100) NOT NULL,
    entity_type VARCHAR(100),
    entity_id UUID,
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_audit_logs_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- ============================================
-- Functions
-- ============================================

-- ============================================
-- Emitens Table
-- ============================================
CREATE TABLE IF NOT EXISTS emitens (
    id BIGSERIAL PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    sector VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_emitens_symbol ON emitens(symbol);
CREATE INDEX idx_emitens_sector ON emitens(sector);
CREATE INDEX idx_emitens_name ON emitens(name);

-- ============================================
-- Stock Price Ticks Table
-- ============================================
CREATE TABLE IF NOT EXISTS stock_price_ticks (
    id BIGSERIAL PRIMARY KEY,
    emiten_id BIGINT NOT NULL,
    price_time TIMESTAMP NOT NULL,
    price NUMERIC(12,2) NOT NULL,
    volume BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_tick UNIQUE (emiten_id, price_time),
    CONSTRAINT fk_stock_price_ticks_emiten_id
        FOREIGN KEY (emiten_id)
        REFERENCES emitens(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_stock_price_ticks_emiten_id ON stock_price_ticks(emiten_id);
CREATE INDEX idx_stock_price_ticks_price_time ON stock_price_ticks(price_time);
CREATE INDEX idx_stock_price_ticks_emiten_time ON stock_price_ticks(emiten_id, price_time DESC);
CREATE INDEX idx_stock_price_ticks_price ON stock_price_ticks(price);
CREATE INDEX idx_stock_price_ticks_volume ON stock_price_ticks(volume);

-- ============================================
-- Stock Price Summary Table
-- ============================================
CREATE TABLE IF NOT EXISTS stock_price_summary (
    id BIGSERIAL PRIMARY KEY,
    emiten_id BIGINT NOT NULL,
    date DATE NOT NULL,
    open_price NUMERIC(12,2) NOT NULL,
    high_price NUMERIC(12,2) NOT NULL,
    low_price NUMERIC(12,2) NOT NULL,
    close_price NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_summary UNIQUE (emiten_id, date),
    CONSTRAINT fk_stock_price_summary_emiten_id
        FOREIGN KEY (emiten_id)
        REFERENCES emitens(id)
        ON DELETE CASCADE,
    CONSTRAINT check_price_order CHECK (
        high_price >= low_price AND
        high_price >= open_price AND
        high_price >= close_price AND
        low_price <= open_price AND
        low_price <= close_price
    )
);

CREATE INDEX idx_stock_price_summary_emiten_id ON stock_price_summary(emiten_id);
CREATE INDEX idx_stock_price_summary_date ON stock_price_summary(date);
CREATE INDEX idx_stock_price_summary_emiten_date ON stock_price_summary(emiten_id, date DESC);
CREATE INDEX idx_stock_price_summary_close_price ON stock_price_summary(close_price);

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to generate daily stock price summary from tick data
CREATE OR REPLACE FUNCTION generate_stock_summary(p_emiten_id BIGINT, p_date DATE)
RETURNS VOID AS $$
DECLARE
    v_open_price NUMERIC(12,2);
    v_high_price NUMERIC(12,2);
    v_low_price NUMERIC(12,2);
    v_close_price NUMERIC(12,2);
BEGIN
    -- Get OHLC from tick data
    SELECT
        FIRST_VALUE(price) OVER (ORDER BY price_time ASC) AS open_price,
        MAX(price) AS high_price,
        MIN(price) AS low_price,
        FIRST_VALUE(price) OVER (ORDER BY price_time DESC) AS close_price
    INTO v_open_price, v_high_price, v_low_price, v_close_price
    FROM stock_price_ticks
    WHERE emiten_id = p_emiten_id
      AND DATE(price_time) = p_date;

    -- Insert or update summary
    INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
    VALUES (p_emiten_id, p_date, v_open_price, v_high_price, v_low_price, v_close_price)
    ON CONFLICT (emiten_id, date) DO UPDATE SET
        open_price = EXCLUDED.open_price,
        high_price = EXCLUDED.high_price,
        low_price = EXCLUDED.low_price,
        close_price = EXCLUDED.close_price,
        updated_at = CURRENT_TIMESTAMP;
END;
$$ LANGUAGE plpgsql;

-- Function to generate summary for all emitens for a specific date
CREATE OR REPLACE FUNCTION generate_all_summaries(p_date DATE DEFAULT CURRENT_DATE)
RETURNS VOID AS $$
BEGIN
    -- Generate summary for each emiten that has tick data on this date
    INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
    SELECT
        emiten_id,
        p_date,
        FIRST_VALUE(price) OVER (PARTITION BY emiten_id ORDER BY price_time ASC) AS open_price,
        MAX(price) AS high_price,
        MIN(price) AS low_price,
        FIRST_VALUE(price) OVER (PARTITION BY emiten_id ORDER BY price_time DESC) AS close_price
    FROM stock_price_ticks
    WHERE DATE(price_time) = p_date
    GROUP BY emiten_id
    ON CONFLICT (emiten_id, date) DO UPDATE SET
        open_price = EXCLUDED.open_price,
        high_price = EXCLUDED.high_price,
        low_price = EXCLUDED.low_price,
        close_price = EXCLUDED.close_price,
        updated_at = CURRENT_TIMESTAMP;
END;
$$ LANGUAGE plpgsql;

-- Trigger to automatically update updated_at
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_emitens_updated_at
    BEFORE UPDATE ON emitens
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_stock_price_summary_updated_at
    BEFORE UPDATE ON stock_price_summary
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- ============================================
-- Views
-- ============================================

-- View for active users
CREATE OR REPLACE VIEW active_users AS
SELECT
    id,
    email,
    full_name,
    created_at,
    updated_at
FROM users
WHERE is_active = TRUE;

-- View for user sessions (valid refresh tokens)
CREATE OR REPLACE VIEW user_sessions AS
SELECT
    rt.id,
    rt.user_id,
    rt.device_info,
    rt.ip_address,
    rt.created_at AS session_start,
    rt.expires_at
FROM refresh_tokens rt
WHERE rt.is_revoked = FALSE
  AND rt.expires_at > CURRENT_TIMESTAMP;

-- ============================================
-- Stock Data Views
-- ============================================

-- View: Latest stock prices
CREATE OR REPLACE VIEW v_latest_stock_prices AS
SELECT DISTINCT ON (spt.emiten_id)
    e.id AS emiten_id,
    e.symbol,
    e.name,
    e.sector,
    spt.price_time,
    spt.price,
    spt.volume,
    spt.created_at AS tick_created_at
FROM emitens e
LEFT JOIN stock_price_ticks spt ON e.id = spt.emiten_id
ORDER BY spt.emiten_id, spt.price_time DESC;

-- View: Stock price summary by date
CREATE OR REPLACE VIEW v_stock_price_summary AS
SELECT
    e.id AS emiten_id,
    e.symbol,
    e.name,
    DATE(spt.price_time) AS price_date,
    MIN(spt.price) AS low_price,
    MAX(spt.price) AS high_price,
    FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time ASC) AS open_price,
    FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time DESC) AS close_price,
    SUM(spt.volume) AS total_volume
FROM emitens e
INNER JOIN stock_price_ticks spt ON e.id = spt.emiten_id
GROUP BY e.id, e.symbol, e.name, DATE(spt.price_time)
ORDER BY e.symbol, DATE(spt.price_time) DESC;

-- View: Top gainers today
CREATE OR REPLACE VIEW v_top_gainers AS
SELECT
    e.symbol,
    e.name,
    e.sector,
    (close_price - open_price) AS change_amount,
    ((close_price - open_price) / open_price * 100) AS change_percent
FROM (
    SELECT
        e.id,
        e.symbol,
        e.name,
        e.sector,
        FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time ASC) AS open_price,
        FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time DESC) AS close_price
    FROM emitens e
    INNER JOIN stock_price_ticks spt ON e.id = spt.emiten_id
    WHERE DATE(spt.price_time) = CURRENT_DATE
) sub
WHERE open_price > 0
ORDER BY change_percent DESC
LIMIT 20;

-- View: Top losers today
CREATE OR REPLACE VIEW v_top_losers AS
SELECT
    e.symbol,
    e.name,
    e.sector,
    (close_price - open_price) AS change_amount,
    ((close_price - open_price) / open_price * 100) AS change_percent
FROM (
    SELECT
        e.id,
        e.symbol,
        e.name,
        e.sector,
        FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time ASC) AS open_price,
        FIRST_VALUE(spt.price) OVER (PARTITION BY e.id, DATE(spt.price_time) ORDER BY spt.price_time DESC) AS close_price
    FROM emitens e
    INNER JOIN stock_price_ticks spt ON e.id = spt.emiten_id
    WHERE DATE(spt.price_time) = CURRENT_DATE
) sub
WHERE open_price > 0
ORDER BY change_percent ASC
LIMIT 20;

-- View: Most active stocks by volume
CREATE OR REPLACE VIEW v_most_active_stocks AS
SELECT
    e.symbol,
    e.name,
    e.sector,
    COUNT(*) AS tick_count,
    SUM(spt.volume) AS total_volume,
    MIN(spt.price) AS min_price,
    MAX(spt.price) AS max_price,
    MAX(spt.price_time) AS last_trade_time
FROM emitens e
INNER JOIN stock_price_ticks spt ON e.id = spt.emiten_id
WHERE DATE(spt.price_time) = CURRENT_DATE
GROUP BY e.id, e.symbol, e.name, e.sector
ORDER BY total_volume DESC
LIMIT 20;

-- View: Stock price summary with emitens info
CREATE OR REPLACE VIEW v_stock_summary_detail AS
SELECT
    sps.id,
    e.symbol,
    e.name,
    e.sector,
    sps.date,
    sps.open_price,
    sps.high_price,
    sps.low_price,
    sps.close_price,
    (sps.close_price - sps.open_price) AS price_change,
    CASE
        WHEN sps.open_price > 0 THEN
            ROUND(((sps.close_price - sps.open_price) / sps.open_price * 100)::NUMERIC, 2)
        ELSE 0
    END AS change_percent,
    sps.created_at,
    sps.updated_at
FROM stock_price_summary sps
INNER JOIN emitens e ON sps.emiten_id = e.id
ORDER BY sps.date DESC, e.symbol;

-- View: Moving averages (50-day and 200-day)
CREATE OR REPLACE VIEW v_moving_averages AS
WITH ma_data AS (
    SELECT
        emiten_id,
        date,
        close_price,
        AVG(close_price) OVER (
            PARTITION BY emiten_id
            ORDER BY date
            ROWS BETWEEN 49 PRECEDING AND CURRENT ROW
        ) AS ma_50,
        AVG(close_price) OVER (
            PARTITION BY emiten_id
            ORDER BY date
            ROWS BETWEEN 199 PRECEDING AND CURRENT ROW
        ) AS ma_200
    FROM stock_price_summary
)
SELECT
    e.symbol,
    e.name,
    m.date,
    m.close_price,
    ROUND(m.ma_50::NUMERIC, 2) AS ma_50,
    ROUND(m.ma_200::NUMERIC, 2) AS ma_200,
    CASE
        WHEN m.ma_50 > m.ma_200 THEN 'Bullish'
        ELSE 'Bearish'
    END AS trend
FROM ma_data m
INNER JOIN emitens e ON m.emiten_id = e.id
ORDER BY e.symbol, m.date DESC;

-- View: Price gaps (opening price different from previous close)
CREATE OR REPLACE VIEW v_price_gaps AS
WITH prev_data AS (
    SELECT
        emiten_id,
        date,
        close_price,
        LEAD(date) OVER (PARTITION BY emiten_id ORDER BY date) AS next_date,
        LEAD(close_price) OVER (PARTITION BY emiten_id ORDER BY date) AS next_close
    FROM stock_price_summary
)
SELECT
    e.symbol,
    e.name,
    p.date,
    p.close_price AS prev_close,
    s.open_price AS next_open,
    (s.open_price - p.close_price) AS gap_amount,
    CASE
        WHEN p.close_price > 0 THEN
            ROUND(((s.open_price - p.close_price) / p.close_price * 100)::NUMERIC, 2)
        ELSE 0
    END AS gap_percent,
    CASE
        WHEN s.open_price > p.close_price THEN 'Gap Up'
        WHEN s.open_price < p.close_price THEN 'Gap Down'
        ELSE 'No Gap'
    END AS gap_type
FROM prev_data p
INNER JOIN stock_price_summary s ON p.next_date = s.date AND p.emiten_id = s.emiten_id
INNER JOIN emitens e ON p.emiten_id = e.id
WHERE p.next_date IS NOT NULL
  AND s.open_price != p.close_price
ORDER BY p.date DESC;
