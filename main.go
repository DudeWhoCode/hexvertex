package main

import (
	"log"
	"github.com/DudeWhoCode/hexvertex/hexserver"
	"github.com/DudeWhoCode/hexvertex/backend"
	"net/http"
	"github.com/mrjones/oauth"
)

var C *oauth.Consumer
var Tokens map[string]*oauth.RequestToken

func main() {
	backend.NewOauth()
	rt := hexserver.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", rt))
}
