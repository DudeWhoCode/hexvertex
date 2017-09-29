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

	//var consumerKey *string = flag.String(
	//	"consumerkey",
	//	"",
	//	"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")
	//
	//var consumerSecret *string = flag.String(
	//	"consumersecret",
	//	"",
	//	"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")
	//
	//var port *int = flag.Int(
	//	"port",
	//	8888,
	//	"Port to listen on.")

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