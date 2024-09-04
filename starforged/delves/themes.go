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
			Name:        `Infested`,
			Description: `Foul creatures dwell here.`,
			Features:    iterateInfestedFeatures,
			Dangers:     iterateInfestedDangers,
		},
		Theme{
			Name:        `Fortified`,
			Description: `Foes defend this place against intruders.`,
			Features:    iterateFortifiedFeatures,
			Dangers:     iterateFortifiedDangers,
		},
		Theme{
			Name:        `Haunted`,
			Description: `Restless spirits are bound to this place.`,
			Features:    iterateHauntedFeatures,
			Dangers:     iterateHauntedDangers,
		},
		Theme{
			Name:        `Hallowed`,
			Description: `The faithful worship here.`,
			Features:    iterateHallowedFeatures,
			Dangers:     iterateHallowedDangers,
		},
		Theme{
			Name:        `Ravaged`,
			Description: `Time, disaster, or strife have taken their toll.`,
			Features:    iterateRavagedFeatures,
			Dangers:     iterateRavagedDangers,
		},
		Theme{
			Name:        `Corrupted`,
			Description: `This place is tainted by dark magic.`,
			Features:    iterateCorruptedFeatures,
			Dangers:     iterateCorruptedDangers,
		},
		Theme{
			Name:        `Ancient`,
			Description: `This place holds the secrets of a bygone age.`,
			Features:    iterateAncientFeatures,
			Dangers:     iterateAncientDangers,
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
