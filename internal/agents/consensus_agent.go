package agents

import (
	"context"
)

type ConsensusAgent struct {
}

func NewConsensusAgent() *ConsensusAgent {
	return &ConsensusAgent{}
}

type ConsensusRequest struct {
	CompanyName    string
	AgentDecisions []AgentDecision
}

type AgentDecision struct {
	AgentName  string
	Decision   string
	Confidence float64
	Reasoning  string
}

type ConsensusResponse struct {
	FinalDecision       string
	ConsensusConfidence float64
	SupportingAgents    []string
	DissentingAgents    []string
}

// ReachConsensus aggregates decisions from multiple agents
func (c *ConsensusAgent) ReachConsensus(ctx context.Context, req *ConsensusRequest) (*ConsensusResponse, error) {
	response := &ConsensusResponse{
		SupportingAgents: []string{},
		DissentingAgents: []string{},
	}

	if len(req.AgentDecisions) == 0 {
		response.FinalDecision = "INSUFFICIENT_DATA"
		response.ConsensusConfidence = 0.0
		return response, nil
	}

	// Count votes and calculate weighted consensus
	votes := make(map[string]float64)
	totalWeight := 0.0

	for _, decision := range req.AgentDecisions {
		weight := decision.Confidence
		votes[decision.Decision] += weight
		totalWeight += weight
	}

	// Find majority decision
	maxVotes := 0.0
	majorityDecision := ""
	for decision, voteWeight := range votes {
		if voteWeight > maxVotes {
			maxVotes = voteWeight
			majorityDecision = decision
		}
	}

	response.FinalDecision = majorityDecision
	response.ConsensusConfidence = maxVotes / totalWeight

	// Categorize agents
	for _, decision := range req.AgentDecisions {
		if decision.Decision == majorityDecision {
			response.SupportingAgents = append(response.SupportingAgents, decision.AgentName)
		} else {
			response.DissentingAgents = append(response.DissentingAgents, decision.AgentName)
		}
	}

	// Apply consensus threshold
	if response.ConsensusConfidence < 0.5 {
		response.FinalDecision = "NO_CONSENSUS"
	}

	return response, nil
}
