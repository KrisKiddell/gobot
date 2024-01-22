package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kriskiddell/plog"
)

var cmd = []*discordgo.ApplicationCommand{
	{
		Name:        "flip-a-coin",
		Description: "Heads for Tails? Who knows!",
	},
	{
		Name:        "be-the-very-best",
		Description: "Do you want to be the very best? Like no one ever was?",
	},
}

func (b *Bot) RegisterSlashCommands() {
	plog.Info.Println("Registering slash commands...")

	for _, v := range cmd {
		plog.Info.Println("Registering slash command:", v.Name)
		cmd, err := b.CreateSlashCommand(v.Name, v.Description)

		if err != nil {
			plog.Error.Println("Unable to load command", v.Name, err)
		}

		plog.Success.Println(v.Name, "Registered!")
		b.Commands = append(b.Commands, cmd)
	}

}

func (b *Bot) CreateSlashCommand(name string, description string) (*discordgo.ApplicationCommand, error) {
	return b.Session.ApplicationCommandCreate(b.AppId, "", &discordgo.ApplicationCommand{
		Name:        name,
		Description: description,
	})
}

func (b *Bot) RemoveSlashCommands() {
	for _, cmd := range b.Commands {
		err := b.Session.ApplicationCommandDelete(b.AppId, "", cmd.ID)
		if err != nil {
			plog.Error.Println("Unable to remove command:", cmd.Name)
			continue
		}
		plog.Success.Println("Removed command", cmd.Name)
	}
}
