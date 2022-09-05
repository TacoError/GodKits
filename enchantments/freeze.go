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
	return 5, 50
}

func (Freeze) MaxLevel() int {
	return 5
}

func (Freeze) Cost(int) (int, int) {
	return 1, 1
}

func (Freeze) Trigger(_ *player.Player, hit *player.Player) {
	hit.SetImmobile()
	for i := 0; i <= 5; i++ {
		hit.Hurt(3, damage.SourceLightning{})
		time.Sleep(250 * time.Millisecond)
	}
	hit.SetMobile()
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
