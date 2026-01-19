-- Migration: Create stock_price_ticks table
-- Version: 000008
-- Description: Create stock_price_ticks table for historical stock price data

-- Create stock_price_ticks table
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

-- Create indexes for performance
CREATE INDEX idx_stock_price_ticks_emiten_id ON stock_price_ticks(emiten_id);
CREATE INDEX idx_stock_price_ticks_price_time ON stock_price_ticks(price_time);
CREATE INDEX idx_stock_price_ticks_emiten_time ON stock_price_ticks(emiten_id, price_time DESC);
CREATE INDEX idx_stock_price_ticks_price ON stock_price_ticks(price);
CREATE INDEX idx_stock_price_ticks_volume ON stock_price_ticks(volume);

-- Add comments
COMMENT ON TABLE stock_price_ticks IS 'Historical stock price tick data';
COMMENT ON COLUMN stock_price_ticks.id IS 'Unique identifier (auto-increment)';
COMMENT ON COLUMN stock_price_ticks.emiten_id IS 'Reference to emiten';
COMMENT ON COLUMN stock_price_ticks.price_time IS 'Timestamp of the price';
COMMENT ON COLUMN stock_price_ticks.price IS 'Stock price (numeric with 2 decimals)';
COMMENT ON COLUMN stock_price_ticks.volume IS 'Trading volume';
COMMENT ON COLUMN stock_price_ticks.created_at IS 'Record creation timestamp';
COMMENT ON CONSTRAINT unique_tick ON stock_price_ticks IS 'Ensures no duplicate price for same emiten and time';
