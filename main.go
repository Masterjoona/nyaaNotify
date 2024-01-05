package main

var Url = "https://nyaa.si"

func main() {
	generateCron, discordWebhook, includeString, regexString, amount, category := ParseCommandParameters()

	if generateCron {
		MakeParameters()
		return
	}

	if TestMatchTitle != "" && (includeString != "" || regexString != "") {
		TestMatches(includeString, regexString, category)
		return
	}

	if discordWebhook == "" {
		Logger("Running in no webhook mode. Only printing matches.")
	}

	today := GetDate()

	if GetField("LastMod") != today && GetField("LastMod") != "" {
		Logger("Seems like some time has passed, cleaning JSON and log file.")
		CleanFiles()
	}

	if IsAmount(amount) {
		Logger("Posted already enough today.")
		SetField("LastMod", today)
		return
	}

	nyaaPosts := GetNyaaPosts()
	for _, nyaaPost := range nyaaPosts {
		title := nyaaPost.Title
		postCategories := []string{nyaaPost.Category, nyaaPost.CategoryId}
		if MatchPost(title, includeString, regexString, category, postCategories) {
			Logger("Found match: " + title)
			println(nyaaPost.Link)
			postURL := nyaaPost.URL

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

			SendEmbed(nyaaPost, discordWebhook)
		}
	}
}
