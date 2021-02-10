package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/raifpy/Go/errHandler"
	"gopkg.in/tucnak/telebot.v2"
)

func req2String(req *http.Request) (int, string) {
	client := http.Client{}
	response, err := client.Do(req)
	data, err := ioutil.ReadAll(response.Body)
	if errHandler.HandlerBool(err) {
		return 0, ""
	}
	return response.StatusCode, string(data)
}

func res2String(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		tgLog(err.Error(), logID)
		return nil
	}
	return body
}

func contains(key string, list []string) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}

func tgLog(text string, id int) {
	log.Println(text)
	if tgBot != nil {
		tgBot.Send(&telebot.User{ID: id}, time.Now().String()+" : "+text)
	}
}
