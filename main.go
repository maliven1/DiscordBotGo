package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!GoPlay"

type answers struct {
	OriginChannelId string
	MusicURL        string
}

var responses map[int]answers = map[int]answers{}

func main() {
	//Cashe := cache.New()

	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}
		SplitUlrMusic := strings.Split(args[1], "")
		CheckUrlMusic := strings.Join(SplitUlrMusic[:29], "")
		if CheckUrlMusic == "https://www.youtube.com/watch" {
			author := discordgo.MessageEmbedAuthor{
				Name: m.Author.Username,
			}
			embed := discordgo.MessageEmbed{
				Title:  "YouTube",
				URL:    args[1],
				Author: &author,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Start Bot")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, os.Interrupt)
	<-sc
	fmt.Println("Stop Bot")
}
