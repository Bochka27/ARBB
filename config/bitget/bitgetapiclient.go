package bitget

type BitgetApiClient struct {
	BitgetRestClient *BitgetRestClient
}

func (p *BitgetApiClient) Init() *BitgetApiClient {
	p.BitgetRestClient = new(BitgetRestClient).Init()
	return p
}

func (p *BitgetApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}
