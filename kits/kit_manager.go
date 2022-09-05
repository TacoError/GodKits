package kits

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

type Kit struct {
	Name  string
	Armor []item.Stack
	Sword item.Stack
}

func (k Kit) GiveKit(player *player.Player) {
	armor := player.Armour()
	armor.Clear()
	inventory := player.Inventory()
	inventory.Clear()
	armor.SetHelmet(k.Armor[0])
	armor.SetChestplate(k.Armor[1])
	armor.SetLeggings(k.Armor[2])
	armor.SetBoots(k.Armor[3])

	_ = inventory.SetItem(1, item.NewStack(item.GoldenApple{}, 32))
	_ = inventory.SetItem(0, k.Sword)

	player.Messagef("§fYou have equipped the kit§7: §e%v", k.Name)
}

var kits = make(map[string]Kit)

func Load() {
	kits["Regular"] = regular
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
