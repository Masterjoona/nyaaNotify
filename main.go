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

	if TestMatchTitle != "" && (includeString != "" || regexString != "") {
		TestMatches(includeString, regexString)
		return
	}

	if discordWebhook == "" {
		Logger("Running in no webhook mode. Only printing matches.")
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
				return
			}

			SetField("PostedURLs", postURL)
			SetField("LastMod", today)

			if discordWebhook == "" {
				continue
			}

			SendEmbed(nyaaPosts[i], discordWebhook)
		}
	}
}
