package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	token = flag.String("token", os.Getenv("TGAPI_TOKEN"), "Telegram Bot API Token. Defaults to TGAPI_TOKEN env variable.")
	debug = flag.Bool("debug", false, "If set, enables debug logging in telegram bot api.")
)

func main() {
	flag.Parse()

	bot, err := tgapi.NewBotAPI(*token)
	if err != nil {
		log.Fatalf("NewBotAPI(%#v): %v", *token, err)
	}

	bot.Debug = *debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for upd := range updates {
		accept(bot, upd)
	}
}

func accept(bot *tgapi.BotAPI, upd tgapi.Update) {
	m := upd.Message
	if m == nil {
		// We want only message updates.
		return
	}

	var w, h int
	if _, err := fmt.Sscanf(m.Text, "%d %d", &w, &h); err != nil {
		bot.Send(tgapi.NewMessage(m.Chat.ID, err.Error()))
		return
	}

	img, err := samplePng(w, h, func(i, j int) color.RGBA {
		return color.RGBA{
			A: 128,
			R: uint8(i % 100),
			G: uint8(j % 100),
			B: uint8((i + j) % 100),
		}
	})
	if err != nil {
		bot.Send(tgapi.NewMessage(m.Chat.ID, err.Error()))
		return
	}

	// Send the image to the user.
	file := tgapi.FileBytes{Name: "img.png", Bytes: img}
	msg := tgapi.NewPhotoUpload(m.Chat.ID, file)
	bot.Send(msg)
}

func samplePng(w, h int, fn func(i, j int) color.RGBA) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 1; j < h; j++ {
			img.SetRGBA(i, j, fn(i, j))
		}
	}

	// Create a png out of img.
	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
