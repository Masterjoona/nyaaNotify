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
	generateCron, discordWebhook, includeString, regexString, amount := ParseCommandParameters()

	if generateCron {
		MakeParameters()
		return
	}
	today := GetDate()
	if GetField("CleanDate") == today || (GetField("LastMod") != today && GetField("LastMod") != "") {
		Logger("Seems like a week has passed, cleaning JSON and log file.")
		CleanFiles()
	}

	if IsAmount(amount) {
		Logger("Posted already enough today.")
		SetField("LastMod", today)
		SetField("CleanDate", GetCleanDateString())
		return
	}

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
		title := post.Title
		if MatchTitle(title, includeString, regexString) {
			Logger("Found match: " + title)
			url := post.URL
			if AlreadyPosted(url) {
				Logger("Already posted: " + title)
				continue
			}
			if IsAmount(amount) {
				Logger("Posted already enough today.")
				SetField("LastMod", today)
				return
			}
			SendEmbed(post, discordWebhook)
			SetField("PostedURLs", url)
		}
	}
}
