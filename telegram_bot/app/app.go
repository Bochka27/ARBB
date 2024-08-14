package app

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	telegram, err := NewTelegram("6718620839:AAHqGIVDWguaWMJA_Om5hIJxcAK1kNBPSr4")
	if err != nil {
		return err
	}
	return telegram.Start()
}
