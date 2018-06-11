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

	c.Archetype.CalculateArchtypeCost()

	c.Body.Dice.Hard = 3
	c.Body.Dice.GoFirst = 1

	c.Skills["Athletics"] = &oneroll.Skill{
		Name:     "Athletics",
		LinkStat: c.Body,
		Dice: &oneroll.DiePool{
			Normal: 3,
			Hard:   0,
			Wiggle: 0,
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
		Type:        "Attack",
		Description: "TK Blast",
		Level:       3,
		CostPerDie:  2,
		Modifiers:   []*oneroll.Modifier{area, ifthen},
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
		Type:        "Useful",
		Description: "Fly",
		Level:       1,
		CostPerDie:  2,
		Modifiers:   []*oneroll.Modifier{gf, boost},
	}

	u.Capacities = []*oneroll.Capacity{
		&speed,
	}

	a.Capacities = []*oneroll.Capacity{
		&rng,
		&mass,
	}

	f.Qualities = []*oneroll.Quality{&a, &u}

	f.CalculatePowerCost()

	f.Effect = "Fly and throw TK blasts."

	c.Powers = map[string]*oneroll.Power{
		"Telekinisis": f}

	c.CalculateCharacterCost()

	fmt.Println(c)

}
