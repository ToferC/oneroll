package main

import (
	"fmt"

	"github.com/toferc/oneroll"
)

func main() {
	c := oneroll.NewWTCharacter("Deacon")

	c.Archtype = &oneroll.Archtype{
		Type:        "Cadaver",
		Sources:     []*oneroll.Source{oneroll.Sources["Genetic"]},
		Permissions: []*oneroll.Permission{oneroll.Permissions["Super"]},
	}

	f := oneroll.NewPower("Telekinisis")

	f.Dice = &oneroll.DiePool{
		Normal: 4,
		Hard:   2,
	}

	e := oneroll.Modifiers["Area"]

	e.Level = 3
	e.Cost = e.ModifierCost()

	e2 := oneroll.Modifiers["Go First"]
	e2.Cost = e2.ModifierCost()

	a := oneroll.Quality{
		Type:        "Attack",
		Description: "TK Blast",
		Level:       3,
		CostPerDie:  2,
		Modifiers:   []*oneroll.Modifier{e},
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
		Modifiers:   []*oneroll.Modifier{e2},
	}

	u.Capacities = []*oneroll.Capacity{
		&speed,
	}

	a.Capacities = []*oneroll.Capacity{
		&rng,
	}

	a.CostPerDie = a.QualityCost()
	u.CostPerDie = u.QualityCost()

	f.Qualities = []*oneroll.Quality{&a, &u}

	f.Cost = f.PowerCost()

	f.Effect = "Fly and throw TK blasts."

	c.Powers = map[string]*oneroll.Power{
		"Telekinisis": f}

	fmt.Println(c)
	fmt.Println(c.Powers["Telekinisis"])
}
