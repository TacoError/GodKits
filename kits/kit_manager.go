package kits

import "github.com/df-mc/dragonfly/server/player"

type Kit struct {
	GiveKit func(player *player.Player)
}

var kits = make(map[string]Kit)

func Load() {
	kits["Regular"] = Kit{GiveKit: GiveRegular}
	kits["SpeedDemon"] = Kit{GiveKit: GiveSpeedDemon}
}

func KitExists(name string) bool {
	if _, ok := kits[name]; ok {
		return true
	}
	return false
}

func GetKitNameList() []string {
	names := make([]string, 0, len(kits))
	for name := range kits {
		names = append(names, name)
	}
	return names
}

func GiveKit(player *player.Player, kit string) {
	kits[kit].GiveKit(player)
	player.Messagef("§eYou have received the kit§7: §a%v§e!", kit)
}
