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

	f := new(oneroll.Power)
	f.NewPower("Flight")

	f.Dice = &oneroll.DiePool{
		Normal: 4,
		Hard:   2,
	}

	e := oneroll.Extras["Area"]
	e2 := oneroll.Extras["Go First"]

	u := oneroll.Quality{
		Type:        "Useful",
		Description: "Fly",
		Level:       1,
		CostPerDie:  2,
		Extras:      []*oneroll.Extra{e, e2},
	}

	u.CostPerDie = u.QualityCost()

	f.Qualities = []*oneroll.Quality{&u}

	f.Cost = f.PowerCost()

	c.Powers = map[string]*oneroll.Power{
		"Flight": f}

	fmt.Println(c)
	fmt.Println(c.Powers["Flight"])
}
