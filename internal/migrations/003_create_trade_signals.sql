-- Trade Signals table
CREATE TABLE IF NOT EXISTS trade_signals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bank_id UUID NOT NULL,
    symbol TEXT NOT NULL,
    action TEXT NOT NULL,
    target_price NUMERIC(10,2),
    confidence NUMERIC(3,2),
    esg_score NUMERIC(3,2),
    status TEXT DEFAULT 'PENDING',
    executed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Enable RLS
ALTER TABLE trade_signals ENABLE ROW LEVEL SECURITY;

-- RLS Policy
CREATE POLICY trade_bank_isolation ON trade_signals
    USING (bank_id = current_setting('app.current_bank')::uuid);

-- Indexes
CREATE INDEX idx_trade_signals_bank_id ON trade_signals(bank_id);
CREATE INDEX idx_trade_signals_symbol ON trade_signals(symbol);
CREATE INDEX idx_trade_signals_status ON trade_signals(status);
