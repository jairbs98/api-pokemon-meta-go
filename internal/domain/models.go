package domain

import (
	"encoding/json"
	"time"
)

type PokemonUsageStats struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	PokemonID       int             `gorm:"column:pokemon_id" json:"pokemon_id"`
	Name            string          `gorm:"column:name" json:"name"`
	Tier            string          `gorm:"column:tier" json:"tier"`
	Generation      string          `gorm:"column:generation" json:"generation"`
	UsagePercentage float64         `gorm:"column:usage_percentage" json:"usage_percentage"`
	Rank            int             `gorm:"column:rank" json:"rank"`
	BaseStats       json.RawMessage `gorm:"column:base_stats" json:"base_stats"`
	Abilities       json.RawMessage `gorm:"column:abilities" json:"abilities"`
	Items           json.RawMessage `gorm:"column:items" json:"items"`
	Moves           json.RawMessage `gorm:"column:moves" json:"moves"`
	Spreads         json.RawMessage `gorm:"column:spreads" json:"spreads"`
	Natures         json.RawMessage `gorm:"column:natures" json:"natures"`
	Teammates       json.RawMessage `gorm:"column:teammates" json:"teammates"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"created_at"`
}

func (PokemonUsageStats) TableName() string {
	return "pokemon_usage_stats"
}

type MetaSnapshotItem struct {
	Rank            int             `gorm:"column:rank" json:"rank"`
	Name            string          `gorm:"column:name" json:"name"`
	UsagePercentage float64         `gorm:"column:usage_percentage" json:"usage_percentage"`
	TopItem         *string         `gorm:"column:top_item" json:"top_item"`
	TopAbility      *string         `gorm:"column:top_ability" json:"top_ability"`
	Items           json.RawMessage `gorm:"column:items" json:"items"`
	Abilities       json.RawMessage `gorm:"column:abilities" json:"abilities"`
	Moves           json.RawMessage `gorm:"column:moves" json:"moves"`
	Spreads         json.RawMessage `gorm:"column:spreads" json:"spreads"`
	Natures         json.RawMessage `gorm:"column:natures" json:"natures"`
	BaseStats       json.RawMessage `gorm:"column:base_stats" json:"base_stats"`
}

type SpeedCreeperItem struct {
	Name            string  `gorm:"column:name" json:"name"`
	UsagePercentage float64 `gorm:"column:usage_percentage" json:"usage_percentage"`
	Velocidad       int     `gorm:"column:velocidad" json:"velocidad"`
	Ataque          *int    `gorm:"column:ataque" json:"ataque"`
}

type WallbreakerItem struct {
	Name            string  `gorm:"column:name" json:"name"`
	UsagePercentage float64 `gorm:"column:usage_percentage" json:"usage_percentage"`
	AtkBase         int     `gorm:"column:atk_base" json:"atk_base"`
	MoveName        string  `gorm:"column:move_name" json:"move_name"`
}

type TrendPoint struct {
	Day   string  `json:"day"`
	Value float64 `json:"value"`
	Label string  `json:"label"`
}

type PokemonTrendResponse struct {
	PokemonName string       `json:"pokemon_name"`
	History     []TrendPoint `json:"history"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PokemonRepository interface {
	GetMetaSnapshot(tier, gen string, limit int) ([]MetaSnapshotItem, error)
	GetSpeedCreepers(tier, gen string) ([]SpeedCreeperItem, error)
	GetWallbreakers(tier, gen, moveType string) ([]WallbreakerItem, error)
	GetPokemonTrend(pokemonName, tier, gen string) (float64, error)
}
