-- Rollback: Drop emitens table
-- Version: 000006
-- Description: Drop emitens table

-- Drop indexes
DROP INDEX IF EXISTS idx_emitens_symbol;
DROP INDEX IF EXISTS idx_emitens_sector;
DROP INDEX IF EXISTS idx_emitens_name;

-- Drop table
DROP TABLE IF EXISTS emitens;
