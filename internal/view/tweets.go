// Publish the tweets
package view

import (
	"context"
	"github/mariogmarq/WarTwitterBot/internal/models"
	"github/mariogmarq/WarTwitterBot/internal/repository"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type View struct {
	client *twitter.Client
}

func CreateClient() View {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CONSUMER_KEY"),
		ClientSecret: os.Getenv("CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	log.Println(config.ClientID)
	log.Println(config.ClientSecret)

	httpClient := config.Client(context.Background())

	client := twitter.NewClient(httpClient)

	return View{client}
}

// Post an event
func (v View) postUpdate(phrase string) {
	log.Printf("Posting %s", phrase)
	_, _, err := v.client.Statuses.Update(phrase, nil)

	for err != nil {
		log.Println(err.Error())
		time.Sleep(time.Minute * 2)
		_, _, err = v.client.Statuses.Update(phrase, nil)
	}
}

func getPhraseKill(apis []models.FighterApi) string {
	phrase, err := repository.GetInstance().GetPhraseByN(len(apis))
	if err != nil {
		panic(err)
	}

	sentence, err := phrase.MapToPlayers(apis...)
	if err != nil {
		panic(err)
	}

	return sentence
}

func (v View) PostKillUpdate(apis []models.FighterApi) {
	phrase := getPhraseKill(apis)
	v.postUpdate(phrase)
}

func (v View) PostWinUpdate(apis []models.FighterApi) {
	var phrase *models.Phrase
	if len(apis) == 1 {
		phrase, _ = models.NewPhrase("$1 ha ganado")
	} else {
		phrase, _ = models.NewPhrase("$1 y $2 han ganado")
	}

	sentence, _ := phrase.MapToPlayers(apis...)
	v.postUpdate(sentence)
}
