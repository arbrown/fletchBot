package main

import (
	"fmt"
	"log"

	"github.com/arbrown/fletchbot/settings"
	"github.com/jzelinskie/geddit"
)

func main() {
	// Load settings from config file
	config, err := settings.ReadSettings("fletchbotsettings.json")
	if err != nil {
		fmt.Println("Some dumb error")
		log.Fatal(err)
	}

	// Launch main loop
	fletchBot(config)

}

func fletchBot(config settings.FletchBotSettings) {
	// Create new OauthSession for Reddit API
	o, err := geddit.NewOAuthSession(
		config.AppID,
		config.AppSecret,
		config.UserAgent,
		"",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = o.LoginAuth(config.UserName, config.Password)
	if err != nil {
		log.Fatal(err)
	}

	opts := geddit.ListingOptions{Limit: 4}

	posts, err := o.SubredditSubmissions("", geddit.NewSubmissions, opts)
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range posts {
		fmt.Printf("%d)\t%s (%d)\n", i, p.Title, p.Score)
		comments, err := o.Comments(p, geddit.NewSubmissions, opts)
		if err != nil {
			log.Fatal(err)
		}
		for j, c := range comments {
			fmt.Printf("\t%d\t%s (%d)\n", j, c.Body, c.Likes)
		}
	}

}
