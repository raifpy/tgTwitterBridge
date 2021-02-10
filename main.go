package main

// Developed by @raifpy

/*
known bugs:
	Telegram Album Media > may lead to spam Twitter
	Twitter authorization may be chance (logout)
*/

import (
	"log"
	"os"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

const defaultJSONPath = "default.json"
const headerJSONPath = "header.json"
const cookiePath = "cookie.txt"
const telegramJSONPATH = "telegram.json"

//var twitter *twitterClient // global twitter

func main() {
	if contains("-background", os.Args) {
		cmd := exec.Command(os.Args[0])

		err := cmd.Start()
		if err == nil {
			os.Exit(0)
		}
		log.Println("Clound't start bridge on background: ", err)

	}
	client, err := newTwitter(defaultJSONPath, headerJSONPath, cookiePath)
	errHandler.HandlerExit(err)

	//twitter = client // global

	initTelegram(client)

}
