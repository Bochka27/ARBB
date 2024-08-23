package connector

import (
	"fmt"
	"github.com/Kucoin/kucoin-go-sdk"
)

//func KuCoinF() {
//f := kumex.NewApiServiceFromEnv()
//f.ActiveContracts()
//}

func KuCoinS() {
	s := kucoin.NewApiServiceFromEnv()
	s.Markets()
	s.Tickers()
	fmt.Println(s.Tickers())
}
