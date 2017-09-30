package backend

import (
	"github.com/mrjones/oauth"
	"flag"
	"os"
)

var Tokens map[string]*oauth.RequestToken
var C *oauth.Consumer

func NewOauth() {
	Tokens = make(map[string]*oauth.RequestToken)
	var consumerKey *string
	var consumerSecret *string
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	consumerKey = &key
	consumerSecret = &secret

	flag.Parse()

	C = oauth.NewConsumer(
		*consumerKey,
		*consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	C.Debug(true)

}