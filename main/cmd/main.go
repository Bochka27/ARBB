package main

import (
	"github.com/Bochka27/ARBB/internal/connector"
)

func main() {
	//a := app.NewApp()
	//
	//if err := a.Run(); err != nil {
	//	panic(err)
	//}
	//connector.NewBybitS()
	//connector.NewBybitF()
	//connector.KuCoinF()
	connector.KuCoinS()
}
