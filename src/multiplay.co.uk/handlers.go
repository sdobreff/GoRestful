package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
)

const serverURL string = "http://private-e7633-cfe1.apiary-mock.com/"

type Regions []struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Region struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Providers []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DataCenters []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Machines []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Crashes []struct {
	ID        int    `json:"id"`
	MachineID int    `json:"machine_id"`
	Body      string `json:"body"`
}

type PSUs struct {
	Columns []string `json:"columns"`
	Values  []struct {
		Num0 int64 `json:"0"`
		Num1 int   `json:"1"`
	} `json:"values"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://private-e7633-cfe1.apiary-mock.com/regions")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Regions
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s[2].ID)
}

func AllRegions(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(serverURL + "regions")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Regions
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func RegionShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Region
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func AllProviders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id + "/providers")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Providers
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func AllDataCenters(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id + "/datacenters")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s DataCenters
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func AllMachines(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id + "/machines")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Machines
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func AllCrashes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id + "/crashes")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Crashes
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func AllPSUs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := http.Get(serverURL + "region/" + id + "/psu")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var s PSUs
	json.Unmarshal(body, &s)
	fmt.Fprintf(w, "%v", s)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "sess")
	if err != nil {
		fmt.Fprintln(w, "aborted")
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		fmt.Fprintln(w, "no state match; possible csrf OR cookies not enabled")
		return
	}
	tkn, err := oauthCfg.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	if err != nil {
		fmt.Fprintln(w, "there was an issue getting your token")
		return
	}

	if !tkn.Valid() {
		fmt.Fprintln(w, "retreived invalid token")
		return
	}

	// client := oauthCfg.Client(oauth2.NoContext, tkn)

	// user, err := client.Get("...")
	// fmt.Fprintln(w, user)
	// user, _, err := client.Users.Get("")
	// if err != nil {
	// 	fmt.Println(w, "error getting name")
	// 	return
	// }

	// session.Values["name"] = user.Email
	session.Values["accessToken"] = tkn.AccessToken
	session.Save(r, w)

	http.Redirect(w, r, "/", 302)
}
