package enchantments

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Freeze struct {
	SwordEnchantment
}

func (Freeze) Name() string {
	return "Freeze"
}

func (Freeze) GetChance() (int, int) {
	return 5, 100
}

func (Freeze) MaxLevel() int {
	return 5
}

func (Freeze) Cost(level int) (int, int) {
	min := 1 + (level-1)*11
	return min, min + 20
}

func (Freeze) Trigger(player *player.Player, hit *player.Player) {
	// whatever freeze enchantment does idk
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
