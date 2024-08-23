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
