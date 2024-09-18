package connector

const (
	//BitGet
	//Bitmex
	//ByBit
	urlBF = "https://api.bybit.com/v5/market/tickers?category=linear"
	urlBS = "https://api.bybit.com/v5/market/tickers?category=spot"
	//Gate
	urlGF = "https://api.gateio.ws/api/v4/futures/usdt/tickers"
	urlGS = "https://api.gateio.ws/api/v4/spot/tickers"
	//Htx
	urlHF  = "https://api.hbdm.com/linear-swap-ex/market/bbo?contract_type=swap&business_type=swap"
	urlHS  = "https://api.huobi.pro/market/tickers"
	urlHFR = "https://api.hbdm.com/linear-swap-api/v1/swap_batch_funding_rate"
	// KuCoin
	urlKCF  = "https://api-futures.kucoin.com/api/v1/allTickers"
	urlKCFR = "https://api-futures.kucoin.com/api/v1/contracts/active"
	urlKCS  = "https://api.kucoin.com/api/v1/market/allTickers"
	// Mexc
	// Okex
	urlOSL = "https://www.okx.com/api/v5/market/tickers?instType=SPOT"
	urlOF  = "https://www.okx.com/api/v5/market/tickers?instType=SWAP"
	urlOFR = "https://www.okx.com/api/v5/public/funding-rate?instId="
)
