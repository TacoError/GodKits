package commands

import (
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
	if kits.KitExists(g.Type) {
		kits.GiveKit(sender, g.Type)
		return
	}
	sender.Messagef("§cThere is no kit with that name! §7(§a%v§7)", strings.Join(kits.GetKitNameList(), ", "))
}
