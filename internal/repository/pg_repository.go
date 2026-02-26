package repository

import (
	"api-pokemon-meta-go/internal/domain"

	"gorm.io/gorm"
)

type pgPokemonRepository struct {
	db *gorm.DB
}

func NewPostgresPokemonRepository(db *gorm.DB) domain.PokemonRepository {
	return &pgPokemonRepository{db: db}
}

func (r *pgPokemonRepository) GetMetaSnapshot(tier, gen string, limit int) ([]domain.MetaSnapshotItem, error) {
	var results []domain.MetaSnapshotItem
	query := `
		SELECT 
			rank, name, usage_percentage, 
			items->>0 as top_item, abilities->>0 as top_ability,
			items, abilities, moves, spreads, natures, base_stats
		FROM pokemon_usage_stats 
		WHERE tier = ? AND generation = ? 
		ORDER BY rank ASC LIMIT ?;
	`
	err := r.db.Raw(query, tier, gen, limit).Scan(&results).Error
	return results, err
}

func (r *pgPokemonRepository) GetSpeedCreepers(tier, gen string) ([]domain.SpeedCreeperItem, error) {
	var results []domain.SpeedCreeperItem
	query := `
		WITH TopSpeed AS (
			SELECT MAX((SELECT value->>'base_stat' FROM jsonb_array_elements(base_stats::jsonb) WHERE value->'stat'->>'name' = 'speed')::int) as max_meta_speed
			FROM pokemon_usage_stats WHERE tier = ? AND generation = ? AND rank <= 10
		)
		SELECT 
			name, usage_percentage,
			(SELECT value->>'base_stat' FROM jsonb_array_elements(base_stats::jsonb) WHERE value->'stat'->>'name' = 'speed')::int as velocidad,
			(SELECT value->>'base_stat' FROM jsonb_array_elements(base_stats::jsonb) WHERE value->'stat'->>'name' = 'attack')::int as ataque
		FROM pokemon_usage_stats, TopSpeed
		WHERE tier = ? AND generation = ?
		  AND (SELECT value->>'base_stat' FROM jsonb_array_elements(base_stats::jsonb) WHERE value->'stat'->>'name' = 'speed')::int > TopSpeed.max_meta_speed
		  AND usage_percentage < 5.0
		ORDER BY velocidad DESC LIMIT 5;
	`
	err := r.db.Raw(query, tier, gen, tier, gen).Scan(&results).Error
	return results, err
}

func (r *pgPokemonRepository) GetWallbreakers(tier, gen, moveType string) ([]domain.WallbreakerItem, error) {
	var results []domain.WallbreakerItem
	query := `
		SELECT DISTINCT
			p.name, p.usage_percentage,
			(SELECT value->>'base_stat' FROM jsonb_array_elements(p.base_stats::jsonb) WHERE value->'stat'->>'name' = 'attack')::int as atk_base,
			m->>'name' as move_name
		FROM pokemon_usage_stats p, jsonb_array_elements(p.moves::jsonb) m
		WHERE p.tier = ? AND p.generation = ?
		  AND m->>'type' = ? AND (m->>'power')::int >= 75 AND p.usage_percentage < 15.0
		ORDER BY atk_base DESC LIMIT 10;
	`
	err := r.db.Raw(query, tier, gen, moveType).Scan(&results).Error
	return results, err
}

func (r *pgPokemonRepository) GetPokemonTrend(pokemonName, tier, gen string) (float64, error) {
	var usage float64
	query := `SELECT usage_percentage FROM pokemon_usage_stats WHERE name = ? AND tier = ? AND generation = ?`
	err := r.db.Raw(query, pokemonName, tier, gen).Scan(&usage).Error
	return usage, err
}
