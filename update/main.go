package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	const webFeed = "https://akashgoswami.com/articles/index.xml"
	const webNotesFeed = "https://akashgoswami.com/notes/index.xml"
	const devFeed = "https://akashgoswami.dev/posts/index.xml"
	const devNotesFeed = "https://akashgoswami.dev/notes/index.xml"
	const title = `<h1 align="center">Hello! I'm Akash 👋🏽</h1>`
	const aboutText = `<p>I'm an experienced software engineer specialising in platform security and backend systems, with 7+ years of experience across fintech and developer security tooling. I currently use Go, Python and TypeScript and Postgres for development.</p>`
	const latestPostsTitle = `<h2>📬 Latest posts</h2>`
	const latestNotesTitle = `<h2>📝 Latest notes</h2>`
	const techStackTitle = `<h2>💻 Tech Stack</h2>`
	const techStackText = `![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) ![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white) ![GraphQL](https://img.shields.io/badge/-GraphQL-E10098?style=for-the-badge&logo=graphql&logoColor=white) ![ApacheCassandra](https://img.shields.io/badge/cassandra-%231287B1.svg?style=for-the-badge&logo=apache-cassandra&logoColor=white) ![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-%2300000f.svg?style=for-the-badge&logo=mysql&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white) ![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white) ![Google Cloud](https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white) ![Cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white) ![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white) ![Hugo](https://img.shields.io/badge/Hugo-black.svg?style=for-the-badge&logo=Hugo) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) ![Figma](https://img.shields.io/badge/figma-%23F24E1E.svg?style=for-the-badge&logo=figma&logoColor=white)`
	const githubStatsTitle = `<h2>📊 GitHub Stats</h2>`
	const githubStatsText = `![](https://github-readme-stats.vercel.app/api?username=Skyth3r&theme=city_light&hide_border=false&include_all_commits=false&count_private=true)<br/>![](https://github-readme-streak-stats.herokuapp.com/?user=Skyth3r&theme=city_light&hide_border=false)`

	socialLinks := `<p align="center">
	<a href="https://akash.lol/" rel="me">Links</a> •
	<a href="https://akashgoswami.com/" rel="me">Website</a> •
	<a href="https://akashgoswami.dev/" rel="me">Dev blog</a> •
	<a href="https://bsky.app/profile/akashgoswami.com" rel="me">Bluesky</a> •
	<a href="https://hachyderm.io/@akashgoswami" rel="me">Mastodon</a>
  	</p>`

	posts := mergeAndSortItems(getFeedItems(webFeed), getFeedItems(devFeed), 3)
	notes := mergeAndSortItems(getFeedItems(webNotesFeed), getFeedItems(devNotesFeed), 3)

	postsSection := buildItemList(posts)
	notesSection := buildItemList(notes)

	date := time.Now().Format("2 Jan 2006")
	updated := `<sub>Last updated on ` + date + `.<sub>`

	data := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n%s\n%s\n\n%s\n\n%s",
		title, aboutText, socialLinks,
		latestPostsTitle, postsSection,
		latestNotesTitle, notesSection,
		techStackTitle, techStackText,
		githubStatsTitle, githubStatsText,
		updated,
	)

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

func getFeedItems(url string) []*gofeed.Item {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(url)
	if err != nil {
		log.Fatalf("Error getting feed: %v", err)
	}
	return feed.Items
}

func itemTime(item *gofeed.Item) time.Time {
	if item.PublishedParsed != nil {
		return *item.PublishedParsed
	}
	if item.UpdatedParsed != nil {
		return *item.UpdatedParsed
	}
	return time.Time{}
}

func mergeAndSortItems(a, b []*gofeed.Item, n int) []*gofeed.Item {
	merged := append(a, b...)
	sort.Slice(merged, func(i, j int) bool {
		return itemTime(merged[i]).After(itemTime(merged[j]))
	})
	if len(merged) > n {
		merged = merged[:n]
	}
	return merged
}

func buildItemList(items []*gofeed.Item) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString("\n- [" + item.Title + "](" + item.Link + ")\n")
	}
	return sb.String()
}
