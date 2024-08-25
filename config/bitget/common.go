package bitget

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
}

func (b *BitgetRestClient) Init() *BitgetRestClient {
	b.ApiKey = ApiKey
	b.ApiSecretKey = SecretKey
	b.BaseUrl = BaseUrl
	b.Passphrase = PASSPHRASE
	b.Signer = new(Signer).Init(SecretKey)
	b.HttpClient = http.Client{
		Timeout: time.Duration(TimeoutSecond) * time.Second,
	}
	return p
}

func (b *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := b.Signer.Sign(POST, uri, params, timesStamp)
	if RSA == SignType {
		sign = b.Signer.SignByRSA(POST, uri, params, timesStamp)
	}
	requestUrl := BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(POST, requestUrl, buffer)

	Headers(request, b.ApiKey, timesStamp, sign, b.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := b.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (b *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := TimesStamp()
	body := BuildGetParams(params)
	fmt.Println(body)

	sign := b.Signer.Sign(GET, uri, body, timesStamp)

	requestUrl := b.BaseUrl + uri + body

	request, err := http.NewRequest(GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	Headers(request, b.ApiKey, timesStamp, sign, b.Passphrase)

	response, err := b.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
