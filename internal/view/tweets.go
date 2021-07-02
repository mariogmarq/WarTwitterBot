// Publish the tweets
package view

import (
	"github/mariogmarq/WarTwitterBot/internal/models"
	"github/mariogmarq/WarTwitterBot/internal/repository"
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
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return View{client}
}

// Post an event
func (v View) postUpdate(phrase string) {
	_, _, err := v.client.Statuses.Update(phrase, nil)

	for err != nil {
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
