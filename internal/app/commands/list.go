package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func init() {
	registeredCommands["list"] = (*Commander).List
}

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products:"

	products := c.productService.List()
	for _, p := range products {
		outputMsgText += "\n"
		outputMsgText += p.Title
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
