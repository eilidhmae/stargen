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

func iterateInfestedFeatures(yield func(Feature) bool) {
	features := []Feature{
		Feature{Name: `Inhabited nest`},
		Feature{Name: `Abandoned nest`},
		Feature{Name: `Ravaged terrain or architecture`},
		Feature{Name: `Remains or carrion`},
		Feature{Name: `Hoarded food`},
	}
	for {
		index := rand.Intn(len(features))
		if !yield(features[index]) {
			return
		}
	}
}

func iterateFortifiedFeatures(yield func(Feature) bool) {
	features := []Feature{
		Feature{Name: `Camp or quarters`},
		Feature{Name: `Guarded location`},
		Feature{Name: `Storage or repository`},
		Feature{Name: `Work or training area`},
		Feature{Name: `Command center or leadership`},
	}
	for {
		index := rand.Intn(len(features))
		if !yield(features[index]) {
			return
		}
	}
}

func iterateHauntedFeatures(yield func(Feature) bool) {
	features := []Feature{
		Feature{Name: `Tomb or burial site`},
		Feature{Name: `Blood was spilled here`},
		Feature{Name: `Unnatural mists or darkness`},
		Feature{Name: `Messages from beyond the grave`},
		Feature{Name: `Apparitions of a person or event`},
	}
	for {
		index := rand.Intn(len(features))
		if !yield(features[index]) {
			return
		}
	}
}

func iterateHallowedFeatures(yield func(Feature) bool) {
	features := []Feature{
		Feature{Name: `Temple or altar`},
		Feature{Name: `Offerings or atonements`},
		Feature{Name: `Religious relic or idol`},
		Feature{Name: `Consecrated ground`},
		Feature{Name: `Dwellings or gathering place`},
	}
	for {
		index := rand.Intn(len(features))
		if !yield(features[index]) {
			return
		}
	}
}
