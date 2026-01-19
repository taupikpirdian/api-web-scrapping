-- Migration: Seed emitens data
-- Version: 000007
-- Description: Insert sample Indonesian stock emitens

-- Insert Indonesian stocks (IHSG)
INSERT INTO emitens (symbol, name, sector) VALUES
-- Banking
('BBCA', 'Bank Central Asia Tbk', 'Finance'),
('BBRI', 'Bank Rakyat Indonesia Tbk', 'Finance'),
('BBNI', 'Bank Negara Indonesia Tbk', 'Finance'),
('BMRI', 'Bank Mandiri Tbk', 'Finance'),

-- Technology
('TLKM', 'Telkom Indonesia Tbk', 'Technology'),
('EXCL', 'XL Axiata Tbk', 'Technology'),
('ISAT', 'Indosat Ooredoo Tbk', 'Technology'),
('FREN', 'Smartfren Telecom Tbk', 'Technology'),

-- Consumer Goods
('UNVR', 'Unilever Indonesia Tbk', 'Consumer Goods'),
('INDF', 'Indofood Sukses Makmur Tbk', 'Consumer Goods'),
('ICBP', 'Indofood CBP Sukses Makmur Tbk', 'Consumer Goods'),
('ULTJ', 'Ultrajaya Milk Industry Tbk', 'Consumer Goods'),

-- Energy & Resources
('ADRO', 'Adaro Energy Indonesia Tbk', 'Energy'),
('ITMG', 'Indo Tambangraya Megah Tbk', 'Energy'),
('PGAS', 'Perusahaan Gas Negara Tbk', 'Energy'),
('PTBA', 'Bukit Asam Tbk', 'Energy'),
('TPIA', 'Chandra Asri Petrochemical Tbk', 'Energy'),

-- Industrial
('UNTR', 'United Tractors Tbk', 'Industrial'),
('INTP', 'Indocement Tunggal Prakasa Tbk', 'Industrial'),
('SMGR', 'Semen Indonesia Tbk', 'Industrial'),
('ANTM', 'Aneka Tambang Tbk', 'Mining'),
('TKIM', 'Pabrik Kertas Tjiwnd Kimia Tbk', 'Industrial'),

-- Property & Real Estate
('BSDE', 'Bumi Serpong Damai Tbk', 'Property'),
('ASII', 'Astra International Tbk', 'Industrial'),
('AUTO', 'Astra Otoparts Tbk', 'Industrial'),
('GJTL', 'Gajah Tunggal Tbk', 'Industrial'),

-- Retail & Trade
('AMRT', 'Sumber Alfaria Trijaya Tbk', 'Retail'),
('ICBP', 'Indofood CBP Sukses Makmur Tbk', 'Consumer Goods'),
('LPPF', 'Matahari Department Store Tbk', 'Retail'),
('MNCN', 'Media Nusantara Citra Tbk', 'Media'),

-- Healthcare
('KLBF', 'Kalbe Farma Tbk', 'Healthcare'),
('PYID', 'Pyridam Farma Tbk', 'Healthcare'),
('MYOR', 'Mayora Indah Tbk', 'Consumer Goods'),

-- Infrastructure
('JSMR', 'Jasa Marga Tbk', 'Infrastructure'),
('WIKA', 'Wijaya Karya Tbk', 'Infrastructure'),
('PTPP', 'PP Tbk', 'Infrastructure'),
('ADHI', 'Adhi Karya Tbk', 'Infrastructure')

ON CONFLICT (symbol) DO NOTHING;

-- Add comment
COMMENT ON TABLE emitens IS 'Populated with Indonesian stock market data';
