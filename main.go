package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/wizardfrag/gitlab-discord/gitlabdiscord"
	"log"
	"github.com/wizardfrag/gitlab-discord/discord"
)

var configFile string

func main() {
	var config gitlabdiscord.Config
	cFile := flag.String("config", "./gitlab-discord.toml", "Specify the config to use for gitlab-discord")
	flag.Parse()

	configFile = *cFile

	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatalln("error loading config file,", err)
	}

	log.Println("starting gitlab-discord")

	_, err = discord.Run(config)
	if err != nil {
		log.Fatal("error starting discord. Stopping!")
	}

	<- make(chan bool)
	return
}
