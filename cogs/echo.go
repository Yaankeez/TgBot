package cogs

import(
	"gopkg.in/telebot.v3"
)

func EchoCog(bot *telebot.Bot){
	bot.Handle(telebot.OnText, func(c telebot.Context) error{
		return c.Send("Чё он сказал? Он сказал:" + c.Text() )
	})
}