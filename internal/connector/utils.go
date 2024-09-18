package connector

import (
	"fmt"
	"io"
	"net/http"
)

func DoGet(url string) []byte {
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return body
}
