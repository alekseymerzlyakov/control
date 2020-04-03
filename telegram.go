package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func telegram(mes string) {
	bot, err := tgbotapi.NewBotAPI("1013858255:AAFfvt1SIglynSzpXG786tMlHwdRfsQ3Zrk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	//

	// Для работы надо узнать чат ID - запускаем закоментированый код - пишем в чат и получаем ID чата
	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID



		msg := tgbotapi.NewMessage(-467584122, mes)
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)

}


