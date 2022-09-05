package main

import (
	"GodKits/commands"
	"GodKits/enchantments"
	"GodKits/handler"
	"GodKits/kits"
	"GodKits/utils"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config := server.DefaultConfig()
	srv := server.New(&config, log)

	srv.CloseOnProgramEnd()
	srv.JoinMessage("§eWelcome, §f%v")
	srv.QuitMessage("§eGoodbye, §f%v")
	srv.SetName("GodKits")

	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}

	kits.Load()
	enchantments.Register()

	cmd.Register(cmd.New("gkit", "Get a kit (/gkit type)", []string{"gkits"}, commands.GKit{}))

	for srv.Accept(func(player *player.Player) {
		utils.CanPvP[player.Name()] = false
		player.Handle(&handler.Handler{
			Player: player,
		})
		player.Inventory().Clear()
		player.Armour().Clear()

		player.Teleport(srv.World().Spawn().Vec3())

		player.Message("§aYou have spawned, use the command §e/gkit §ato equip a GKit!")
	}) {
	}
}
