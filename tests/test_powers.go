package main

import (
	"fmt"

	"github.com/toferc/oneroll"
)

func main() {
	c := oneroll.NewWTCharacter("Deacon")

	c.Archtype = &oneroll.Archtype{
		Type: "Cadaver",
		Sources: []*oneroll.Source{oneroll.Sources["Genetic"],
			oneroll.Sources["Technological"]},
		Permissions: []*oneroll.Permission{oneroll.Permissions["Super"]},
	}

	c.Archtype.CalculateArchtypeCost()

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

	speed := oneroll.Capacity{
		Type:  "Speed",
		Value: "250kph",
	}

	u := oneroll.Quality{
		Type:        "Useful",
		Description: "Fly",
		Level:       1,
		CostPerDie:  2,
		Modifiers:   []*oneroll.Modifier{gf},
	}

	u.Capacities = []*oneroll.Capacity{
		&speed,
	}

	a.Capacities = []*oneroll.Capacity{
		&rng,
	}

	f.Qualities = []*oneroll.Quality{&a, &u}

	f.CalculatePowerCost()

	f.Effect = "Fly and throw TK blasts."

	c.Powers = map[string]*oneroll.Power{
		"Telekinisis": f}

	fmt.Println(c)

}
