package delves

import "math/rand"

type Feature struct {
	Name string
}

func iterateWildFeatures(yield func(Feature) bool) {
	features := []Feature{
		Feature{Name: `Denizen's Lair`},
		Feature{Name: `Territorial markings`},
		Feature{Name: `Impressive flora or fauna`},
		Feature{Name: `Hunting ground or watering hole`},
		Feature{Name: `Remains or carrion`},
	}
	for {
		index := rand.Intn(len(features))
		if !yield(features[index]) {
			return
		}
	}
}
