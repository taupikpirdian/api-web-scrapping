# Stock Data Database Schema

## Tables

### 1. emitens
Table untuk menyimpan data emiten/saham.

**Columns:**
- `id` (BIGSERIAL) - Primary key, auto-increment
- `symbol` (VARCHAR(10)) - Kode saham (UNIQUE, contoh: BBCA, TLKM)
- `name` (VARCHAR(100)) - Nama perusahaan
- `sector` (VARCHAR(50)) - Sektor industri
- `created_at` (TIMESTAMP) - Waktu pembuatan record
- `updated_at` (TIMESTAMP) - Waktu terakhir update

**Indexes:**
- `idx_emitens_symbol` - Untuk lookup by symbol
- `idx_emitens_sector` - Untuk filter by sector
- `idx_emitens_name` - Untuk search by name

**Trigger:**
- `update_emitens_updated_at` - Auto update `updated_at` saat record di-update

### 2. stock_price_ticks
Table untuk menyimpan data harga saham historis (tick data).

**Columns:**
- `id` (BIGSERIAL) - Primary key, auto-increment
- `emiten_id` (BIGINT) - Foreign key ke table emitens
- `price_time` (TIMESTAMP) - Waktu harga tercatat
- `price` (NUMERIC(12,2)) - Harga saham (2 desimal)
- `volume` (BIGINT) - Volume trading (opsional)
- `created_at` (TIMESTAMP) - Waktu pembuatan record

**Constraints:**
- `unique_tick` - UNIQUE(emiten_id, price_time) - Mencegah duplicate harga untuk emiten & waktu yang sama
- `fk_stock_price_ticks_emiten_id` - Foreign key ke emitens(id) dengan CASCADE delete

**Indexes:**
- `idx_stock_price_ticks_emiten_id` - Untuk query by emiten
- `idx_stock_price_ticks_price_time` - Untuk query by waktu
- `idx_stock_price_ticks_emiten_time` - Composite index untuk query harga emiten tertentu berdasarkan waktu
- `idx_stock_price_ticks_price` - Untuk query by harga
- `idx_stock_price_ticks_volume` - Untuk query by volume

## Views

### v_latest_stock_prices
Menampilkan harga terakhir untuk setiap emiten.

```sql
SELECT * FROM v_latest_stock_prices WHERE symbol = 'BBCA';
```

### v_stock_price_summary
Menampilkan summary harga harian (Open, High, Low, Close, Volume).

```sql
SELECT * FROM v_stock_price_summary
WHERE symbol = 'BBCA'
ORDER BY price_date DESC
LIMIT 30;
```

### v_top_gainers
Top 20 saham dengan kenaikan harga terbesar hari ini.

```sql
SELECT * FROM v_top_gainers;
```

### v_top_losers
Top 20 saham dengan penurunan harga terbesar hari ini.

```sql
SELECT * FROM v_top_losers;
```

### v_most_active_stocks
Top 20 saham paling aktif berdasarkan volume hari ini.

```sql
SELECT * FROM v_most_active_stocks;
```

## Common Queries

### Insert Emitens

```sql
INSERT INTO emitens (symbol, name, sector)
VALUES ('BBCA', 'Bank Central Asia Tbk', 'Finance');
```

### Insert Stock Price (Upsert untuk handle duplicate)

```sql
INSERT INTO stock_price_ticks (emiten_id, price_time, price, volume)
VALUES (1, '2024-01-15 09:00:00', 9250, 1000000)
ON CONFLICT (emiten_id, price_time) DO UPDATE SET
    price = EXCLUDED.price,
    volume = EXCLUDED.volume;
```

### Get Latest Price by Symbol

```sql
SELECT e.symbol, e.name, spt.price, spt.price_time
FROM emitens e
LEFT JOIN stock_price_ticks spt ON e.id = spt.emiten_id
WHERE e.symbol = 'BBCA'
ORDER BY spt.price_time DESC
LIMIT 1;
```

### Get Price History (Last 30 Days)

```sql
SELECT
    DATE(spt.price_time) AS date,
    MIN(spt.price) AS low,
    MAX(spt.price) AS high,
    FIRST_VALUE(spt.price) OVER (ORDER BY spt.price_time ASC) AS open,
    FIRST_VALUE(spt.price) OVER (ORDER BY spt.price_time DESC) AS close,
    SUM(spt.volume) OVER () AS volume
FROM stock_price_ticks spt
WHERE spt.emiten_id = 1
  AND spt.price_time >= NOW() - INTERVAL '30 days'
GROUP BY DATE(spt.price_time), spt.price, spt.volume
ORDER BY date DESC;
```

### Get All Emitens by Sector

```sql
SELECT * FROM emitens
WHERE sector = 'Finance'
ORDER BY symbol;
```

### Search Emitens by Name

```sql
SELECT * FROM emitens
WHERE name ILIKE '%bank%'
ORDER BY symbol;
```

### Get Price Movement (Percentage Change)

```sql
WITH latest AS (
    SELECT DISTINCT ON (emiten_id)
        emiten_id, price, price_time
    FROM stock_price_ticks
    ORDER BY emiten_id, price_time DESC
),
previous AS (
    SELECT DISTINCT ON (emiten_id)
        emiten_id, price, price_time
    FROM stock_price_ticks
    WHERE price_time < (
        SELECT MAX(price_time) FROM stock_price_ticks
    )
    ORDER BY emiten_id, price_time DESC
)
SELECT
    e.symbol,
    e.name,
    l.price AS latest_price,
    p.price AS previous_price,
    ((l.price - p.price) / p.price * 100) AS change_percent
FROM emitens e
JOIN latest l ON e.id = l.emiten_id
LEFT JOIN previous p ON e.id = p.emiten_id;
```

### Get Average Price by Day

```sql
SELECT
    DATE(price_time) AS date,
    AVG(price) AS avg_price,
    MIN(price) AS min_price,
    MAX(price) AS max_price,
    SUM(volume) AS total_volume
FROM stock_price_ticks
WHERE emiten_id = 1
  AND price_time >= NOW() - INTERVAL '7 days'
GROUP BY DATE(price_time)
ORDER BY date DESC;
```

### Get Price Range (High/Low) in Period

```sql
SELECT
    e.symbol,
    MIN(spt.price) AS period_low,
    MAX(spt.price) AS period_high,
    MAX(spt.price) - MIN(spt.price) AS price_range,
    (MAX(spt.price) - MIN(spt.price)) / MIN(spt.price) * 100 AS range_percent
FROM emitens e
JOIN stock_price_ticks spt ON e.id = spt.emiten_id
WHERE e.symbol = 'BBCA'
  AND spt.price_time >= NOW() - INTERVAL '30 days'
GROUP BY e.symbol;
```

### Get Top Volume by Emitens

```sql
SELECT
    e.symbol,
    e.name,
    COUNT(spt.id) AS tick_count,
    SUM(spt.volume) AS total_volume,
    AVG(spt.price) AS avg_price
FROM emitens e
JOIN stock_price_ticks spt ON e.id = spt.emiten_id
WHERE DATE(spt.price_time) = CURRENT_DATE
GROUP BY e.id, e.symbol, e.name
HAVING SUM(spt.volume) > 0
ORDER BY total_volume DESC
LIMIT 10;
```

## Bulk Operations

### Bulk Insert Stock Prices

```sql
COPY stock_price_ticks (emiten_id, price_time, price, volume)
FROM '/path/to/stock_prices.csv'
DELIMITER ','
CSV HEADER;
```

### Bulk Upsert (PostgreSQL 15+)

```sql
INSERT INTO stock_price_ticks (emiten_id, price_time, price, volume)
VALUES
    (1, '2024-01-15 09:00:00', 9250, 1000000),
    (1, '2024-01-15 09:01:00', 9260, 1500000),
    (2, '2024-01-15 09:00:00', 3500, 2000000)
ON CONFLICT (emiten_id, price_time) DO UPDATE SET
    price = EXCLUDED.price,
    volume = EXCLUDED.volume;
```

## Maintenance

### Clean Old Price Data (Keep Last 1 Year)

```sql
DELETE FROM stock_price_ticks
WHERE price_time < NOW() - INTERVAL '1 year';
```

### Archive Old Data to Another Table

```sql
CREATE TABLE stock_price_ticks_archive AS
SELECT * FROM stock_price_ticks
WHERE price_time < NOW() - INTERVAL '6 months';

DELETE FROM stock_price_ticks
WHERE price_time < NOW() - INTERVAL '6 months';
```

### Recalculate Statistics

```sql
ANALYZE emitens;
ANALYZE stock_price_ticks;
```

### Reindex Tables

```sql
REINDEX TABLE emitens;
REINDEX TABLE stock_price_ticks;
```

## Performance Tips

1. **Use Composite Index**: Query yang sering menggunakan `emiten_id` dan `price_time` bersama-sama sudah terindeks
2. **Partition by Date**: Untuk data besar, pertimbangkan partitioning by date
3. **Use Views**: Views yang sudah dibuat sudah teroptimasi untuk query umum
4. **Bulk Operations**: Gunakan COPY atau batch INSERT untuk data besar
5. **Regular Vacuum**: Jalankan VACUUM ANALYZE secara rutin

## Sample Data

Emitens yang sudah di-seed (migration 000007):
- **Banking**: BBCA, BBRI, BBNI, BMRI
- **Technology**: TLKM, EXCL, ISAT, FREN
- **Consumer Goods**: UNVR, INDF, ICBP, ULTJ
- **Energy**: ADRO, ITMG, PGAS, PTBA, TPIA
- **Industrial**: UNTR, INTP, SMGR, ANTM, TKIM
- **Property**: BSDE, ASII, AUTO, GJTL
- **Retail**: AMRT, LPPF, MNCN
- **Healthcare**: KLBF, PYID, MYOR
- **Infrastructure**: JSMR, WIKA, PTPP, ADHI

## Connection String Examples

```bash
# Application (Go)
DATABASE_URL=postgres://admin:secret@localhost:5432/api_web_scrapping?sslmode=disable

# With connection pooling
DATABASE_URL=postgres://admin:secret@localhost:5432/api_web_scrapping?sslmode=disable&pool_max_conns=10

# psql command
psql -h localhost -p 5432 -U admin -d api_web_scrapping
```

## Next Steps

1. ✅ Run migrations: `make db-migrate`
2. ✅ Verify tables: `\dt` in psql
3. ✅ Test insert data
4. ✅ Query using views
5. ✅ Setup application connection
6. ✅ Implement scraping logic
