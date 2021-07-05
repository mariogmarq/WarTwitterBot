package view

import (
	"os"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
)

func TestUpload(t *testing.T) {
	client := CreateClient()
	media_id, err := client.postImage(os.Getenv("IMAGES_DIR") + "/test_image.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, _, err = client.client.Statuses.Update("Hello World!", &twitter.StatusUpdateParams{MediaIds: []int64{media_id}})
	if err != nil {
		t.Fatal(err.Error())
	}
}
