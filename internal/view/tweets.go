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
	"github.com/dghubble/oauth1"
)

type View struct {
	client *twitter.Client
}

func CreateClient() View {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	httpClient := config.Client(context.TODO(), oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET")))

	client := twitter.NewClient(httpClient)

	return View{client}
}

// Post an event
// TODO Add picture of player
func (v View) postUpdate(phrase string) {
	log.Printf("Posting %s", phrase)
	_, _, err := v.client.Statuses.Update(phrase, nil)

	for err != nil {
		log.Println(err.Error())
		time.Sleep(time.Minute * 2)
		_, _, err = v.client.Statuses.Update(phrase, nil)
	}
}

func getPhraseKill(apis []*models.FighterApi) string {
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

func (v View) PostKillUpdate(apis []*models.FighterApi) {
	phrase := getPhraseKill(apis)
	v.postUpdate(phrase)
}

func (v View) PostWinUpdate(api *models.FighterApi) {
	var phrase *models.Phrase
	phrase, _ = models.NewPhrase("$1 ha ganado")

	sentence, _ := phrase.MapToPlayers(api)
	v.postUpdate(sentence)
}
