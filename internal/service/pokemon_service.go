package service

import (
	"math"
	"math/rand"
	"time"

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
	currentUsage, err := s.repo.GetPokemonTrend(pokemonName, tier, gen)
	if err != nil {
		return nil, err
	}

	history := make([]domain.TrendPoint, 0)
	daysMap := map[time.Weekday]string{
		time.Monday:    "Lun",
		time.Tuesday:   "Mar",
		time.Wednesday: "Mié",
		time.Thursday:  "Jue",
		time.Friday:    "Vie",
		time.Saturday:  "Sáb",
		time.Sunday:    "Dom",
	}

	today := time.Now()

	for i := 6; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		variance := 0.9 + rand.Float64()*(1.1-0.9)
		simulatedValue := currentUsage * variance

		if i == 0 {
			simulatedValue = currentUsage
		}

		history = append(history, domain.TrendPoint{
			Day:   daysMap[date.Weekday()],
			Value: math.Round(simulatedValue*100) / 100,
			Label: date.Weekday().String(),
		})
	}

	return &domain.PokemonTrendResponse{
		PokemonName: pokemonName,
		History:     history,
	}, nil
}
