// Publish the tweets
package view

import (
	"testing"
)

func TestCreateClient(t *testing.T) {
	wrapper := CreateClient()
	client := wrapper.client

	_, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

}
