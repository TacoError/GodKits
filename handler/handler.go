package handler

import (
	"GodKits/enchantments"
	"fmt"
	"github.com/df-mc/atomic"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Handler struct {
	player.NopHandler
	Player *player.Player

	CanPvP    atomic.Bool
	LastHitBy atomic.Value[*player.Player]
}

func (h *Handler) HandleItemDrop(ctx *event.Context, _ *entity.Item) {
	ctx.Cancel()
}

func (h *Handler) HandleFoodLoss(ctx *event.Context, _ int, _ int) {
	ctx.Cancel()
}

func (h *Handler) HandleItemDamage(ctx *event.Context, _ item.Stack, _ int) {
	ctx.Cancel()
}

func (h *Handler) HandleItemPickup(ctx *event.Context, _ item.Stack) {
	ctx.Cancel()
}

func (h *Handler) HandleDeath(damage.Source) {
	lastHitBy := h.LastHitBy.Load()
	if lastHitBy == nil {
		_, _ = chat.Global.WriteString(fmt.Sprintf("§e%v §fhas died.", h.Player.Name()))
		return
	}
	_, _ = chat.Global.WriteString(fmt.Sprintf("§e%v §fhas been killed by §e%v§f!", h.Player.Name(), lastHitBy.Name()))
}

func (h *Handler) HandleAttackEntity(ctx *event.Context, e world.Entity, _, _ *float64, _ *bool) {
	if p, ok := e.(*player.Player); ok {
		if !p.Handler().(*Handler).CanPvP.Load() {
			h.Player.Message("§cYou cannot hit this player, as they do not have a kit equipped")
			ctx.Cancel()
			return
		}

		p.Handler().(*Handler).LastHitBy.Store(h.Player)
		held, _ := h.Player.HeldItems()
		for _, enchantment := range held.Enchantments() {
			id, _ := item.EnchantmentID(enchantment.Type())
			if id < 50 {
				continue
			}
			if _, ok := enchantment.Type().(enchantments.Enchantment); !ok {
				continue
			}
			enchantments.RunEnchantmentCheckOnHit(enchantment.Type().(enchantments.Enchantment), enchantment.Level(), h.Player, p)
		}
		return
	}
	ctx.Cancel()
}

func (h *Handler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()
	_, _ = chat.Global.WriteString(fmt.Sprintf("§f%v§7: §e%v", h.Player.Name(), *message))
}

func (h *Handler) HandleRespawn(*mgl64.Vec3, **world.World) {
	h.CanPvP.Store(false)
	h.Player.Message("§aYou have respawned, use the command §e/gkit §ato re-equip another GKit!")
}
