package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    token := os.Getenv("TOKEN")

    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil { // If we got a message
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            switch update.Message.Command() {
            case "help":
                commandHelp(bot, update.Message)
            case "list":
                commandList(bot, update.Message)
            default:
                defaultBehavior(bot, update.Message)
            }
        }
    }
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
    msg.ReplyToMessageID = inputMessage.MessageID
    bot.Send(msg)
}

func commandHelp(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
        "/help — help"
        + "\n/list — list of products"
    )
    bot.Send(msg)
}

func commandList(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "List")
    bot.Send(msg)
}
