package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Logger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// session, err := store.Get(r, "sess")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// tok := session.Values["accessToken"]

		// if tok == nil && name != "Auth" {
		// 	b := make([]byte, 16)
		// 	rand.Read(b)

		// 	state := base64.URLEncoding.EncodeToString(b)

		// 	session, _ := store.Get(r, "sess")
		// 	session.Values["state"] = state
		// 	session.Save(r, w)

		// 	url := oauthCfg.AuthCodeURL(state)
		// 	http.Redirect(w, r, url, 302)

		// }

		if "" == r.Header.Get("token") {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return
		}

		rawToken := r.Header.Get("token")

		token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Printf("Token for user %v is valid!", claims["name"])
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return
		}

		start := time.Now()
		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
