package main

import (
	"log"
	"mevbot/bot"
	"mevbot/config"
)

func main() {

	env := config.GetEnv()
	inToken := env.GetBaoTokenAddr()
	outToken := env.GetLocTokenAddr()
	dexAddress := env.GetDexAddr()

	bot, err := bot.NewMevBot(inToken, outToken, dexAddress)
	if err != nil {
		log.Fatal("Error create bot")
	}
	bot.Run()

}
