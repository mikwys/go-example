package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/mikwys/go-example/internal"
	"go.m3o.com/twitter"
)

func main() {
	indexTpl := template.Must(template.ParseFiles("web/index.gohtml"))
	timelineTpl := template.Must(template.ParseFiles("web/timeline.gohtml"))
	config, err := internal.LoadEnvConfig()
	if err != nil {
		panic(err)
	}

	// setup dependencies
	twitterService := twitter.NewTwitterService(config.Token())

	http.HandleFunc("/", NewIndexHandler(indexTpl))
	http.HandleFunc("/timeline", NewTweetListHandler(timelineTpl, twitterService))

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

func NewTweetListHandler(tpl *template.Template, twSvc *twitter.TwitterService) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userName := req.URL.Query().Get("user")
		log.Println("user param", userName)
		if strings.TrimSpace(userName) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// load user details
		userDetails, err := twSvc.User(&twitter.UserRequest{
			Username: userName,
		})
		if err != nil {
			log.Println(fmt.Sprintf("could not load user %s details. err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// load recent tweets
		timeline, err := twSvc.Timeline(&twitter.TimelineRequest{
			Username: userName,
		})
		if err != nil {
			log.Println(fmt.Sprintf("could not load user %s tweets. err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, internal.UserTimelineModel{
			Details:  *userDetails,
			Timeline: *timeline,
		})
		if err != nil {
			log.Println(fmt.Sprintf("could execute template (user %s) err= %v", userName, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
