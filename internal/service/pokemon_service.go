package service

import (
	"api-pokemon-meta-go/internal/domain"
)

type PokemonService struct {
	repo domain.PokemonRepository
}

func NewPokemonService(repo domain.PokemonRepository) *PokemonService {
	return &PokemonService{repo: repo}
}

func (s *PokemonService) GetMetaSnapshot(tier, gen string, limit int) ([]domain.MetaSnapshotItem, error) {
	return s.repo.GetMetaSnapshot(tier, gen, limit)
}

func (s *PokemonService) GetSpeedCreepers(tier, gen string) ([]domain.SpeedCreeperItem, error) {
	return s.repo.GetSpeedCreepers(tier, gen)
}

func (s *PokemonService) GetWallbreakers(tier, gen, moveType string) ([]domain.WallbreakerItem, error) {
	return s.repo.GetWallbreakers(tier, gen, moveType)
}

func (s *PokemonService) GetPokemonTrend(pokemonName, tier, gen string) (*domain.PokemonTrendResponse, error) {
	history, err := s.repo.GetPokemonTrend(pokemonName, tier, gen)
	if err != nil {
		return nil, err
	}

	if history == nil {
		history = []domain.TrendPoint{}
	}

	return &domain.PokemonTrendResponse{
		PokemonName: pokemonName,
		History:     history,
	}, nil
}

func (s *PokemonService) GetStallIndex(tier, gen string) (*domain.StallIndexResponse, error) {
	index, err := s.repo.GetStallIndex(tier, gen)
	if err != nil {
		return nil, err
	}

	return &domain.StallIndexResponse{
		StallPercentage: index,
	}, nil
}
