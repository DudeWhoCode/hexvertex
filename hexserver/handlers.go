package hexserver

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/DudeWhoCode/hexvertex/backend"
	"path"
)


func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func RedirectUserToTwitter(w http.ResponseWriter, r *http.Request) {
	tokenUrl := fmt.Sprintf("http://%s/maketoken", r.Host)
	token, requestUrl, err := backend.C.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to save the token, we'll need it for AuthorizeToken()
	backend.Tokens[token.Token] = token
	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
}

func GetTwitterToken(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")
	accessToken, err := backend.C.AuthorizeToken(backend.Tokens[tokenKey], verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	client, err := backend.C.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Get(
		"https://api.twitter.com/1.1/statuses/user_timeline.json?count=50&screen_name=dudewhocode")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "The newest item in your home timeline is: "+string(bits))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fileName := path.Join("templates", "login.html")
	http.ServeFile(w,r,fileName)
}