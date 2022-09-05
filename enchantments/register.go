package enchantments

import (
	"github.com/df-mc/dragonfly/server/item"
)

var customEnchants = make(map[int]Enchantment)

func Register() {
	customEnchants[50] = Freeze{}

	for id, enchantment := range customEnchants {
		item.RegisterEnchantment(id, enchantment)
	}
}
