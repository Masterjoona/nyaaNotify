package main

import "github.com/gtuk/discordwebhook"

func SendEmbed(post NyaaPost, description string, discordWebhook string) {
	seeds := "Seeders"
	leechs := "Leechers"
	cate := "Category"
	trueBool := true
	fields := &[]discordwebhook.Field{
		{
			Name:   &cate,
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
	}
	thumb := discordwebhook.Thumbnail{
		Url: &post.CategoryImg,
	}
	embed := discordwebhook.Embed{
		Title:       &post.Title,
		Url:         &post.URL,
		Description: &description,
		Fields:      fields,
		Thumbnail:   &thumb,
	}
	username := "Nyaa.si Notification"
	embeds := &[]discordwebhook.Embed{embed}
	message := discordwebhook.Message{
		Username: &username,
		Embeds:   embeds,
	}
	err := discordwebhook.SendMessage(discordWebhook, message)
	if err != nil {
		Logger("Error sending message: " + err.Error())
		panic(err)
	}
}
