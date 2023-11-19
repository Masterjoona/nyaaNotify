package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/halkyon/go-editor-capture-input/pkg/editor"
)

var name string

func ParseCommandParameters() (bool, string, string, string, string, string, string) {
	var generateCron bool
	var discordWebhook string
	var includeString string
	var regexString string
	var amount string

	var shortenerToken string
	var shortenerURL string

	flag.BoolVar(&generateCron, "generate", false, "generate a cron expression")
	flag.StringVar(&discordWebhook, "webhook", "", "discord webhook url")
	flag.StringVar(&shortenerToken, "token", "", "see readme")
	flag.StringVar(&shortenerURL, "shortener", "", "see readme")
	flag.StringVar(&includeString, "include", "", "strings to look for")
	flag.StringVar(&regexString, "regex", "", "regex to match.")
	flag.StringVar(&amount, "amount", "", "stop posting after x links have been posted.")
	flag.StringVar(&name, "name", "", "name for logging and posted.txt")
	flag.Parse()
	return generateCron, discordWebhook, shortenerToken, includeString, regexString, shortenerURL, amount
}

func parseParameters(parameters string) (cron, include, regex, webhook, nameName, apiUrl, token, amount string) {
	lines := strings.Split(parameters, "\n")

	prefixes := map[string]*string{
		"cron=":      &cron,
		"include=":   &include,
		"regex=":     &regex,
		"webhook=":   &webhook,
		"name=":      &nameName,
		"shortener=": &apiUrl,
		"token=":     &token,
		"amount=":    &amount,
	}

	for _, line := range lines {
		for prefix, target := range prefixes {
			if strings.HasPrefix(line, prefix) {
				*target = strings.TrimSpace(strings.Split(line[len(prefix):], "#")[0])
			}
		}
	}
	name = nameName
	return
}

func MakeParameters() {
	weekday := time.Now().Weekday()
	hour := time.Now().Hour()
	defaultCron := "# Run every 10 minutes between " + fmt.Sprintf("%d:00", hour) + " and " + fmt.Sprintf("%d:00", hour+1) + " on " + weekday.String() + ".\n" + fmt.Sprintf("cron=*/10 %d-%d * * %d", hour, hour+1, weekday)
	defaultAmount := "# Stop posting after x links have been posted.\namount=1"
	defaultInclude := "\n# Look for banana, apple while ignoring orange. Case insensitive.\ninclude=banana,apple,;orange"
	defaultRegex := "# Look for case insenstive banana, apple or orange. Golang regex flavor.\n#regex=(?i)banana|apple|orange\n"
	defaultWebhook := "webhook=https://discord.com/api/webhooks/123/abcdef\n"
	defaultName := "# Name for logging and posted.txt\nname=" + weekday.String() + "_" + fmt.Sprintf("%d", hour)
	defaultURLShortenAPI := "\n# Api url for shortening urls.\n#shortener="
	defajltURLShortenToken := "# Token for shortening urls.\n#token="
	defaultText := defaultCron + "\n" + defaultAmount + "\n" + defaultInclude + "\n" + defaultRegex + "\n" + defaultWebhook + "\n" + defaultName + "\n" + defaultURLShortenAPI + "\n" + defajltURLShortenToken
	editor := editor.New([]byte(defaultText), "parameters.sh")
	output, err := editor.Run()
	if err != nil {
		panic(err)
	}

	cron, include, regex, webhook, name, shortenerURL, shortenerToken, amount := parseParameters(string(output))
	if cron == "" || webhook == "" || (include == "" && regex == "") || name == "" || ((shortenerURL == "" || shortenerToken == "") && (shortenerURL != "" || shortenerToken != "")) {
		Logger("cron=" + cron)
		Logger("include=" + include)
		Logger("regex=" + regex)
		Logger("webhook=" + webhook)
		Logger("name=" + name)
		Logger("amount=" + amount)
		if (shortenerURL == "" && shortenerToken != "") || (shortenerURL != "" && shortenerToken == "") {
			Logger("shortener=" + shortenerURL)
			Logger("token=" + shortenerToken)
		}
		panic("Invalid parameters")
	}

	maybeRegex := OptionalParam("-regex", regex)
	maybeInclude := OptionalParam("-include", include)
	maybeShorten := OptionalParam("-shortener", shortenerURL) + OptionalParam("-token", shortenerToken)
	fullCommand := cron + " " + executablePath + "/nyaaNotify" + " -name='" + name + "' -webhook='" + webhook + "' -amount='" + amount + "'" + maybeInclude + maybeRegex + maybeShorten
	println(fullCommand)
}
