-- Enable pgcrypto extension for encryption
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ESG Scores table with RLS
CREATE TABLE IF NOT EXISTS esg_scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bank_id UUID NOT NULL,
    company_name TEXT,
    company_encrypted BYTEA,
    revenue_encrypted BYTEA,
    carbon_emissions_encrypted BYTEA,
    esg_score NUMERIC(3,2) NOT NULL,
    trading_signal JSONB,
    risk_action TEXT,
    audit_hash TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Enable Row Level Security
ALTER TABLE esg_scores ENABLE ROW LEVEL SECURITY;

-- RLS Policy: Bank isolation
CREATE POLICY bank_isolation ON esg_scores
    USING (bank_id = current_setting('app.current_bank')::uuid);

-- Indexes for performance
CREATE INDEX idx_esg_scores_bank_id ON esg_scores(bank_id);
CREATE INDEX idx_esg_scores_created_at ON esg_scores(created_at DESC);
CREATE INDEX idx_esg_scores_company_name ON esg_scores(company_name);
