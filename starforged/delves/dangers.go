package delves

import "math/rand"

type Danger struct {
	Name string
}

func iterateWildDangers(yield func(Danger) bool) {
	dangers := []Danger{
		Danger{Name: `Denizen hunts`},
		Danger{Name: `Denizen strikes without warning`},
		Danger{Name: `Denizen leverages the environment`},
		Danger{Name: `Denizen wields unexpected abilities`},
		Danger{Name: `Denizen guided by a greater threat`},
		Danger{Name: `Denizen protects something`},
		Danger{Name: `Hazardous terrain`},
		Danger{Name: `Weather or environmental threat`},
		Danger{Name: `Benign aspect becomes a threat`},
		Danger{Name: `Overzealous hunter`},
		Danger{Name: `Disturbing evidence of a victim's fate`},
		Danger{Name: `Ill-fated victim of danger`},
	}
	for {
		if !yield(dangers[rand.Intn(len(dangers))]) {
			return
		}
	}
}

func iterateInfestedDangers(yield func(Danger) bool) {
	dangers := []Danger{
		Danger{Name: `Denizens swarm and attack`},
		Danger{Name: `Toxic or sickening environment`},
		Danger{Name: `Denizen stalks you`},
		Danger{Name: `Denizen takes or destroys something`},
		Danger{Name: `Denizen reveals surprising cleverness`},
		Danger{Name: `Denizen guided by a greater threat`},
		Danger{Name: `Denizen blocks the path`},
		Danger{Name: `Denizen funnels you down a new path`},
		Danger{Name: `Denizen undermines the path`},
		Danger{Name: `Denizen lays in wait`},
		Danger{Name: `Trap or snare`},
		Danger{Name: `Victim's horrible fate is revealed`},
	}
	for {
		if !yield(dangers[rand.Intn(len(dangers))]) {
			return
		}
	}
}

func iterateFortifiedDangers(yield func(Danger) bool) {
	dangers := []Danger{
		Danger{Name: `Denizen patrols the area`},
		Danger{Name: `Denizen on guard`},
		Danger{Name: `Denizen ready to sound the alarm`},
		Danger{Name: `Denizen sets an ambush`},
		Danger{Name: `Denizen lures you into a trap`},
		Danger{Name: `Denizens converge on this area`},
		Danger{Name: `Pets or underlings`},
		Danger{Name: `Unexpected alliance revealed`},
		Danger{Name: `Nefarious plans revealed`},
		Danger{Name: `Unexpected leader revealed`},
		Danger{Name: `Trap`},
		Danger{Name: `Alarm trigger`},
	}
	for {
		if !yield(dangers[rand.Intn(len(dangers))]) {
			return
		}
	}
}

func iterateHauntedDangers(yield func(Danger) bool) {
	dangers := []Danger{
		Danger{Name: `Denizen haunts this area`},
		Danger{Name: `Unsettling sounds or foreboding signs`},
		Danger{Name: `Denizen attacks without warning`},
		Danger{Name: `Denizen makes a costly demand`},
		Danger{Name: `Denizen seizes your body or mind`},
		Danger{Name: `Denizen taunts or lures you`},
		Danger{Name: `A disturbing truth is revealed`},
		Danger{Name: `Frightening visions`},
		Danger{Name: `The environment is used against you`},
		Danger{Name: `Trickery leads you astray`},
		Danger{Name: `True nature of this place is revealed`},
		Danger{Name: `Sudden, shocking manifestation`},
	}
	for {
		if !yield(dangers[rand.Intn(len(dangers))]) {
			return
		}
	}
}
