package enchantments

import (
	"GodKits/utils"
	"fmt"
	"github.com/df-mc/dragonfly/server/item"
)

func LoreItem(stack *item.Stack) {
	var enchants []string
	for _, enchantment := range stack.Enchantments() {
		id, _ := item.EnchantmentID(enchantment.Type())

		// Custom enchants start at id 50
		if id > 49 {
			enchants = append(enchants, fmt.Sprintf("§r§f%v §e%v", enchantment.Type().Name(), utils.IntegerToRoman(enchantment.Level())))
		}
	}
	*stack = stack.WithLore(enchants...)
}
