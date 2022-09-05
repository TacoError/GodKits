package enchantments

import (
	"github.com/df-mc/dragonfly/server/item"
)

var CustomEnchants = make(map[int]Enchantment)

func Register() {
	CustomEnchants[50] = Freeze{}

	for id, enchantment := range CustomEnchants {
		item.RegisterEnchantment(id, enchantment)
	}
}
