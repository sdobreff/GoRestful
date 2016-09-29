package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func DefaultHandler(inner http.Handler, name string) http.Handler {

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

		invalidTokenMsg := "WHAT? Invalid Token? F*** off!"

		if "" == r.Header.Get("token") {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, invalidTokenMsg)
			FilePrintln(Warning, invalidTokenMsg)
			return
		}

		rawToken := r.Header.Get("token")

		token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Printf("Token for user %v is valid!", claims["name"])
			Info.Printf("Token for user %v is valid!", claims["name"])
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, invalidTokenMsg)
			FilePrintln(Warning, invalidTokenMsg)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, invalidTokenMsg)
			FilePrintln(Warning, invalidTokenMsg)
			return
		}

		inner.ServeHTTP(w, r)

		Info.Printf(
			"%s\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			name,
		)
	})
}
