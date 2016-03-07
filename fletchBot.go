package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/arbrown/fletchbot/settings"
	"github.com/jzelinskie/geddit"
)

func main() {
	// Load settings from config file
	config, err := settings.ReadSettings("fletchbotsettings.json")
	if err != nil {
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

	// prepare quotes map
	quotes := make(map[*regex.Regexp]string)
	for k, s := range config.CommentQuotes {
		quotes[regexp.Compile(k)] = s
	}

	// prepare comments set (comments already replied to)
	commentsSeen := make(map[string]bool)

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
			for reg, resp := range quotes {
				if reg.MatchString(c.Body) {
					fmt.Printf("Found a match!  I should reply %s here\n", resp)
				}
			}
		}
	}

}
