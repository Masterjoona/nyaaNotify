package main

import (
	"fmt"

	"github.com/gtuk/discordwebhook"
)

func makeDescription(post NyaaPost, shortenerToken string, shortenerURL string) string {
	comments := ""
	if post.Comments != "0" {
		comments = post.Comments + " comments"
	}
	if shortenerToken == "" || shortenerURL == "" {
		torrent := "[Torrent](" + post.Torrent + ")"
		return fmt.Sprintf("Size: %s | %s | <t:%s:R> | %s", post.Size, torrent, post.Date, comments)
	}
	shortURL, err := ShortenURL(post.Magnet, shortenerToken, shortenerURL)
	if err != nil {
		Logger("Error shortening url: " + err.Error())
		return "[Torrent](" + post.Torrent + ") | Error shortening url."
	}
	magnet := "[Magnet](" + shortURL + ")"

	return fmt.Sprintf("Size: %s | %s | <t:%s:R> | %s", post.Size, magnet, post.Date, comments)

}

func SendEmbed(post NyaaPost, discordWebhook string) {
	seeds := "Seeders"
	leechs := "Leechers"
	completed := "Completed"

	category := "Category"
	trueBool := true

	fields := []discordwebhook.Field{
		{
			Name:   &category,
			Value:  &post.Category,
			Inline: &trueBool,
		},
		{
			Name:   &seeds,
			Value:  &post.Seed,
			Inline: &trueBool,
		},
		{
			Name:   &leechs,
			Value:  &post.Leech,
			Inline: &trueBool,
		},
		{
			Name:   &completed,
			Value:  &post.Completed,
			Inline: &trueBool,
		},
	}
	thumb := discordwebhook.Thumbnail{
		Url: &post.CategoryImg,
	}
	description := makeDescription(post, ShortenerToken, ShortenerURL)
	embed := discordwebhook.Embed{
		Title:       &post.Title,
		Url:         &post.URL,
		Description: &description,
		Fields:      &fields,
		Thumbnail:   &thumb,
	}

	username := "Nyaa.si Notification"
	embeds := []discordwebhook.Embed{embed}

	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &embeds,
	}

	err := discordwebhook.SendMessage(discordWebhook, message)
	if err != nil {
		Logger("Error sending message: " + err.Error())
		panic(err)
	}
}
