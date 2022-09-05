package enchantments

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"math/rand"
)

func RunEnchantmentCheckOnHit(enchantment Enchantment, level int, player *player.Player, hit *player.Player) {
	multiplier, max := enchantment.GetChance()
	if (level * multiplier) < rand.Intn(max) {
		player.SendJukeboxPopup(fmt.Sprintf("§e%v §fhas triggered!", enchantment.Name()))
		go enchantment.Trigger(player, hit)
	}
}
