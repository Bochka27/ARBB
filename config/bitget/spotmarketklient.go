package bitget

type SpotMarketClient struct {
	BitgetRestClient *BitgetRestClient
}

func (s *SpotMarketClient) Init() *SpotMarketClient {
	s.BitgetRestClient = new(BitgetRestClient).Init()
	return s
}

func (s *SpotMarketClient) Currencies() (string, error) {
	params := NewParams()
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/public/currencies", params)
	return resp, err
}

func (s *SpotMarketClient) Products(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/public/products", params)
	return resp, err
}

func (s *SpotMarketClient) Product(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/public/product", params)
	return resp, err
}

func (s *SpotMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/market/fills", params)
	return resp, err
}

func (s *SpotMarketClient) Depth(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/market/depth", params)
	return resp, err
}

func (s *SpotMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/market/tickers", params)
	return resp, err
}

func (s *SpotMarketClient) Ticker(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/market/ticker", params)
	return resp, err
}

func (s *SpotMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := s.BitgetRestClient.DoGet("/api/spot/v1/market/candles", params)
	return resp, err
}
