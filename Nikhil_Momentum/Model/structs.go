package model

type WeeklyStatus struct {
	WeekID   string
	Date     string
	Value    int
	Holdings []Stock
	Actions  []Transaction
}

type Stock struct {
	Symbol       string
	BuyPrice     int
	Quantity     int
	AveragePrice int
}

type Transaction struct {
	Symbol string
	Price  int
	Action string
	Date   string
}

type Porfolio struct {
	Holdings []Stock
}

type StockPrices struct {
	Index    string
	Name     string
	Industry string
	Symbol   string
	MCap     string
	// Date     string
	Prices []map[int]string
}

type RawData struct {
	AllRows []StockPrices
}
