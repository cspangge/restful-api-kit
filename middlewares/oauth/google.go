package oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	tools "restful-api-kit/utilities"
)

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
}

const oauthStateString = "random"

var googleOauthConfig = &oauth2.Config{
	ClientID:     "83107980591-tvaau8j9u8k2r3a63lk37b0cvs9aktrg.apps.googleusercontent.com",
	ClientSecret: "1uewvzFTtJrSXeeNoht5QYfC",
	RedirectURL:  "https://localhost:8000",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: endpoint,
}

func GoogleAuth(c *gin.Context) {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).

	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	state := c.Request.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(state)

	code := c.Request.FormValue("code")
	fmt.Println(code)
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	fmt.Println(token)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n\n", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			tools.CheckErr(err)
		}
	}(response.Body)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		tools.CheckErr(err)
	}
	_, err2 := fmt.Fprintf(c.Writer, "Content: %s\n", contents)
	if err2 != nil {
		tools.CheckErr(err2)
	}
}