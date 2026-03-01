package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
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

	// 4. Get real stock price - try Yahoo first, then Alpha Vantage
	stockData, err := r.GetStockPrice(analysis.StockSymbol)
	if err == nil {
		analysis.CurrentPrice = stockData.Price
	} else {
		// Try Alpha Vantage as fallback
		price, err2 := r.GetAlphaVantagePrice(analysis.StockSymbol)
		if err2 == nil {
			analysis.CurrentPrice = price
		} else {
			// Log both errors but continue with $0 price
			fmt.Printf("Failed to get stock price for %s: Yahoo=%v, AlphaVantage=%v\n",
				analysis.StockSymbol, err, err2)
		}
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

	resp, err := r.httpClient.Get(url) // #nosec G107 G704
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
		return 0, fmt.Errorf("no news articles found for company: %s", company)
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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add User-Agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := r.httpClient.Do(req) // #nosec G107 G704
	if err != nil {
		return nil, fmt.Errorf("yahoo finance request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("yahoo finance returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	chart, ok := data["chart"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid chart data")
	}

	result, ok := chart["result"].([]interface{})
	if !ok || len(result) == 0 {
		return nil, fmt.Errorf("no result data")
	}

	resultData, ok := result[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid result format")
	}

	meta, ok := resultData["meta"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no meta data")
	}

	priceVal, ok := meta["regularMarketPrice"]
	if !ok {
		return nil, fmt.Errorf("no price data")
	}

	price, ok := priceVal.(float64)
	if !ok {
		return nil, fmt.Errorf("price is not a number")
	}

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
	} else {
		// HOLD: Small price movement based on ESG and sentiment
		// Range: -5% to +10% based on score
		priceChange = (esgScore-5.0)*2.0 + (sentiment-0.5)*8.0
		targetPrice = currentPrice * (1 + priceChange/100)
	}

	return signal, targetPrice, priceChange
}

// Assess financing risk
func (r *RealTimeAgents) AssessRisk(esgScore, sentiment float64) (string, []string) {
	reasons := []string{}
	positiveReasons := []string{}

	// Negative factors
	if esgScore < 4.0 {
		reasons = append(reasons, "Low ESG score indicates sustainability risks")
	}
	if sentiment < 0.4 {
		reasons = append(reasons, "Negative news sentiment")
	}
	if esgScore < 3.0 {
		reasons = append(reasons, "High regulatory compliance risk")
	}

	// Positive factors (based on actual scores)
	if esgScore >= 7.0 {
		positiveReasons = append(positiveReasons, "Excellent ESG score ("+fmt.Sprintf("%.1f", esgScore)+"/10)")
	} else if esgScore >= 5.5 {
		positiveReasons = append(positiveReasons, "Good ESG score ("+fmt.Sprintf("%.1f", esgScore)+"/10)")
	} else if esgScore >= 4.0 {
		positiveReasons = append(positiveReasons, "Moderate ESG score ("+fmt.Sprintf("%.1f", esgScore)+"/10)")
	}

	if sentiment >= 0.65 {
		positiveReasons = append(positiveReasons, "Strong positive news sentiment")
	} else if sentiment >= 0.5 {
		positiveReasons = append(positiveReasons, "Neutral to positive news sentiment")
	} else if sentiment >= 0.4 {
		positiveReasons = append(positiveReasons, "Mixed news sentiment")
	}

	// Determine action
	action := "APPROVE"
	if len(reasons) >= 2 {
		action = "REJECT"
		return action, reasons
	} else if len(reasons) == 1 {
		action = "REVIEW"
		return action, append(reasons, positiveReasons...)
	}

	// No negative reasons - show positive analysis
	return action, positiveReasons
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

// Get stock price from Alpha Vantage (fallback)
func (r *RealTimeAgents) GetAlphaVantagePrice(symbol string) (float64, error) {
	apiKey := os.Getenv("ALPHA_VANTAGE_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("ALPHA_VANTAGE_KEY not set")
	}

	// Alpha Vantage uses different symbols (no .NS suffix for Indian stocks)
	cleanSymbol := strings.TrimSuffix(symbol, ".NS")

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s",
		cleanSymbol, apiKey)

	resp, err := r.httpClient.Get(url) // #nosec G107 G704
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	quote, ok := data["Global Quote"].(map[string]interface{})
	if !ok || len(quote) == 0 {
		return 0, fmt.Errorf("no quote data")
	}

	priceStr, ok := quote["05. price"].(string)
	if !ok {
		return 0, fmt.Errorf("no price field")
	}

	price := 0.0
	_, err = fmt.Sscanf(priceStr, "%f", &price)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price: %w", err)
	}

	if price == 0 {
		return 0, fmt.Errorf("invalid price")
	}

	return price, nil
}

// GetHistoricalPrices gets historical stock prices for calculating returns
func (r *RealTimeAgents) GetHistoricalPrices(symbol string) (map[string]float64, error) {
	apiKey := os.Getenv("ALPHA_VANTAGE_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ALPHA_VANTAGE_KEY not set")
	}

	// Remove .NS suffix for Alpha Vantage
	cleanSymbol := strings.TrimSuffix(symbol, ".NS")

	// Get daily time series (last 100 days)
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s",
		cleanSymbol, apiKey)

	resp, err := r.httpClient.Get(url) // #nosec G107 G704
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	timeSeries, ok := data["Time Series (Daily)"].(map[string]interface{})
	if !ok || len(timeSeries) == 0 {
		return nil, fmt.Errorf("no historical data available")
	}

	// Extract prices by date
	prices := make(map[string]float64)
	for date, dayData := range timeSeries {
		dayMap, ok := dayData.(map[string]interface{})
		if !ok {
			continue
		}
		closeStr, ok := dayMap["4. close"].(string)
		if !ok {
			continue
		}
		price := 0.0
		_, err = fmt.Sscanf(closeStr, "%f", &price)
		if err == nil && price > 0 {
			prices[date] = price
		}
	}

	return prices, nil
}

// CalculateHistoricalReturns calculates REAL returns for different time periods with actual dates
func (r *RealTimeAgents) CalculateHistoricalReturns(symbol string, currentPrice float64) []map[string]interface{} {
	prices, err := r.GetHistoricalPrices(symbol)
	if err != nil || len(prices) == 0 {
		return []map[string]interface{}{}
	}

	now := time.Now()
	returns := []map[string]interface{}{}

	// Define periods to check (looking back in time for REAL data)
	periods := []struct {
		name   string
		months int
	}{
		{"1 Month", 1},
		{"3 Months", 3},
		{"6 Months", 6},
		{"1 Year", 12},
		{"2 Years", 24},
		{"5 Years", 60},
	}

	for _, period := range periods {
		targetDate := now.AddDate(0, -period.months, 0)

		// Find closest available date in PAST (not future)
		var closestDate string
		var closestPrice float64
		minDiff := time.Hour * 24 * 365 * 10 // 10 years max

		for dateStr, price := range prices {
			priceDate, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				continue
			}

			// Only consider dates BEFORE or ON the target date
			if priceDate.After(targetDate) {
				continue
			}

			diff := targetDate.Sub(priceDate)
			if diff < 0 {
				diff = -diff
			}
			if diff < minDiff {
				minDiff = diff
				closestDate = dateStr
				closestPrice = price
			}
		}

		// Only add if we found valid historical data
		if closestPrice > 0 && closestDate != "" {
			// Calculate ACTUAL return from past to now
			returnAmount := currentPrice - closestPrice
			returnPct := (returnAmount / closestPrice) * 100

			returns = append(returns, map[string]interface{}{
				"period":        period.name,
				"start_date":    closestDate,
				"end_date":      now.Format("2006-01-02"),
				"start_price":   closestPrice,
				"end_price":     currentPrice,
				"return_amount": returnAmount,
				"return_pct":    returnPct,
				"is_positive":   returnPct >= 0,
			})
		}
	}

	return returns
}

// CalculateInvestmentProjections calculates future value based on historical average returns
func (r *RealTimeAgents) CalculateInvestmentProjections(symbol string, currentPrice float64, historicalReturns []map[string]interface{}) []map[string]interface{} {
	projections := []map[string]interface{}{}

	// Calculate average annual return from REAL historical data
	var totalReturn float64
	var count int
	for _, ret := range historicalReturns {
		if returnPct, ok := ret["return_pct"].(float64); ok {
			totalReturn += returnPct
			count++
		}
	}

	// Use historical average or conservative 8% if no data
	annualReturnRate := 0.08
	avgReturnPct := 8.0
	if count > 0 {
		avgReturnPct = totalReturn / float64(count)
		annualReturnRate = avgReturnPct / 100.0
	}

	now := time.Now()

	// Define projection periods
	periods := []struct {
		name   string
		months int
	}{
		{"1 Month", 1},
		{"3 Months", 3},
		{"6 Months", 6},
		{"1 Year", 12},
		{"2 Years", 24},
		{"5 Years", 60},
	}

	for _, period := range periods {
		// Calculate compound return based on historical average
		years := float64(period.months) / 12.0
		futurePrice := currentPrice * math.Pow(1+annualReturnRate, years)
		returnAmount := futurePrice - currentPrice
		returnPct := (returnAmount / currentPrice) * 100

		futureDate := now.AddDate(0, period.months, 0)

		projections = append(projections, map[string]interface{}{
			"period":        period.name,
			"months":        period.months,
			"start_date":    now.Format("2006-01-02"),
			"end_date":      futureDate.Format("2006-01-02"),
			"current_price": currentPrice,
			"future_price":  futurePrice,
			"return_amount": returnAmount,
			"return_pct":    returnPct,
			"is_positive":   returnPct >= 0,
			"based_on":      fmt.Sprintf("Historical avg: %.1f%% annual return", avgReturnPct),
		})
	}

	return projections
}
