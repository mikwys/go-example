package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/mikwys/go-example/internal"
)

func main() {
	indexTpl := template.Must(template.ParseFiles("web/index.gohtml"))
	timelineTpl := template.Must(template.ParseFiles("web/timeline.gohtml"))
	config, err := internal.LoadEnvConfig()
	if err != nil {
		panic(err)
	}

	// setup dependencies
	twitterClient := internal.NewM3OTwitterClient(config.Token())

	http.HandleFunc("/", NewIndexHandler(indexTpl))
	http.HandleFunc("/timeline", NewTweetListHandler(timelineTpl, twitterClient))
	http.HandleFunc("/timeline-json", NewTweetListHandlerJSON(twitterClient))

	log.Println(fmt.Sprintf("HTTP Listener at: %s", config.Port()))
	if err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port()), nil); err != http.ErrServerClosed {
		log.Printf("unexpected http error %v", err)
	}
	log.Println("bye!")
}

func NewIndexHandler(tpl *template.Template) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		err := tpl.Execute(w, nil)
		if err != nil {
			log.Println(fmt.Sprintf("could execute template %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func NewTweetListHandler(tpl *template.Template, twSvc internal.TwitterClient) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userName := req.URL.Query().Get("user")
		log.Println("user param", userName)
		if strings.TrimSpace(userName) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, err := twSvc.Load(userName)
		if err != nil {
			log.Println(fmt.Sprintf("could not load user %s details/tweets. err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, internal.UserTimelineModel{
			Details:  user.Details,
			Timeline: user.Timeline,
		})
		if err != nil {
			log.Println(fmt.Sprintf("could not execute template (user %s) err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
func NewTweetListHandlerJSON(twSvc internal.TwitterClient) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userName := req.URL.Query().Get("user")
		log.Println("user param", userName)
		if strings.TrimSpace(userName) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, err := twSvc.Load(userName)
		if err != nil {
			log.Println(fmt.Sprintf("could not load user %s details/tweets. err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
