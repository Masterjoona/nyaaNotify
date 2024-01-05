package main

import "encoding/xml"

type NyaaPost struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	URL         string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Seeders     string `xml:"seeders"`
	Leechers    string `xml:"leechers"`
	Downloads   string `xml:"downloads"`
	InfoHash    string `xml:"infoHash"`
	CategoryId  string `xml:"categoryId"`
	Category    string `xml:"category"`
	Size        string `xml:"size"`
	Comments    string `xml:"comments"`
	Trusted     string `xml:"trusted"`
	Remake      string `xml:"remake"`
	Description string `xml:"description"` // This is not the actual description.
}

type Rss struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	AtomLink    string `xml:"atom:link"`
	Channel     struct {
		NyaaPosts []NyaaPost `xml:"item"`
	} `xml:"channel"`
}

// Copilot notified me of similiar code in https://github.com/Zhousiru/NyaaHub/blob/aa12f6a5fc9326bbd7bf717608bd74da3dfa530f/scheduler/lib/rss/model.go
// Their project is really cool and better than mine in many ways, so check it out!

func GetNyaaPosts() []NyaaPost {
	content := FetchNyaa(Url + "/?page=rss")
	rss := parseXML(content)
	posts := rss.getPosts()
	return posts
}

func parseXML(content string) *Rss {
	rss := &Rss{}
	xml.Unmarshal([]byte(content), rss)
	return rss
}

func (rss *Rss) getPosts() []NyaaPost {
	posts := []NyaaPost{}
	posts = append(posts, rss.Channel.NyaaPosts...)
	return posts
}
