package mexc

type DetailInformation struct {
	Success bool         `json:"success"`
	Code    int64        `json:"code"`
	Data    []DataEntity `json:"data"`
}

type DataEntity struct {
	Symbol                     string   `json:"symbol"`
	DisplayName                string   `json:"displayName"`
	DisplayNameEn              string   `json:"displayNameEn"`
	PositionOpenType           int64    `json:"positionOpenType"`
	BaseCoin                   string   `json:"baseCoin"`
	QuoteCoin                  string   `json:"quoteCoin"`
	SettleCoin                 string   `json:"settleCoin"`
	ContractSize               float64  `json:"contractSize"`
	MinLeverage                int64    `json:"minLeverage"`
	MaxLeverage                int64    `json:"maxLeverage"`
	PriceScale                 int64    `json:"priceScale"`
	VolScale                   int64    `json:"volScale"`
	AmountScale                int64    `json:"amountScale"`
	PriceUnit                  float64  `json:"priceUnit"`
	VolUnit                    int64    `json:"volUnit"`
	MinVol                     int64    `json:"minVol"`
	MaxVol                     int64    `json:"maxVol"`
	BidLimitPriceRate          float64  `json:"bidLimitPriceRate"`
	AskLimitPriceRate          float64  `json:"askLimitPriceRate"`
	TakerFeeRate               float64  `json:"takerFeeRate"`
	MakerFeeRate               float64  `json:"makerFeeRate"`
	MaintenanceMarginRate      float64  `json:"maintenanceMarginRate"`
	InitialMarginRate          float64  `json:"initialMarginRate"`
	RiskBaseVol                int64    `json:"riskBaseVol"`
	RiskIncrVol                int64    `json:"riskIncrVol"`
	RiskIncrMmr                float64  `json:"riskIncrMmr"`
	RiskIncrImr                float64  `json:"riskIncrImr"`
	RiskLevelLimit             int64    `json:"riskLevelLimit"`
	PriceCoefficientVariation  float64  `json:"priceCoefficientVariation"`
	IndexOrigin                []string `json:"indexOrigin"`
	State                      int64    `json:"state"`
	IsNew                      bool     `json:"isNew"`
	IsHot                      bool     `json:"isHot"`
	IsHidden                   bool     `json:"isHidden"`
	ConceptPlate               []string `json:"conceptPlate"`
	RiskLimitType              string   `json:"riskLimitType"`
	MaxNumOrders               []int64  `json:"maxNumOrders"`
	MarketOrderMaxLevel        int64    `json:"marketOrderMaxLevel"`
	MarketOrderPriceLimitRate1 float64  `json:"marketOrderPriceLimitRate1"`
	MarketOrderPriceLimitRate2 float64  `json:"marketOrderPriceLimitRate2"`
	TriggerProtect             float64  `json:"triggerProtect"`
	Appraisal                  int64    `json:"appraisal"`
	ShowAppraisalCountdown     int64    `json:"showAppraisalCountdown"`
	AutomaticDelivery          int64    `json:"automaticDelivery"`
	ApiAllowed                 bool     `json:"apiAllowed"`
}
