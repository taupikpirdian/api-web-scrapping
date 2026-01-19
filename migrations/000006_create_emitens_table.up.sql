-- Migration: Create emitens table
-- Version: 000006
-- Description: Create emitens table for stock/securities information

-- Create emitens table
CREATE TABLE IF NOT EXISTS emitens (
    id BIGSERIAL PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    sector VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_emitens_symbol ON emitens(symbol);
CREATE INDEX idx_emitens_sector ON emitens(sector);
CREATE INDEX idx_emitens_name ON emitens(name);

-- Add comments
COMMENT ON TABLE emitens IS 'Stock/securities emitens data';
COMMENT ON COLUMN emitens.id IS 'Unique identifier (auto-increment)';
COMMENT ON COLUMN emitens.symbol IS 'Stock symbol/ticker (unique)';
COMMENT ON COLUMN emitens.name IS 'Company or security name';
COMMENT ON COLUMN emitens.sector IS 'Industry sector classification';
COMMENT ON COLUMN emitens.created_at IS 'Record creation timestamp';
COMMENT ON COLUMN emitens.updated_at IS 'Last update timestamp';
