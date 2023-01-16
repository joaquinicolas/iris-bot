package api

import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
)

type Bot interface {
	Run() error
}

type TelegramBot struct {
	bot   *bt.Bot
	Token string
}

func NewTelegramBot(token string) *TelegramBot {
	up := cfg.DefaultUpdateConfigs()
	cf := cfg.BotConfigs{
		BotAPI: cfg.DefaultBotAPI,
		APIKey: token, UpdateConfigs: up,
		Webhook:        false,
		LogFileAddress: cfg.DefaultLogFile,
	}
	bot, err := bt.NewBot(&cf)
	if err == nil {
		return &TelegramBot{
			Token: token,
			bot:   bot,
		}
	}

	return nil
}

func (tb *TelegramBot) Run() error {
	return tb.bot.Run()
}
