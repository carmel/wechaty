package main

import (
	"fmt"
	"go-wechaty/wechaty"
	wp "go-wechaty/wechaty-puppet"
	"go-wechaty/wechaty-puppet/schemas"
	"go-wechaty/wechaty/user"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	var bot = wechaty.NewWechaty(wechaty.WithPuppetOption(wp.Option{
		Token: "",
	}))

	bot.OnScan(func(qrCode string, status schemas.ScanStatus, data string) {
		fmt.Printf("Scan QR Code to login: %v\nhttps://wechaty.github.io/qrcode/%s\n", status, qrCode)
	}).OnLogin(func(user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnMessage(onMessage).OnLogout(func(user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	var err = bot.Start()
	if err != nil {
		panic(err)
	}

	var quitSig = make(chan os.Signal)
	signal.Notify(quitSig, os.Interrupt, os.Kill)

	select {
	case <-quitSig:
		log.Fatal("exit.by.signal")
	}
}

func onMessage(message *user.Message) {
	log.Println(message)

	if message.Self() {
		log.Println("Message discarded because its outgoing")
	}

	if message.Age() > 2*60*time.Second {
		log.Println("Message discarded because its TOO OLD(than 2 minutes)")
	}

	if message.Type() != schemas.MessageTypeText || message.Text() != "#ding" {
		log.Println("Message discarded because it does not match #ding")
		return
	}

	// 1. reply 'dong'
	_, err := message.Say("dong")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("REPLY: dong")

	//	// 2. reply image(qrcode image)
	//	fileBox, _ := file_box.NewFileBoxFromUrl("https://wechaty.github.io/wechaty/images/bot-qr-code.png", "", nil)
	//	_, err = message.SayFile(fileBox)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Printf("REPLY: %s\n", fileBox)
}
