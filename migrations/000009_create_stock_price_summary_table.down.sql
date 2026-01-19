-- Rollback: Drop stock_price_summary table
-- Version: 000009
-- Description: Drop stock_price_summary table

-- Drop indexes
DROP INDEX IF EXISTS idx_stock_price_summary_emiten_id;
DROP INDEX IF EXISTS idx_stock_price_summary_date;
DROP INDEX IF EXISTS idx_stock_price_summary_emiten_date;
DROP INDEX IF EXISTS idx_stock_price_summary_close_price;

-- Drop trigger
DROP TRIGGER IF EXISTS update_stock_price_summary_updated_at ON stock_price_summary;

-- Drop constraints
ALTER TABLE stock_price_summary
DROP CONSTRAINT IF EXISTS fk_stock_price_summary_emiten_id;

ALTER TABLE stock_price_summary
DROP CONSTRAINT IF EXISTS unique_summary;

ALTER TABLE stock_price_summary
DROP CONSTRAINT IF EXISTS check_price_order;

-- Drop table
DROP TABLE IF EXISTS stock_price_summary;
