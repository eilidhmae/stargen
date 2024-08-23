package delves

import "iter"

type Theme struct {
	Name        string
	Description string
	Features    func(yield func(Feature) bool)
	Dangers     func(yield func(Danger) bool)
}

func NewTheme() Theme {
	theme := Theme{
		Name:        `Wild`,
		Description: `Nature prevails in this place.`,
		Features:    iterateWildFeatures,
		Dangers:     iterateWildDangers,
	}
	return theme
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
