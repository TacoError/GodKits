package commands

import (
	"GodKits/handler"
	"GodKits/kits"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
)

type GKit struct {
	Type string
}

func (g GKit) Run(src cmd.Source, _ *cmd.Output) {
	sender := src.(*player.Player)
	if h, ok := sender.Handler().(*handler.Handler); ok {
		if h.CanPvP.Load() {
			sender.Message("§cYou cannot do this whilst already in a kit")
			return
		}
	}

	if kits.KitExists(g.Type) {
		session := sender.Handler().(*handler.Handler)
		session.CanPvP.Store(true)
		kits.GiveKit(sender, g.Type)
		return
	}
	sender.Messagef("§cThere is no kit with that name! §7(§a%v§7)", strings.Join(kits.GetKitNameList(), ", "))
}
