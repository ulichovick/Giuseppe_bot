package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Secret struct {
	Secret string
}

func main() {
	secret, errs := ioutil.ReadFile("./secrets.json")
	if errs != nil {
		log.Fatal("Error leyendo el archivo del secret!: ", errs)
	}
	var data Secret
	errs = json.Unmarshal(secret, &data)
	if errs != nil {
		log.Fatal("Error extrayendo el secret: ", errs)
	}

	bot, err := tgbotapi.NewBotAPI(data.Secret)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Cuenta autorizada  %s \n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)

	updates := bot.GetUpdatesChan(u)

	go getUpdates(updates, bot)

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

func getUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {

	for update := range updates {
		if update.Message != nil {
			log.Printf("%s escribio %s", update.Message.From.FirstName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Give me a second please, im generating ypur answer")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			/* msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text) */
			cmd := exec.Command("python3", "llama.py", "-w", string(update.Message.Text))
			out, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
			text := string(out)
			log.Println(out)
			msg_ans := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg_ans.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg_ans)
		}
	}
}
