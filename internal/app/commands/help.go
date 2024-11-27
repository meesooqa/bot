package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func init() {
	registeredCommands["help"] = (*Commander).Help
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help — help"+
			"\n/list — list of products",
	)
	c.bot.Send(msg)
}
