package config

import (
	"arknights_bot/utils/telebot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

var Arknights *tgbotapi.BotAPI

var TeleBot *telebot.Bot

func Bot() error {
	token := viper.GetString("bot.token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println(err)
		return err
	}
	Arknights = bot
	log.Println("机器人初始化完成")
	return nil
}
