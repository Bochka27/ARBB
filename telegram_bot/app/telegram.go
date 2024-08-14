package app

import (
	"gopkg.in/telebot.v3"
	"time"
)

type Telegram struct {
	bot      *telebot.Bot
	commands *telebot.Command
}

func NewTelegram(token string) (*Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Millisecond},
	})
	return &Telegram{bot: bot}, err
}

func (t *Telegram) Start() error {
	t.bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello, World!")
	})
	t.bot.Handle("/help", func(c telebot.Context) error {
		return c.Send("Available commands:\n/start - Start the bot" +
			"\n/help - Get help information" +
			"\n/rules - Get rules")
	})
	t.bot.Handle("/rules", func(c telebot.Context) error {
		return c.Send("There is no rules!")
	})
	t.bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Reply(c.Message().Text)
	})
	t.bot.Handle(telebot.OnText, t.sHandler)
	t.bot.Start()
	return nil
}

func (t *Telegram) sHandler(c telebot.Context) error {
	return nil
}
