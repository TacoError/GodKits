package enchantments

import (
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type Freeze struct {
	Enchantment
}

func (Freeze) Name() string {
	return "Freeze"
}

func (Freeze) GetChance() (int, int) {
	return 5, 250
}

func (Freeze) MaxLevel() int {
	return 5
}

func (Freeze) Cost(int) (int, int) {
	return 1, 1
}

func (Freeze) Trigger(_ *player.Player, hit *player.Player) {
	go func() {
		hit.SetImmobile()
		times := 0
		ticker := time.NewTicker(time.Millisecond * 250)
		defer func() {
			ticker.Stop()
			hit.SetMobile()
		}()
		for range ticker.C {
			if times > 5 {
				return
			}
			hit.Hurt(3, damage.SourceLightning{})
			times++
		}
	}()
}

func (Freeze) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

func (Freeze) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

func (Freeze) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && (t.ToolType() == item.TypeSword || t.ToolType() == item.TypeAxe)
}
