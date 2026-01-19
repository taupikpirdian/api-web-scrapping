# Stock Price Summary Documentation

## Table: stock_price_summary

Table ini menyimpan data harga saham harian dalam format OHLC (Open, High, Low, Close). Ini adalah format standar untuk data pasar keuangan.

### Schema

```sql
CREATE TABLE stock_price_summary (
    id BIGSERIAL PRIMARY KEY,
    emiten_id BIGINT NOT NULL,              -- Foreign key ke emitens
    date DATE NOT NULL,                     -- Tanggal summary (2026-01-19)
    open_price NUMERIC(12,2) NOT NULL,      -- Harga pembukaan
    high_price NUMERIC(12,2) NOT NULL,      -- Harga tertinggi
    low_price NUMERIC(12,2) NOT NULL,       -- Harga terendah
    close_price NUMERIC(12,2) NOT NULL,     -- Harga penutupan
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (emiten_id, date),               -- Satu summary per emiten per hari
    FOREIGN KEY (emiten_id) REFERENCES emitens(id) ON DELETE CASCADE,
    CHECK (high_price >= low_price AND ...) -- Validasi logika harga
);
```

### Constraints

1. **unique_summary**: Mencegah duplicate summary untuk emiten dan tanggal yang sama
2. **fk_stock_price_summary_emiten_id**: Foreign key ke emitens dengan CASCADE delete
3. **check_price_order**: Validasi bahwa harga masuk akal:
   - high_price ≥ low_price
   - high_price ≥ open_price
   - high_price ≥ close_price
   - low_price ≤ open_price
   - low_price ≤ close_price

### Indexes

- `idx_stock_price_summary_emiten_id` - Query by emiten
- `idx_stock_price_summary_date` - Query by tanggal
- `idx_stock_price_summary_emiten_date` - Composite index (emiten + tanggal)
- `idx_stock_price_summary_close_price` - Query by harga penutupan

## Auto-Generate Summary Functions

### 1. generate_stock_summary(p_emiten_id, p_date)

Generate summary untuk satu emiten pada tanggal tertentu dari tick data.

```sql
-- Generate summary untuk BBCA (id=1) untuk hari ini
SELECT generate_stock_summary(1, CURRENT_DATE);

-- Generate summary untuk tanggal spesifik
SELECT generate_stock_summary(1, '2026-01-19'::DATE);
```

### 2. generate_all_summaries(p_date)

Generate summary untuk SEMUA emitens yang memiliki tick data pada tanggal tersebut.

```sql
-- Generate summary untuk hari ini
SELECT generate_all_summaries();

-- Generate summary untuk tanggal spesifik
SELECT generate_all_summaries('2026-01-19'::DATE);
```

## Views

### v_stock_summary_detail

View lengkap dengan informasi emiten dan perubahan harga.

```sql
SELECT * FROM v_stock_summary_detail
WHERE symbol = 'BBCA'
ORDER BY date DESC
LIMIT 30;
```

**Columns:**
- symbol, name, sector
- date
- open_price, high_price, low_price, close_price
- price_change (close - open)
- change_percent (persentase perubahan)

### v_moving_averages

View untuk moving average 50-hari dan 200-hari (indikator teknikal).

```sql
SELECT * FROM v_moving_averages
WHERE symbol = 'BBCA'
ORDER BY date DESC
LIMIT 10;
```

**Columns:**
- symbol, name
- date
- close_price
- ma_50 (Moving Average 50 hari)
- ma_200 (Moving Average 200 hari)
- trend ('Bullish' jika MA50 > MA200, 'Bearish' jika sebaliknya)

### v_price_gaps

View untuk mendeteksi price gaps (ketika harga open berbeda signifikan dari close sebelumnya).

```sql
SELECT * FROM v_price_gaps
WHERE symbol = 'BBCA'
ORDER BY date DESC
LIMIT 20;
```

**Columns:**
- symbol, name
- date (tanggal gap)
- prev_close (harga close hari sebelumnya)
- next_open (harga open hari ini)
- gap_amount (selisih harga)
- gap_percent (persentase gap)
- gap_type ('Gap Up', 'Gap Down', atau 'No Gap')

## Common Queries

### Insert Summary Manual

```sql
-- Insert summary baru
INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
VALUES (1, '2026-01-19', 9200, 9300, 9150, 9250);

-- Upsert (insert atau update jika sudah ada)
INSERT INTO stock_price_summary (emiten_id, date, open_price, high_price, low_price, close_price)
VALUES (1, '2026-01-19', 9200, 9300, 9150, 9250)
ON CONFLICT (emiten_id, date) DO UPDATE SET
    open_price = EXCLUDED.open_price,
    high_price = EXCLUDED.high_price,
    low_price = EXCLUDED.low_price,
    close_price = EXCLUDED.close_price;
```

### Get Historical Data by Symbol

```sql
SELECT
    date,
    open_price,
    high_price,
    low_price,
    close_price,
    (close_price - open_price) AS change,
    ROUND(((close_price - open_price) / open_price * 100)::NUMERIC, 2) AS change_pct
FROM v_stock_summary_detail
WHERE symbol = 'BBCA'
ORDER BY date DESC
LIMIT 90;  -- Last 90 days (3 months)
```

### Get Price Range in Period

```sql
SELECT
    symbol,
    MIN(low_price) AS period_low,
    MAX(high_price) AS period_high,
    MAX(high_price) - MIN(low_price) AS price_range,
    ROUND(((MAX(high_price) - MIN(low_price)) / MIN(low_price) * 100)::NUMERIC, 2) AS range_pct
FROM v_stock_summary_detail
WHERE symbol = 'BBCA'
  AND date >= NOW() - INTERVAL '52 weeks'
GROUP BY symbol;
```

### Get Weekly/Monthly Summary

```sql
-- Monthly OHLC (dengan manipulasi date)
SELECT
    TO_CHAR(date, 'YYYY-MM') AS month,
    FIRST_VALUE(open_price) OVER (ORDER BY date ASC) AS month_open,
    MAX(high_price) AS month_high,
    MIN(low_price) AS month_low,
    FIRST_VALUE(close_price) OVER (ORDER BY date DESC) AS month_close,
    SUM(volume) AS total_volume
FROM stock_price_summary
WHERE emiten_id = 1
  AND date >= NOW() - INTERVAL '12 months'
GROUP BY TO_CHAR(date, 'YYYY-MM'), date, open_price, close_price, volume
ORDER BY month DESC;
```

### Compare Two Emitens

```sql
WITH bbca AS (
    SELECT date, close_price AS bbca_close
    FROM v_stock_summary_detail
    WHERE symbol = 'BBCA'
),
tlkm AS (
    SELECT date, close_price AS tlkm_close
    FROM v_stock_summary_detail
    WHERE symbol = 'TLKM'
)
SELECT
    COALESCE(bbca.date, tlkm.date) AS date,
    bbca.bbca_close,
    tlkm.tlkm_close,
    CASE
        WHEN bbca.bbca_close > tlkm.tlkm_close THEN 'BBCA Higher'
        ELSE 'TLKM Higher'
    END AS comparison
FROM bbca
FULL OUTER JOIN tlkm ON bbca.date = tlkm.date
ORDER BY date DESC
LIMIT 30;
```

### Get Top Performers by Date Range

```sql
SELECT
    symbol,
    name,
    sector,
    MIN(date) AS start_date,
    MAX(date) AS end_date,
    FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date ASC) AS start_price,
    FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date DESC) AS end_price,
    (FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date DESC) -
     FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date ASC)) AS price_change,
    ROUND(((FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date DESC) -
            FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date ASC)) /
           FIRST_VALUE(close_price) OVER (PARTITION BY symbol ORDER BY date ASC) * 100)::NUMERIC, 2) AS performance_pct
FROM v_stock_summary_detail
WHERE date >= NOW() - INTERVAL '30 days'
GROUP BY symbol, name, sector, date, close_price
HAVING COUNT(*) >= 20  -- Minimal 20 hari data
ORDER BY performance_pct DESC
LIMIT 10;
```

### Calculate Technical Indicators

```sql
-- RSI (Relative Strength Index) - Simplified
WITH avg_gain AS (
    SELECT
        date,
        close_price,
        CASE WHEN close_price > LAG(close_price) OVER (ORDER BY date)
             THEN close_price - LAG(close_price) OVER (ORDER BY date)
             ELSE 0 END AS gain,
        CASE WHEN close_price < LAG(close_price) OVER (ORDER BY date)
             THEN LAG(close_price) OVER (ORDER BY date) - close_price
             ELSE 0 END AS loss
    FROM stock_price_summary
    WHERE emiten_id = 1
    AND date >= NOW() - INTERVAL '15 days'
),
avg_calc AS (
    SELECT
        date,
        close_price,
        AVG(gain) OVER (ORDER BY date ROWS BETWEEN 13 PRECEDING AND CURRENT ROW) AS avg_gain,
        AVG(loss) OVER (ORDER BY date ROWS BETWEEN 13 PRECEDING AND CURRENT ROW) AS avg_loss
    FROM avg_gain
)
SELECT
    date,
    close_price,
    ROUND(100 - (100 / (1 + (avg_gain / NULLIF(avg_loss, 0))))::NUMERIC, 2) AS rsi
FROM avg_calc
ORDER BY date DESC;
```

### Find Breakout Stocks

```sql
-- Stocks yang break resistance (harga close tertinggi dalam 52 weeks)
SELECT
    symbol,
    name,
    date,
    close_price,
    MAX(close_price) OVER (PARTITION BY symbol ORDER BY date ROWS BETWEEN 251 PRECEDING AND CURRENT ROW) AS high_52_week,
    CASE
        WHEN close_price = MAX(close_price) OVER (PARTITION BY symbol ORDER BY date ROWS BETWEEN 251 PRECEDING AND CURRENT ROW)
        THEN '52-Week High'
        ELSE 'Normal'
    END AS status
FROM v_stock_summary_detail
WHERE date >= CURRENT_DATE - INTERVAL '1 week'
  AND close_price >= MAX(close_price) OVER (PARTITION BY symbol ORDER BY date ROWS BETWEEN 251 PRECEDING AND CURRENT ROW) * 0.99
ORDER BY date DESC, symbol;
```

## Automated Workflow

### Generate Summary from Tick Data

```sql
-- 1. Pastikan tick data sudah ada
SELECT COUNT(*) FROM stock_price_ticks
WHERE DATE(price_time) = CURRENT_DATE;

-- 2. Generate summary untuk hari ini
SELECT generate_all_summaries(CURRENT_DATE);

-- 3. Verifikasi hasil
SELECT * FROM v_stock_summary_detail
WHERE date = CURRENT_DATE
ORDER BY symbol;
```

### Schedule Auto-Generation (PostgreSQL cron)

```sql
-- Membuat extension pg_cron (jika available)
CREATE EXTENSION IF NOT EXISTS pg_cron;

-- Schedule generate summary setiap hari jam 17:00
SELECT cron.schedule('generate-daily-summary', '0 17 * * *',
    $$SELECT generate_all_summaries(CURRENT_DATE)$$);
```

### Trigger Auto-Generate on Insert

```sql
-- Trigger untuk auto-generate summary saat tick data di-insert
CREATE OR REPLACE FUNCTION trigger_generate_summary()
RETURNS TRIGGER AS $$
BEGIN
    -- Generate summary untuk emiten dan tanggal dari tick baru
    PERFORM generate_stock_summary(NEW.emiten_id, DATE(NEW.price_time));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER auto_generate_summary
AFTER INSERT ON stock_price_ticks
FOR EACH ROW
EXECUTE FUNCTION trigger_generate_summary();
```

## Best Practices

1. **Generate Summary End of Day**: Generate summary setelah market tutup untuk memastikan semua tick data sudah terkumpul
2. **Use Functions**: Gunakan fungsi `generate_stock_summary()` atau `generate_all_summaries()` untuk konsistensi
3. **Validate Data**: Constraint `check_price_order` memastikan data valid
4. **Partition for Large Data**: Pertimbangkan partitioning by year atau month jika data sangat besar
5. **Archive Old Data**: Archive data lebih dari 5-10 tahun ke table terpisah

## Performance Tips

1. **Composite Index**: Query yang sering filter by emiten + date sudah terindeks
2. **Use Views**: Views yang disediakan sudah teroptimasi
3. **Batch Operations**: Gunakan fungsi `generate_all_summaries()` untuk batch processing
4. **Regular Vacuum**: Jalankan `VACUUM ANALYZE stock_price_summary` secara rutin
5. **Monitor Index Usage**: Cek apakah semua indexes digunakan dengan `EXPLAIN ANALYZE`

## Data Example

```sql
-- Sample data
emiten_id | date       | open_price | high_price | low_price | close_price
----------|------------|------------|------------|-----------|-------------
1         | 2026-01-19 | 9200       | 9300       | 9150      | 9250
1         | 2026-01-18 | 9150       | 9250       | 9100      | 9180
1         | 2026-01-17 | 9100       | 9200       | 9050      | 9150
```

## Related Tables

- **emitens** - Master data emiten/saham
- **stock_price_ticks** - Tick data intraday (raw data untuk summary)
- **v_stock_summary_detail** - View dengan detail lengkap
- **v_moving_averages** - View untuk indikator teknikal
- **v_price_gaps** - View untuk deteksi gap harga
