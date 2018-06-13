package main

import (
	"fmt"

	"github.com/toferc/oneroll"
)

func main() {
	c := oneroll.NewWTCharacter("Deacon")

	c.Archetype = &oneroll.Archetype{
		Type: "Cadaver",
		Sources: []*oneroll.Source{oneroll.Sources["Genetic"],
			oneroll.Sources["Technological"]},
		Permissions: []*oneroll.Permission{oneroll.Permissions["Super"]},
	}

	oneroll.UpdateCost(c.Archetype)

	// Test HyperStat
	boost1 := oneroll.Modifiers["Booster"]
	boost1.Level = 2

	bgf := oneroll.Modifiers["Go First"]
	bgf.Level = 2

	hbq := oneroll.Quality{
		Type:       "Attack",
		Name:       "Hyper-Body",
		Level:      3,
		CostPerDie: 2,
		Modifiers:  []*oneroll.Modifier{boost1, bgf},
	}

	c.Body.HyperStat = &oneroll.HyperStat{
		Name: "Hyper-Body",
		Dice: &oneroll.DiePool{
			Hard: 3,
		},
		Qualities: []*oneroll.Quality{&hbq},
		Effect:    "Attacks fast and does W+2S.",
	}

	c.Body.HyperStat.Dice.Wiggle = 1

	useful := oneroll.Quality{
		Type:  "Useful",
		Name:  "Hyper-Athletics",
		Level: 1,
	}

	c.Skills["Athletics"] = &oneroll.Skill{
		Name:     "Athletics",
		LinkStat: c.Body,
		Dice: &oneroll.DiePool{
			Normal: 3,
			Hard:   0,
			Wiggle: 0,
		},
		HyperSkill: &oneroll.HyperSkill{
			Name: "Hyper_Athmetics",
			Dice: &oneroll.DiePool{
				Hard: 2,
			},
			Qualities: []*oneroll.Quality{
				&useful,
			},
		},
	}

	f := oneroll.NewPower("Telekinisis")

	f.Dice = &oneroll.DiePool{
		Normal: 4,
		Hard:   2,
	}

	area := oneroll.Modifiers["Area"]
	ifthen := oneroll.Modifiers["If/Then"]
	ifthen.Info = "Only when angry"

	area.Level = 3

	gf := oneroll.Modifiers["Go First"]

	boost := oneroll.Modifiers["Booster"]

	boost.Level = 4

	a := oneroll.Quality{
		Type:       "Attack",
		Name:       "TK Blast",
		Level:      3,
		CostPerDie: 2,
		Modifiers:  []*oneroll.Modifier{area, ifthen},
	}

	rng := oneroll.Capacity{
		Type:  "Range",
		Value: "500m",
	}

	mass := oneroll.Capacity{
		Type: "Mass",
	}

	speed := oneroll.Capacity{
		Type:  "Speed",
		Value: "250kph",
	}

	u := oneroll.Quality{
		Type:       "Useful",
		Name:       "Fly",
		Level:      1,
		CostPerDie: 2,
		Modifiers:  []*oneroll.Modifier{gf, boost},
	}

	u.Capacities = []*oneroll.Capacity{
		&speed,
	}

	a.Capacities = []*oneroll.Capacity{
		&rng,
		&mass,
	}

	f.Qualities = []*oneroll.Quality{&a, &u}

	f.Effect = "Fly and throw TK blasts."

	c.Powers = map[string]*oneroll.Power{
		"Telekinisis": f}

	oneroll.UpdateCost(c)

	fmt.Println(c)

	activePower := c.Powers["Telekinisis"].Qualities[0]

	ds := activePower.FormatDiePool(1)

	actionType := fmt.Sprintf("%s %s", activePower.Type, activePower.Name)

	r := oneroll.Roll{
		Actor:  c,
		Action: actionType,
	}

	r.Resolve(ds)

	fmt.Println(r)

	r1 := oneroll.Roll{
		Actor:  c,
		Action: c.Skills["Athletics"].Name,
	}

	ath := c.Skills["Athletics"]

	athString := ath.FormatDiePool(1)

	r1.Resolve(athString)

	fmt.Println(r1)

}
