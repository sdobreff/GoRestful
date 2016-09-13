package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sess")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tok := session.Values["accessToken"]

		if tok == nil && name != "Auth" {
			b := make([]byte, 16)
			rand.Read(b)

			state := base64.URLEncoding.EncodeToString(b)

			session, _ := store.Get(r, "sess")
			session.Values["state"] = state
			session.Save(r, w)

			url := oauthCfg.AuthCodeURL(state)
			http.Redirect(w, r, url, 302)

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
