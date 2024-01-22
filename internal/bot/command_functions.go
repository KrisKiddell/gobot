package bot

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func CreateCommandFuncs() map[string]func() *discordgo.InteractionResponseData {
	commandFuncs := make(map[string]func() *discordgo.InteractionResponseData)

	commandFuncs["flip-a-coin"] = flipACoin
	commandFuncs["be-the-very-best"] = beTheVeryBest

	return commandFuncs
}

func flipACoin() *discordgo.InteractionResponseData {

	res := func() string {
		if rand.Intn(2) == 0 {
			return "Heads"
		} else {
			return "Tails"
		}
	}()

	return &discordgo.InteractionResponseData{
		Content: res,
	}
}

func beTheVeryBest() *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Content: "https://www.google.com",
	}
}
