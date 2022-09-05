package kits

import (
	"GodKits/enchantments"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
)

var regular = Kit{
	Name: "Regular",
	Armor: []item.Stack{
		item.NewStack(item.Helmet{
			Tier: item.ArmourTierDiamond{},
		}, 1).WithEnchantments(
			item.NewEnchantment(enchantment.Protection{}, 4),
			item.NewEnchantment(enchantment.Unbreaking{}, 3),
		),
		item.NewStack(item.Chestplate{
			Tier: item.ArmourTierDiamond{},
		}, 1).WithEnchantments(
			item.NewEnchantment(enchantment.Protection{}, 4),
			item.NewEnchantment(enchantment.Unbreaking{}, 3),
		),
		item.NewStack(item.Leggings{
			Tier: item.ArmourTierDiamond{},
		}, 1).WithEnchantments(
			item.NewEnchantment(enchantment.Protection{}, 4),
			item.NewEnchantment(enchantment.Unbreaking{}, 3),
		),
		item.NewStack(item.Boots{
			Tier: item.ArmourTierDiamond{},
		}, 1).WithEnchantments(
			item.NewEnchantment(enchantment.Protection{}, 4),
			item.NewEnchantment(enchantment.Unbreaking{}, 3),
		),
	},
	Sword: item.NewStack(item.Sword{
		Tier: item.ToolTierDiamond,
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Sharpness{}, 3),
		item.NewEnchantment(enchantment.Unbreaking{}, 2),
		item.NewEnchantment(enchantments.Freeze{}, 4),
	),
}
