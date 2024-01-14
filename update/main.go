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
	const webFeed = "https://akashgoswami.com/articles/index.xml"
	const devFeed = "https://akashgoswami.dev/posts/index.xml"
	const title = `<h2 align="center">Hello! I'm Akash 👋🏽</h2>`
	const latestPostsTitle = `<h3>Latest posts</h3>`

	webFeedItem := getLatestFeedItem(webFeed)
	devFeedItem := getLatestFeedItem(devFeed)

	date := time.Now().Format("2 Jan 2006")

	socialLinks := `<p align="center">
						<a href="https://akash.lol/" rel="me">Links</a> •
						<a href="https://akashgoswami.com/" rel="me">Website</a> •
						<a href="https://akashgoswami.dev/" rel="me">Dev blog</a> •
						<a href="https://hachyderm.io/@akashgoswami" rel="me">Mastodon</a> •
						<a href="https://bsky.app/profile/akashgoswami.com" rel="me">Bluesky</a>
  					</p>`

	article := `<p>Latest article from my website: <a href="` + webFeedItem.Link + `">` + webFeedItem.Title + `</a>. You can also subscribe to my <a href="` + webFeed + `">article RSS feed.</a></p>`
	devArticle := `<p>Latest post from my dev blog: <a href="` + devFeedItem.Link + `">` + devFeedItem.Title + `</a>. You can also subscribe to my <a href="` + devFeed + `">dev post RSS feed.</a></p>`
	updated := `<sub>Last updated on ` + date + `.<sub>`
	data := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n\n%s", title, socialLinks, latestPostsTitle, article, devArticle, updated)

	file, err := os.Create("../README.md")
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

func getLatestFeedItem(input string) *gofeed.Item {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(input)
	if err != nil {
		log.Fatalf("Error getting feed: %v", err)
	}
	return feed.Items[0]
}
