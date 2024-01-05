# nyaaNotify

A simple Go program that checks for new torrents on [nyaa](https://nyaa.si) using cron and sends a discord webhook notification.

## Usage

1. Clone the repository and build the binary

```sh
git clone https://github.com/Masterjoona/nyaaNotify
cd nyaaNotify && go build
```

2. `./nyaaNotify -generate` and read the comments.
3. Copy the printed line to your crontab.

## Flags

### First time flags

-   `-generate` opens your `$EDITOR` and lets you set parameters. `-g` also works.

### Required flags

-   `-webhook` webhook url
-   `-name` Name for logging and posted files
-   `-include` words that the title must include. `,` separate with commas and words starting with `;` are excluded (Intuitive, right?)
-   `-regex` lets you set a regex that the title must match. Golang flavor.
-   `-amount` how many torrents will be sent to discord.
-   `-category` category to search for. For example `1_4` or the string `Anime - Raw`

> [!NOTE]  
> Either `-include` or `-regex` must be set. If both are set, regex will be checked first.

> [!TIP]
> Use `-include` for keyword filtering (e.g., `-include="jujutsu,kaisen,1080,sub"`). Be mindful of multiple matches on an episode day, for that use `-amount` to limit the amount of torrents sent. You can do `(eng|[ani])` for example to match either `eng` or `[ani]` in the title.

> [!TIP]
> You can set variables in crontab `var=value` and reference them by `$var`. Doing this you can avoid long lines with the same webhook url or the filepath!

### Optional flags

-   `-testTitle` lets you test the regex and include flags. `title1;title2`

## Modules

-   [go-editor-capture-input](https://github.com/halkyon/go-editor-capture-input)
-   [discordwebhook](https://github.com/gtuk/discordwebhook)

## Screenshots

![screenshot](https://bin.masterjoona.dev/u/DaNTbR.png)
![screenshot](https://bin.masterjoona.dev/u/L7Zw6K.png)

_What `./nyaaNotify -generate` looks like_

## Notes

-   I can't guarantee it always works as intended but I tried to make it good enough.
-   Beginner project, don't expect too much.
-   I'm not responsible for anything you do with this program.
-   Found [this](https://github.com/Zhousiru/NyaaHub) which is basically this but better.
