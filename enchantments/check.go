package enchantments

import (
	"github.com/df-mc/dragonfly/server/player"
)

func RunEnchantmentCheckOnHit(enchantment Enchantment, player *player.Player, hit *player.Player) {
	if enchant, ok := enchantment.(SwordEnchantment); ok {
		enchant.Trigger(player, hit)
	}
}
