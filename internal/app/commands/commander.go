package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/meesooqa/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	c.bot.Send(msg)
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help — help"+
			"\n/list — list of products",
	)
	c.bot.Send(msg)
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
