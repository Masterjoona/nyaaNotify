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

var Url = "https://nyaa.si"

func main() {
	generateCron, discordWebhook, includeString, regexString, amount := ParseCommandParameters()

	if generateCron {
		MakeParameters()
		return
	}
	today := GetDate()
	if GetField("LastMod") != today && GetField("LastMod") != "" {
		Logger("Seems like a week has passed, cleaning JSON and log file.")
		CleanFiles()
	}

	if IsAmount(amount) {
		Logger("Posted already enough today.")
		SetField("LastMod", today)
		return
	}

	allMatches, nyaaPosts := MatchPosts()
	for i, matches := range allMatches {
		if matches[5] == "" {
			matches[5] = "0" // comments
		}

		nyaaPosts[i] = CreateNyaaPost(matches)
		title := nyaaPosts[i].Title

		if MatchTitle(title, includeString, regexString) {
			Logger("Found match: " + title)
			postURL := nyaaPosts[i].URL

			if AlreadyPosted(postURL) {
				Logger("Already posted: " + title)
				continue
			}

			if IsAmount(amount) {
				Logger("Posted already enough today.")
				SetField("LastMod", today)
				return
			}

			SendEmbed(nyaaPosts[i], discordWebhook)
			SetField("PostedURLs", postURL)
		}
	}
}
