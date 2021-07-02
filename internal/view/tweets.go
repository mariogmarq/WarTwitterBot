// Publish the tweets
package view

import (
	"os"

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

// Post a kill event
func (v View) PostUpdate() {

}
