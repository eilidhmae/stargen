package delves

import (
	"iter"
	"math/rand"
)

type Theme struct {
	Name        string
	Description string
	Features    func(yield func(Feature) bool)
	Dangers     func(yield func(Danger) bool)
}

func NewTheme() Theme {
	themes := []Theme{
		Theme{
			Name:        `Wild`,
			Description: `Nature prevails in this place.`,
			Features:    iterateWildFeatures,
			Dangers:     iterateWildDangers,
		},
		Theme{
			Name: `Infested`,
			Description: `Foul creatures dwell here.`,
			Features: iterateInfestedFeatures,
			Dangers: iterateInfestedDangers,
		},
	}
	return themes[rand.Intn(len(themes))]
}

func (t *Theme) GetFeatureString() string {
	next, stop := iter.Pull(t.Features)
	defer stop()
	v, ok := next()
	if !ok {
		return ""
	}
	return v.Name
}

func (t *Theme) GetDangerString() string {
	next, stop := iter.Pull(t.Dangers)
	defer stop()
	v, ok := next()
	if !ok {
		return ""
	}
	return v.Name
}
