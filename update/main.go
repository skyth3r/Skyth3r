package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	const rssFeed = "https://akashgoswami.com/articles/index.xml"
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(rssFeed)
	if err != nil {
		log.Fatalf("Error getting feed: %v", err)
	}

	latestFeedItem := feed.Items[0]

	date := time.Now().Format("2 Jan 2006")

	title := `<h3 align="center">Hello! I'm Akash 👋🏽</h3>`
	socialLinks := `<br><p align="center">
						<a href="https://akashgoswami.com/">Website</a> •
						<a href="https://twitter.com/akashgoswami_">Twitter</a> •
						<a href="https://bsky.app/profile/akashgoswami.com">Bluesky</a> •
						<a href="https://hachyderm.io/@akashgoswami" rel="me">Mastodon</a>
  					</p>`
	article := `<br><p>Latest article from my website: **[` + latestFeedItem.Title + `](` + latestFeedItem.Link + `)**. You can also subscribe to my [**article RSS feed**](rssFeed).<p>`
	updated := `<br><sub>Last updated on ` + date + `.<sub>`
	data := fmt.Sprintf("%s%s%s%s", title, socialLinks, article, updated)

	file, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("Unable to create README file. Error: %v", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		log.Fatalf("Error writing data to README: %v", err)
	}
	file.Sync()
}
