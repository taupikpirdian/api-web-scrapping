-- Rollback: Drop stock_price_ticks table
-- Version: 000008
-- Description: Drop stock_price_ticks table

-- Drop indexes
DROP INDEX IF EXISTS idx_stock_price_ticks_emiten_id;
DROP INDEX IF EXISTS idx_stock_price_ticks_price_time;
DROP INDEX IF EXISTS idx_stock_price_ticks_emiten_time;
DROP INDEX IF EXISTS idx_stock_price_ticks_price;
DROP INDEX IF EXISTS idx_stock_price_ticks_volume;

-- Drop foreign key constraint
ALTER TABLE stock_price_ticks
DROP CONSTRAINT IF EXISTS fk_stock_price_ticks_emiten_id;

-- Drop unique constraint
ALTER TABLE stock_price_ticks
DROP CONSTRAINT IF EXISTS unique_tick;

-- Drop table
DROP TABLE IF EXISTS stock_price_ticks;
