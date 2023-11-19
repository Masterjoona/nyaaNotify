package main

type NyaaPost struct {
	Category    string
	Title       string
	Magnet      string
	Torrent     string
	URL         string
	Size        string
	Date        string
	Seed        string
	Leech       string
	Completed   string
	Comments    string
	CategoryImg string
}

var url = "https://nyaa.si"

func main() {
	generateCron, discordWebhook, shortenerToken, includeString, regexString, shortenerURL, amount := ParseCommandParameters()

	if generateCron {
		MakeParameters()
		return
	}
	if ShouldntPost() {
		//println("Not posting.")
		return
	}
	CleanPosted()
	matches, nyaaPosts := MatchPosts()
	for i, match := range matches {
		if match[5] == "" {
			match[5] = "0"
		}
		nyaaPosts[i] = NyaaPost{
			Category:    match[1],
			URL:         url + match[3],
			Title:       match[4],
			Torrent:     url + match[8],
			Magnet:      match[9],
			Size:        match[10],
			Date:        match[11],
			Seed:        match[12],
			Leech:       match[13],
			Completed:   match[14],
			Comments:    match[5],
			CategoryImg: url + match[2],
		}
	}
	for _, post := range nyaaPosts {
		if CheckTitle(post.Title, includeString, regexString) {
			if CheckIfAlreadyPosted(post) {
				continue
			}
			description := MakeDescription(post, shortenerToken, shortenerURL)
			if IsOverAmount(amount) {
				//println("Over amount.")
				StorePosted(GetDate())
				return
			}
			SendEmbed(post, description, discordWebhook)
			StorePosted(post.URL)
		}
	}
}
