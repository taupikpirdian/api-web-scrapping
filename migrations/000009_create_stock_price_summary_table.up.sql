-- Migration: Create stock_price_summary table
-- Version: 000009
-- Description: Create stock_price_summary table for daily OHLC (Open, High, Low, Close) data

-- Create stock_price_summary table
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

-- Create indexes for performance
CREATE INDEX idx_stock_price_summary_emiten_id ON stock_price_summary(emiten_id);
CREATE INDEX idx_stock_price_summary_date ON stock_price_summary(date);
CREATE INDEX idx_stock_price_summary_emiten_date ON stock_price_summary(emiten_id, date DESC);
CREATE INDEX idx_stock_price_summary_close_price ON stock_price_summary(close_price);

-- Add comments
COMMENT ON TABLE stock_price_summary IS 'Daily stock price summary (OHLC)';
COMMENT ON COLUMN stock_price_summary.id IS 'Unique identifier (auto-increment)';
COMMENT ON COLUMN stock_price_summary.emiten_id IS 'Reference to emiten';
COMMENT ON COLUMN stock_price_summary.date IS 'Summary date (YYYY-MM-DD)';
COMMENT ON COLUMN stock_price_summary.open_price IS 'Opening price for the day';
COMMENT ON COLUMN stock_price_summary.high_price IS 'Highest price for the day';
COMMENT ON COLUMN stock_price_summary.low_price IS 'Lowest price for the day';
COMMENT ON COLUMN stock_price_summary.close_price IS 'Closing price for the day';
COMMENT ON COLUMN stock_price_summary.created_at IS 'Record creation timestamp';
COMMENT ON COLUMN stock_price_summary.updated_at IS 'Last update timestamp';
COMMENT ON CONSTRAINT unique_summary ON stock_price_summary IS 'Ensures one summary per emiten per day';
COMMENT ON CONSTRAINT check_price_order ON stock_price_summary IS 'Validates price logic (high >= low, etc)';
