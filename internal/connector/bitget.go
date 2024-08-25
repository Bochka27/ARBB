package connector

import (
	"fmt"
	"github.com/Bochka27/ARBB/config/bitget"
)

func BitGetF() {
	client := new(bitget.MixMarketClient).Init()
	params := bitget.NewParams()
	params["productType"] = "umcbl"

	resp, err := client.Tickers(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func BitGetS() {
	client := new(bitget.SpotMarketClient).Init()
	params := bitget.NewParams()

	resp, err := client.Tickers(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
