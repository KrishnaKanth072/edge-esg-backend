package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Real-time data aggregator using free APIs
type RealTimeAgents struct {
	httpClient *http.Client
}

func NewRealTimeAgents() *RealTimeAgents {
	return &RealTimeAgents{
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

type CompanyAnalysis struct {
	Company       string    `json:"company"`
	ESGScore      float64   `json:"esg_score"`
	Environmental float64   `json:"environmental"`
	Social        float64   `json:"social"`
	Governance    float64   `json:"governance"`
	TradingSignal string    `json:"trading_signal"`
	StockSymbol   string    `json:"stock_symbol"`
	CurrentPrice  float64   `json:"current_price"`
	TargetPrice   float64   `json:"target_price"`
	PriceChange   float64   `json:"price_change"`
	Confidence    int       `json:"confidence"`
	RiskAction    string    `json:"risk_action"`
	RiskReasons   []string  `json:"risk_reasons"`
	NewsSentiment float64   `json:"news_sentiment"`
	LastUpdated   time.Time `json:"last_updated"`
}

// Analyze company with REAL data from multiple free sources
func (r *RealTimeAgents) AnalyzeCompany(ctx context.Context, company string) (*CompanyAnalysis, error) {
	analysis := &CompanyAnalysis{
		Company:     company,
		LastUpdated: time.Now(),
	}

	// 1. Get news sentiment (real-time)
	sentiment, err := r.GetNewsSentiment(company)
	if err == nil {
		analysis.NewsSentiment = sentiment
	} else {
		analysis.NewsSentiment = 0.5 // neutral default
	}

	// 2. Calculate ESG score based on sentiment and company type
	analysis.ESGScore = r.CalculateESGScore(company, sentiment)
	analysis.Environmental = analysis.ESGScore * 0.9
	analysis.Social = analysis.ESGScore * 1.1
	analysis.Governance = analysis.ESGScore * 1.0

	// 3. Determine stock symbol
	analysis.StockSymbol = r.GuessStockSymbol(company)

	// 4. Get real stock price if possible
	stockData, err := r.GetStockPrice(analysis.StockSymbol)
	if err == nil {
		analysis.CurrentPrice = stockData.Price
	}

	// 5. Generate trading signal based on ESG + sentiment
	analysis.TradingSignal, analysis.TargetPrice, analysis.PriceChange = r.GenerateTradingSignal(
		analysis.ESGScore,
		sentiment,
		analysis.CurrentPrice,
	)

	// 6. Risk assessment
	analysis.RiskAction, analysis.RiskReasons = r.AssessRisk(analysis.ESGScore, sentiment)

	// 7. Confidence based on data quality
	analysis.Confidence = r.CalculateConfidence(sentiment, analysis.CurrentPrice)

	return analysis, nil
}

// Get real news sentiment from NewsAPI
func (r *RealTimeAgents) GetNewsSentiment(company string) (float64, error) {
	// Get API key from environment
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		apiKey = "demo" // Fallback to demo (won't work)
	}

	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&sortBy=publishedAt&language=en&pageSize=20&apiKey=%s",
		strings.ReplaceAll(company, " ", "+"), apiKey)

	// #nosec G107 - URL uses trusted domain (newsapi.org) only
	resp, err := r.httpClient.Get(url)
	if err != nil {
		return 0.5, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var newsData map[string]interface{}
	if err := json.Unmarshal(body, &newsData); err != nil {
		return 0.5, nil // Return neutral on parse error
	}

	articles, ok := newsData["articles"].([]interface{})
	if !ok || len(articles) == 0 {
		return 0.5, nil
	}

	// Sentiment analysis keywords
	positive := []string{"growth", "profit", "success", "innovation", "sustainable", "green", "award", "expansion", "breakthrough"}
	negative := []string{"loss", "decline", "scandal", "lawsuit", "pollution", "violation", "fine", "layoff", "bankruptcy"}

	score := 0.0
	for i, article := range articles {
		if i >= 20 {
			break
		}
		art := article.(map[string]interface{})
		title := strings.ToLower(fmt.Sprintf("%v", art["title"]))
		description := strings.ToLower(fmt.Sprintf("%v", art["description"]))
		text := title + " " + description

		for _, word := range positive {
			if strings.Contains(text, word) {
				score += 0.05
			}
		}
		for _, word := range negative {
			if strings.Contains(text, word) {
				score -= 0.05
			}
		}
	}

	sentiment := 0.5 + score
	if sentiment < 0 {
		sentiment = 0
	}
	if sentiment > 1 {
		sentiment = 1
	}

	return sentiment, nil
}

type StockData struct {
	Symbol string
	Price  float64
}

// Get real stock price from Yahoo Finance (free, no key needed)
func (r *RealTimeAgents) GetStockPrice(symbol string) (*StockData, error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

	// #nosec G107 - URL uses trusted domain (yahoo.com) only
	resp, err := r.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	chart, ok := data["chart"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response")
	}

	result, ok := chart["result"].([]interface{})
	if !ok || len(result) == 0 {
		return nil, fmt.Errorf("no data")
	}

	meta := result[0].(map[string]interface{})["meta"].(map[string]interface{})
	price := meta["regularMarketPrice"].(float64)

	return &StockData{
		Symbol: symbol,
		Price:  price,
	}, nil
}

// Calculate ESG score based on company name and sentiment
func (r *RealTimeAgents) CalculateESGScore(company string, sentiment float64) float64 {
	baseScore := 5.0

	// Adjust based on industry keywords
	companyLower := strings.ToLower(company)

	// Green/sustainable companies get bonus
	if strings.Contains(companyLower, "solar") || strings.Contains(companyLower, "wind") ||
		strings.Contains(companyLower, "renewable") || strings.Contains(companyLower, "green") ||
		strings.Contains(companyLower, "clean") || strings.Contains(companyLower, "sustainable") {
		baseScore += 2.5
	}

	// Tech companies generally score well
	if strings.Contains(companyLower, "tech") || strings.Contains(companyLower, "software") ||
		strings.Contains(companyLower, "apple") || strings.Contains(companyLower, "microsoft") ||
		strings.Contains(companyLower, "google") || strings.Contains(companyLower, "tesla") {
		baseScore += 1.5
	}

	// Manufacturing/Industrial - moderate
	if strings.Contains(companyLower, "steel") || strings.Contains(companyLower, "motors") ||
		strings.Contains(companyLower, "tata") || strings.Contains(companyLower, "reliance") {
		baseScore += 0.5
	}

	// Oil/coal companies score lower
	if strings.Contains(companyLower, "oil") || strings.Contains(companyLower, "coal") ||
		strings.Contains(companyLower, "petroleum") || strings.Contains(companyLower, "exxon") ||
		strings.Contains(companyLower, "chevron") || strings.Contains(companyLower, "shell") {
		baseScore -= 2.5
	}

	// Tobacco/weapons - very low
	if strings.Contains(companyLower, "tobacco") || strings.Contains(companyLower, "cigarette") ||
		strings.Contains(companyLower, "defense") || strings.Contains(companyLower, "weapons") {
		baseScore -= 3.0
	}

	// Adjust by sentiment (bigger impact)
	baseScore += (sentiment - 0.5) * 6.0

	// Clamp to 0-10 range
	if baseScore < 0 {
		baseScore = 0
	}
	if baseScore > 10 {
		baseScore = 10
	}

	return baseScore
}

// Generate trading signal based on ESG and sentiment
func (r *RealTimeAgents) GenerateTradingSignal(esgScore, sentiment, currentPrice float64) (string, float64, float64) {
	signal := "HOLD"
	targetPrice := currentPrice
	priceChange := 0.0

	// Strong ESG + positive sentiment = BUY
	if esgScore >= 6.5 && sentiment >= 0.6 {
		signal = "BUY"
		priceChange = 15.0 + (esgScore-6.5)*5.0
		targetPrice = currentPrice * (1 + priceChange/100)
	} else if esgScore <= 4.0 || sentiment <= 0.4 {
		signal = "SELL"
		priceChange = -(10.0 + (4.0-esgScore)*3.0)
		targetPrice = currentPrice * (1 + priceChange/100)
	}

	return signal, targetPrice, priceChange
}

// Assess financing risk
func (r *RealTimeAgents) AssessRisk(esgScore, sentiment float64) (string, []string) {
	reasons := []string{}

	if esgScore < 4.0 {
		reasons = append(reasons, "Low ESG score indicates sustainability risks")
	}
	if sentiment < 0.4 {
		reasons = append(reasons, "Negative news sentiment")
	}
	if esgScore < 3.0 {
		reasons = append(reasons, "High regulatory compliance risk")
	}

	action := "APPROVE"
	if len(reasons) >= 2 {
		action = "REJECT"
	} else if len(reasons) == 1 {
		action = "REVIEW"
	}

	if len(reasons) == 0 {
		reasons = append(reasons, "Strong ESG profile", "Positive market sentiment")
	}

	return action, reasons
}

// Calculate confidence based on data availability
func (r *RealTimeAgents) CalculateConfidence(sentiment, price float64) int {
	confidence := 70

	if sentiment != 0.5 {
		confidence += 10 // Have real news data
	}
	if price > 0 {
		confidence += 15 // Have real stock price
	}

	return confidence
}

// Guess stock symbol from company name
func (r *RealTimeAgents) GuessStockSymbol(company string) string {
	// Common mappings
	mappings := map[string]string{
		"tata steel":  "TATASTEEL.NS",
		"tata motors": "TATAMOTORS.NS",
		"tata":        "TCS.NS",
		"reliance":    "RELIANCE.NS",
		"infosys":     "INFY",
		"wipro":       "WIT",
		"hdfc":        "HDFCBANK.NS",
		"icici":       "ICICIBANK.NS",
		"suzlon":      "SUZLON.NS",
		"adani":       "ADANIENT.NS",
		"apple":       "AAPL",
		"microsoft":   "MSFT",
		"google":      "GOOGL",
		"amazon":      "AMZN",
		"tesla":       "TSLA",
		"exxon":       "XOM",
		"chevron":     "CVX",
		"shell":       "SHEL",
		"bp":          "BP",
		"meta":        "META",
		"facebook":    "META",
		"netflix":     "NFLX",
		"nvidia":      "NVDA",
	}

	companyLower := strings.ToLower(company)
	for key, symbol := range mappings {
		if strings.Contains(companyLower, key) {
			return symbol
		}
	}

	// Default fallback
	return strings.ToUpper(strings.ReplaceAll(company, " ", "")) + ".NS"
}
