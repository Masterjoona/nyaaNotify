package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/halkyon/go-editor-capture-input/pkg/editor"
)

var Name string
var TestMatchTitle string

func ParseCommandParameters() (generateCron bool, discordWebhook, includeString, regexString, amount, category string) {

	flag.BoolVar(&generateCron, "generate", false, "generate a cron expression")
	flag.BoolVar(&generateCron, "g", false, "generate a cron expression")
	flag.StringVar(&discordWebhook, "webhook", "", "discord webhook url")
	flag.StringVar(&includeString, "include", "", "strings to look for")
	flag.StringVar(&regexString, "regex", "", "regex to match.")
	flag.StringVar(&amount, "amount", "", "stop posting after x links have been posted.")
	flag.StringVar(&Name, "name", "", "name for logging and posted.txt")
	flag.StringVar(&category, "category", "", "category to look for.")
	flag.StringVar(&TestMatchTitle, "testTitle", "", "title to test the matching on. for multiple titles, separate with ;")
	flag.Parse()

	var executablePathLocal = getExecutablePath()
	PostedFile = executablePathLocal + "/" + Name + "_posted.json"
	LogFile = executablePathLocal + "/" + Name + "_log.txt"
	ExecutablePath = executablePathLocal

	return generateCron, discordWebhook, includeString, regexString, amount, category
}

func parseParameters(parameters string) (cron, include, regex, webhook, nameName, amount, category, path string) {
	lines := strings.Split(parameters, "\n")

	prefixes := map[string]*string{
		"cron=":     &cron,
		"include=":  &include,
		"regex=":    &regex,
		"webhook=":  &webhook,
		"name=":     &nameName,
		"amount=":   &amount,
		"category=": &category,
		"path=":     &path,
	}

	for _, line := range lines {
		for prefix, target := range prefixes {
			if strings.HasPrefix(line, prefix) {
				*target = strings.TrimSpace(strings.Split(line[len(prefix):], "#")[0])
			}
		}
	}
	Name = nameName
	return
}

func MakeParameters() {
	weekday := time.Now().Weekday()
	hour := time.Now().Hour()
	defaultCron := "# Run every 10 minutes between " + fmt.Sprintf("%d:00", hour) + " and " + fmt.Sprintf("%d:00", hour+1) + " on " + weekday.String() + ".\n" + fmt.Sprintf("cron=*/10 %d-%d * * %d", hour, hour+1, weekday)
	defaultAmount := "# Stop posting after x links have been posted.\namount=1"
	defaultInclude := "\n# Look for banana, apple or blueberry and while ignoring orange. Case insensitive.\n#include=banana,(apple|blueberry),;orange\n"
	defaultRegex := "# Look for case insenstive banana, apple or orange. Golang regex flavor.\n#regex=(?i)banana|apple|orange\n"
	defaultWebhook := "webhook=discord\n"
	defaultName := "# Name for logging and posted.txt\nname=" + weekday.String() + "_" + fmt.Sprintf("%d", hour)
	defaultCategory := "# Category to look for. Can either be number_number or the string equivalent.\n#category=1_2\n#category=Anime - English-translated"
	defaultPath := "# Path to nyaaNotify executable. Mostly for having a variable in crontab.\n#path=./nyaaNotify\n"
	defaultText := defaultCron + "\n" + defaultAmount + "\n" + defaultInclude + "\n" + defaultRegex + "\n" + defaultWebhook + "\n" + defaultName + "\n" + defaultCategory + "\n" + defaultPath
	editor := editor.New([]byte(defaultText), "parameters.sh")
	output, err := editor.Run()
	if err != nil {
		Logger("Error running editor: " + err.Error())
		panic(err)
	}

	cron, include, regex, webhook, name, amount, category, path := parseParameters(string(output))
	if cron == "" || webhook == "" || (include == "" && regex == "") || name == "" {
		Logger("cron=" + cron)
		Logger("include=" + include)
		Logger("regex=" + regex)
		Logger("webhook=" + webhook)
		Logger("name=" + name)
		Logger("amount=" + amount)
		Logger("category=" + category)
		panic("Invalid parameters!")
	}

	matching := ParamCreation("regex", regex) + ParamCreation("include", include)
	base := cron + " " + ExecutablePath + "/nyaaNotify"
	if path != "" {
		base = cron + " " + path
	}
	fullCommand := base + ParamCreation("name", name) + ParamCreation("webhook", webhook) + ParamCreation("amount", amount) + matching + ParamCreation("category", category)
	println(fullCommand)
}
