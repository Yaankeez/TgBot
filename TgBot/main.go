package main
import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
	"TgBot/cogs"
)

func main(){
	godotenv.Load()
	token:=os.Getenv("BOT_TOKEN")
	if(token==""){
		log.Fatal("Gde token???")
	}
	bot_settings := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot,err := telebot.NewBot(bot_settings)
	if err != nil{
		log.Fatal(err)
	}

	cogs.StartMessageCog(bot)
	cogs.EchoCog(bot)
	cogs.ScheduleCogs(bot)

	log.Println("Bot was started")
	bot.Start()
}

//имеха старался
