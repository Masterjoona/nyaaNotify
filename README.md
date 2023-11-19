# nyaaNotify

A simple Go program that checks for new torrents on [nyaa](https://nyaa.si) using cron and sends a discord webhook notification.

## Usage
1. Clone the repository and build the binary

`git clone https://github.com/Masterjoona/nyaaNotify`
`cd nyaaNotify && go build`
2. `./nyaaNotify -generate` and read the comments.
3. Copy the printed line to your crontab.

## Flags
### First time flags
- `-generate` opens your `$EDITOR` and lets you set parameters
### Required flags
- `-webhook` webhook url
- `-name` Name for logging and posted.txt
- `-include` words that the title must include. `,` separate with commas and words starting with `;` are excluded (Intuitive, right?)
- `-regex` lets you set a regex that the title must match. Golang flavor.
- `-amount` how many torrents will be sent to discord.
> [!NOTE]  
> Either `-include` or `-regex` must be set. If both are set, regex will be checked first.

> [!TIP]
> `-include` is often sufficient without the need for regex. For instance: `-include="jujutsu,kaisen,1080,sub`. However, keep in mind that this might match multiple torrents on release day. To control the quantity, you can use `-amount.`"
### Optional flags
- `-shortener` input a url shortener. Because `[Magnet](magnet:?xt=urn:btih:...)` won't get markdowned in discord.
- `-token` the token for the url shortener

## Modules
- [go-editor-capture-input](https://github.com/halkyon/go-editor-capture-input)
- [discordwebhook](https://github.com/gtuk/discordwebhook)

## Screenshots
![screenshot](https://bin.masterjoona.dev/u/DaNTbR.png)
![screenshot](https://bin.masterjoona.dev/u/L7Zw6K.png)

*What `./nyaaNotify -generate` looks like*

## Notes
- I can't guarantee it always works as intended but I tried to make it good enough.
- Beginner project, don't expect too much.
- I'm not responsible for anything you do with this program.