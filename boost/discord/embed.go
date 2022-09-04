package discord

import dg "github.com/bwmarrin/discordgo"

const (
	embedColor = 0x1F7EAD // color being used in python script so when changing this, change it in python too
	footerText = "https://discord.gg/ZR2EspdHFQ ● Icons by: icons8.com"
)

// returns new embed
func NewEmbed() *Embed {
	return &Embed{&dg.MessageEmbed{}}
}

type Embed struct {
	*dg.MessageEmbed
}

// sets up the embed to have the color and footer good to go
func (e *Embed) SetupEmbed() *Embed {
	e.Color = embedColor
	e.SetFooter("attachment://boost.png", footerText)

	return e
}

func (e *Embed) SetFooter(iconUrl, text string) *Embed {
	e.Footer = &dg.MessageEmbedFooter{
		IconURL: iconUrl,
		Text:    text,
	}

	return e
}

func (e *Embed) SetAuthor(iconUrl, name, url string) *Embed {
	e.Author = &dg.MessageEmbedAuthor{
		IconURL: iconUrl,
		Name:    name,
		URL:     url,
	}

	return e
}

func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) AddField(name, value string, inline bool) *Embed {
	e.Fields = append(e.Fields, &dg.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})

	return e
}
