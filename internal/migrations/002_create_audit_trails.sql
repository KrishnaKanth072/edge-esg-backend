-- Audit Trail table for compliance
CREATE TABLE IF NOT EXISTS audit_trails (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bank_id UUID NOT NULL,
    user_id TEXT NOT NULL,
    action TEXT NOT NULL,
    resource TEXT,
    details JSONB,
    ip_address TEXT,
    blockchain_tx TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Enable RLS
ALTER TABLE audit_trails ENABLE ROW LEVEL SECURITY;

-- RLS Policy
CREATE POLICY audit_bank_isolation ON audit_trails
    USING (bank_id = current_setting('app.current_bank')::uuid);

-- Indexes
CREATE INDEX idx_audit_trails_bank_id ON audit_trails(bank_id);
CREATE INDEX idx_audit_trails_user_id ON audit_trails(user_id);
CREATE INDEX idx_audit_trails_created_at ON audit_trails(created_at DESC);
