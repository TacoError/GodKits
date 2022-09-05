package enchantments

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

type Enchantment interface {
	item.EnchantmentType
	GetChance() (int, int)
}

type SwordEnchantment interface {
	Enchantment
	Trigger(player *player.Player, hit *player.Player)
}
