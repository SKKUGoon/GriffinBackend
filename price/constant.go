package price

const (
	// base point
	BINANCE = "https://api.binance.com/api/v3"
	// request endpoint
	TRADE = "/trades"
	// needed currency
	ETHEREUM = "ETHUSDT"
	POLYGON  = "MATICUSDT"
	USDC     = "USDCUSDT"
)

type BinanceRequest struct {
	Symbol string
	Limit  string
}

type BinanceResponse struct {
	Id            int    `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty`
	UnixTime      int    `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	IsBestMatch   bool   `json:"isBestMatch"`
}

type PriceInformation struct {
	Ethereum float64 `json:"ethusdt"`
	Polygon  float64 `json:"maticusdt"`
	USDC     float64 `json:"usdcusdt"`
}

type MonthlyPayHistory struct {
	Date     string `json:"date"`
	Ethereum int    `json:"eth"`
	Polygon  int    `json:"matic"`
	USDC     int    `json:"usdc"`
}
