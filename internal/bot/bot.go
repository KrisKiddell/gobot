package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kriskiddell/plog"
)

type Bot struct {
	Token    string
	AppId    string
	Session  *discordgo.Session
	Commands []*discordgo.ApplicationCommand
}

func NewBot(token string, app_id string) (*Bot, error) {

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, err
	}

	session.AddHandler(messageHandler)
	session.AddHandler(interactionHandler(CreateCommandFuncs()))
	session.Identify.Intents = discordgo.IntentGuildMessages

	bot := &Bot{
		Token:   token,
		AppId:   app_id,
		Session: session,
	}

	return bot, err

}

func messageHandler(s *discordgo.Session, i *discordgo.MessageCreate) {
	plog.Info.Println("A message has been posted")
}

func interactionHandler(cmdfncs map[string]func() *discordgo.InteractionResponseData) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: cmdfncs[i.ApplicationCommandData().Name](),
		})
	}
}
