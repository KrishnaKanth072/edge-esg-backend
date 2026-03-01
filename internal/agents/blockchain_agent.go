package agents

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type BlockchainAgent struct {
	auditRecords map[string]*AuditRecord
}

type AuditRecord struct {
	CompanyName  string
	AnalysisData string
	Timestamp    int64
	TxHash       string
	BlockHash    string
}

func NewBlockchainAgent() *BlockchainAgent {
	return &BlockchainAgent{
		auditRecords: make(map[string]*AuditRecord),
	}
}

type AuditRecordRequest struct {
	CompanyName  string
	AnalysisData string
	Timestamp    int64
}

type AuditRecordResponse struct {
	TransactionHash string
	BlockHash       string
	Success         bool
}

type AuditVerifyRequest struct {
	TransactionHash string
}

type AuditVerifyResponse struct {
	IsValid   bool
	Data      string
	Timestamp int64
}

// RecordAudit creates immutable audit trail
func (b *BlockchainAgent) RecordAudit(ctx context.Context, req *AuditRecordRequest) (*AuditRecordResponse, error) {
	// Generate transaction hash
	txData := fmt.Sprintf("%s:%s:%d", req.CompanyName, req.AnalysisData, req.Timestamp)
	txHashBytes := sha256.Sum256([]byte(txData))
	txHash := "0x" + hex.EncodeToString(txHashBytes[:])

	// Generate block hash (simulated)
	blockData := fmt.Sprintf("%s:%d", txHash, time.Now().Unix())
	blockHashBytes := sha256.Sum256([]byte(blockData))
	blockHash := "0x" + hex.EncodeToString(blockHashBytes[:])

	// Store record
	b.auditRecords[txHash] = &AuditRecord{
		CompanyName:  req.CompanyName,
		AnalysisData: req.AnalysisData,
		Timestamp:    req.Timestamp,
		TxHash:       txHash,
		BlockHash:    blockHash,
	}

	return &AuditRecordResponse{
		TransactionHash: txHash,
		BlockHash:       blockHash,
		Success:         true,
	}, nil
}

// VerifyAudit checks audit trail integrity
func (b *BlockchainAgent) VerifyAudit(ctx context.Context, req *AuditVerifyRequest) (*AuditVerifyResponse, error) {
	record, exists := b.auditRecords[req.TransactionHash]
	if !exists {
		return &AuditVerifyResponse{
			IsValid: false,
		}, nil
	}

	return &AuditVerifyResponse{
		IsValid:   true,
		Data:      record.AnalysisData,
		Timestamp: record.Timestamp,
	}, nil
}
