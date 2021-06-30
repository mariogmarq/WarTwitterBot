// Publish the tweets
package view

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type View struct {
	client *twitter.Client
}

func CreateClient() View {
	config := oauth1.NewConfig("Consumerkey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return View{client}
}

// Post a kill event
func (v View) PostUpdate() {

}
