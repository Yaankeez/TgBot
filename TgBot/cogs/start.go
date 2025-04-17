package cogs

import(
	"gopkg.in/telebot.v3"
)

func StartMessageCog(bot *telebot.Bot){
	bot.Handle("/start", func(c telebot.Context) error{
		return c.Send("Egor Kundenko privet")
	})
}