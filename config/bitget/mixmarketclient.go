package bitget

type MixMarketClient struct {
	BitgetRestClient *BitgetRestClient
}

func (m *MixMarketClient) Init() *MixMarketClient {
	m.BitgetRestClient = new(BitgetRestClient).Init()
	return p
}

func (m *MixMarketClient) Contracts(params map[string]string) (string, error) {
	resp, err := m.BitgetRestClient.DoGet("/api/mix/v1/market/contracts", params)
	return resp, err
}

func (m *MixMarketClient) Depth(params map[string]string) (string, error) {
	resp, err := m.BitgetRestClient.DoGet("/api/mix/v1/market/depth", params)
	return resp, err
}

func (m *MixMarketClient) Ticker(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/ticker", params)
	return resp, err
}

func (m *MixMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := m.BitgetRestClient.DoGet("/api/mix/v1/market/tickers", params)
	return resp, err
}

func (m *MixMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := m.BitgetRestClient.DoGet("/api/mix/v1/market/fills", params)
	return resp, err
}

func (m *MixMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := m.BitgetRestClient.DoGet("/api/mix/v1/market/candles", params)
	return resp, err
}
