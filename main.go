package main

import (
	"github/mariogmarq/WarTwitterBot/internal/view"
	"github/mariogmarq/WarTwitterBot/internal/warlogic"
)

func main() {
	messages, win := warlogic.StartGame()
	twitterClient := view.CreateClient()

	for message := range messages {
		twitterClient.PostKillUpdate(message)
	}

	twitterClient.PostWinUpdate(<-win)
}
