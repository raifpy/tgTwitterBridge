package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/raifpy/Go/errHandler"
	"gopkg.in/tucnak/telebot.v2"
)

var tgBot *telebot.Bot = nil
var logID int

func getTelegramJSONValues() telegramConfigStruct {
	data, err := ioutil.ReadFile(telegramJSONPATH)
	if err != nil {
		log.Fatalln(err)
	}

	var tgStruct telegramConfigStruct
	err = json.Unmarshal(data, &tgStruct)
	errHandler.HandlerExit(err)

	logID = int(tgStruct.LogID)

	return tgStruct

}

func initTelegram(twitter *twitterClient) {
	log.Println("Starting Telegram Bot")
	values := getTelegramJSONValues()
	bot, err := telebot.NewBot(telebot.Settings{
		//Token:  string(tgTokenData),
		Token:  values.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	errHandler.HandlerExit(err)
	tgBot = bot
	log.Println("Started Telegram Bot")
	bot.Handle(telebot.OnChannelPost, func(m *telebot.Message) {
		if m.Chat.ID != values.ChannelID {
			log.Println(m.Chat.ID, " not accepted")
			return
		}

		//fmt.Println(m)

		if m.Animation != nil {
			text := m.Caption
			log.Println("Animation", text)
			readClose, err := bot.GetFile(&m.Animation.File)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			data, err := ioutil.ReadAll(readClose)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			readClose.Close()
			media, err := twitter.uploadMedia("mp4", data)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}

			twitter.sendMedia(text, []string{media})

		} else if m.Photo != nil {
			text := m.Caption
			fmt.Println("Photo", text)
			fmt.Println("¿?*")

			readClose, err := bot.GetFile(&m.Photo.File)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			data, err := ioutil.ReadAll(readClose)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			readClose.Close()
			media, err := twitter.uploadMedia("png", data)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			twitter.sendMedia(text, []string{media})

		} else if m.Video != nil {
			text := m.Caption
			fmt.Println("Video", text)
			readClose, err := bot.GetFile(&m.Video.File)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			data, err := ioutil.ReadAll(readClose)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			readClose.Close()
			media, err := twitter.uploadMedia("mp4", data)
			if err != nil {
				tgLog(err.Error(), int(values.LogID))
				return
			}
			twitter.sendMedia(text, []string{media})
		} else {
			fmt.Println(m.Text)
			twitter.sendTweet(m.Text)
		}

	})
	fmt.Println(channelInfo(values.ChannelID, bot))
	log.Println("Waiting any channel post")
	bot.Start()
}

func channelInfo(channelID int64, bot *telebot.Bot) string {
	chat, err := bot.ChatByID(fmt.Sprint(channelID))
	if err != nil {
		tgLog(err.Error(), logID)
		return ""
	}
	return fmt.Sprintf("Channel: %s\nUserName: %s\nDescription: %s",
		chat.Title, chat.Username, chat.Description,
	)

}
