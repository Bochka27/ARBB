package main

import (
	"github.com/Bochka27/ARBB/telegram_bot/app"
)

func main() {
	a := app.NewApp()

	if err := a.Run(); err != nil {
		panic(err)
	}

	//client := bybit_connector.NewBybitHttpClient("1dhAvur2BYKqatGLuL", "VtkRNaoJZHQMZylFw9NmoYMTI7834hMA13Vy", bybit_connector.WithBaseURL(bybit_connector.MAINNET))
	//orderResult := client.NewTickersService()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(bybit_connector.PrettyPrint(orderResult))
}
