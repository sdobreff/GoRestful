package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

const (
	defaultConfigFile = "config.json"

	githubAuthorizeUrl = "https://accounts.google.com/o/oauth2/auth"
	githubTokenUrl     = "https://accounts.google.com/o/oauth2/token"
	redirectUrl        = "http://localhost:3000/auth"
)

type Config struct {
	ClientSecret string `json:"clientSecret"`
	ClientID     string `json:"clientID"`
	Secret       string `json:"secret"`
}

var (
	cfg      *Config
	oauthCfg *oauth2.Config
	store    *sessions.CookieStore

	// scopes
	scopes = []string{"https://www.googleapis.com/auth/plus.profile.emails.read"}
)

func main() {

	var err error
	cfg, err = loadConfig(defaultConfigFile)
	if err != nil {
		panic(err)
	}

	store = sessions.NewCookieStore([]byte(cfg.Secret))

	oauthCfg = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  githubAuthorizeUrl,
			TokenURL: githubTokenUrl,
		},
		RedirectURL: redirectUrl,
		Scopes:      scopes,
	}

	router := NewRouter()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))
}

func loadConfig(file string) (*Config, error) {
	var config Config

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// package main

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/sessions"
// 	"golang.org/x/oauth2"
// )

// const (
// 	defaultLayout = "templates/layout.html"
// 	templateDir   = "templates/"

// 	defaultConfigFile = "config.json"

// 	githubAuthorizeUrl = "https://accounts.google.com/o/oauth2/auth"
// 	githubTokenUrl     = "https://accounts.google.com/o/oauth2/token"
// 	redirectUrl        = "http://localhost:3000/auth"
// )

// type Config struct {
// 	ClientSecret string `json:"clientSecret"`
// 	ClientID     string `json:"clientID"`
// 	Secret       string `json:"secret"`
// }

// var (
// 	cfg      *Config
// 	oauthCfg *oauth2.Config
// 	store    *sessions.CookieStore

// 	// scopes
// 	scopes = []string{"https://www.googleapis.com/auth/plus.profile.emails.read"}

// 	tmpls = map[string]*template.Template{}
// )

// func loadConfig(file string) (*Config, error) {
// 	var config Config

// 	b, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = json.Unmarshal(b, &config); err != nil {
// 		return nil, err
// 	}

// 	return &config, nil
// }

// func main() {
// 	tmpls["home.html"] = template.Must(template.ParseFiles(templateDir+"home.html", defaultLayout))

// 	var err error
// 	cfg, err = loadConfig(defaultConfigFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	store = sessions.NewCookieStore([]byte(cfg.Secret))

// 	oauthCfg = &oauth2.Config{
// 		ClientID:     cfg.ClientID,
// 		ClientSecret: cfg.ClientSecret,
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  githubAuthorizeUrl,
// 			TokenURL: githubTokenUrl,
// 		},
// 		RedirectURL: redirectUrl,
// 		Scopes:      scopes,
// 	}

// 	r := mux.NewRouter()
// 	r.HandleFunc("/", HomeHandler)
// 	r.HandleFunc("/start", StartHandler)
// 	r.HandleFunc("/auth", CallbackHandler)

// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
// 	http.Handle("/", r)

// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "3000"
// 	}

// 	log.Fatalln(http.ListenAndServe(":"+port, nil))
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpls["home.html"].ExecuteTemplate(w, "base", map[string]interface{}{})
// }

// func StartHandler(w http.ResponseWriter, r *http.Request) {
// 	b := make([]byte, 16)
// 	rand.Read(b)

// 	state := base64.URLEncoding.EncodeToString(b)

// 	session, _ := store.Get(r, "sess")
// 	session.Values["state"] = state
// 	session.Save(r, w)

// 	url := oauthCfg.AuthCodeURL(state)
// 	http.Redirect(w, r, url, 302)
// }

// func CallbackHandler(w http.ResponseWriter, r *http.Request) {
// 	session, err := store.Get(r, "sess")
// 	if err != nil {
// 		fmt.Fprintln(w, "aborted")
// 		return
// 	}

// 	if r.URL.Query().Get("state") != session.Values["state"] {
// 		fmt.Fprintln(w, "no state match; possible csrf OR cookies not enabled")
// 		return
// 	}
// 	tkn, err := oauthCfg.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
// 	if err != nil {
// 		fmt.Fprintln(w, "there was an issue getting your token")
// 		return
// 	}

// 	if !tkn.Valid() {
// 		fmt.Fprintln(w, "retreived invalid token")
// 		return
// 	}

// 	// client := oauthCfg.Client(oauth2.NoContext, tkn)

// 	// user, err := client.Get("...")
// 	// fmt.Fprintln(w, user)
// 	// user, _, err := client.Users.Get("")
// 	// if err != nil {
// 	// 	fmt.Println(w, "error getting name")
// 	// 	return
// 	// }

// 	// session.Values["name"] = user.Email
// 	session.Values["accessToken"] = tkn.AccessToken
// 	session.Save(r, w)

// 	http.Redirect(w, r, "/", 302)
// }
