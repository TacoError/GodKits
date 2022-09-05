package kits

import (
	"GodKits/enchantments"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/player"
)

func GiveRegular(player *player.Player) {
	helmet := item.NewStack(item.Helmet{
		Tier: item.ArmourTierDiamond{},
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Protection{}, 4),
		item.NewEnchantment(enchantment.Unbreaking{}, 3),
	)

	chestplate := item.NewStack(item.Chestplate{
		Tier: item.ArmourTierDiamond{},
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Protection{}, 4),
		item.NewEnchantment(enchantment.Unbreaking{}, 3),
	)

	leggings := item.NewStack(item.Leggings{
		Tier: item.ArmourTierDiamond{},
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Protection{}, 4),
		item.NewEnchantment(enchantment.Unbreaking{}, 3),
	)

	boots := item.NewStack(item.Boots{
		Tier: item.ArmourTierDiamond{},
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Protection{}, 4),
		item.NewEnchantment(enchantment.Unbreaking{}, 3),
	)

	enchantments.LoreItem(&helmet)
	enchantments.LoreItem(&chestplate)
	enchantments.LoreItem(&leggings)
	enchantments.LoreItem(&boots)

	sword := item.NewStack(item.Sword{
		Tier: item.ToolTierDiamond,
	}, 1).WithEnchantments(
		item.NewEnchantment(enchantment.Sharpness{}, 3),
		item.NewEnchantment(enchantment.Unbreaking{}, 2),
		item.NewEnchantment(enchantments.Freeze{}, 4),
	)
	enchantments.LoreItem(&sword)

	player.Inventory().Clear()
	player.Armour().Clear()
	player.Armour().SetHelmet(helmet)
	player.Armour().SetChestplate(chestplate)
	player.Armour().SetLeggings(leggings)
	player.Armour().SetBoots(boots)
	player.Inventory().SetItem(0, sword)
	player.Inventory().SetItem(1, item.NewStack(item.EnchantedApple{}, 32))
}
