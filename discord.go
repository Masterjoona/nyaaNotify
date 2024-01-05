package main

import (
	"fmt"
	"time"

	"github.com/gtuk/discordwebhook"
)

func makeDescription(post NyaaPost) string {
	comments := ""
	if post.Comments != "0" {
		comments = post.Comments + " comments"
	}
	torrent := "[Torrent](" + post.Link + ")"
	timestamp, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", post.PubDate)
	return fmt.Sprintf("Size: %s | %s | <t:%d:R> | %s", post.Size, torrent, timestamp.Unix(), comments)
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
			Value:  &post.Seeders,
			Inline: &trueBool,
		},
		{
			Name:   &leechs,
			Value:  &post.Leechers,
			Inline: &trueBool,
		},
		{
			Name:   &completed,
			Value:  &post.Downloads,
			Inline: &trueBool,
		},
	}
	categoryImgLink := Url + "/static/img/icons/nyaa/" + post.CategoryId + ".png"
	thumb := discordwebhook.Thumbnail{
		Url: &categoryImgLink,
	}
	description := makeDescription(post)
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
