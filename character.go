package oneroll

import (
	"fmt"
)

// Character represents a full character in the ORE game
type Character struct {
	ID           int64
	Name         string
	Body         *Statistic
	Coordination *Statistic
	Sense        *Statistic
	Mind         *Statistic
	Command      *Statistic
	Charm        *Statistic
	BaseWill     int
	Willpower    int
	Skills       map[string]*Skill
	Archtype     *Archtype
	HyperStats   map[string]*HyperStat
	HyperSkills  map[string]*HyperSkill
	Permissions  map[string]*Permission
	Powers       map[string]*Power
	HitLocations map[string]*Location
	PointCost    int
}

// Display character
func (c *Character) String() string {

	text := fmt.Sprintf("\n%s (%d pts)\n", c.Name, c.PointCost)

	if c.Archtype.Type != "" {
		text += fmt.Sprint(c.Archtype)
	}

	text += "\n\nStats:\n"

	text += ShowSkills(c, false)

	text += fmt.Sprintf("\nBase Will:%d\n", c.BaseWill)
	text += fmt.Sprintf("Willpower: %d\n", c.Willpower)

	text += fmt.Sprintf("\nHit Locations:\n")

	for _, loc := range c.HitLocations {
		text += fmt.Sprintf("%s\n", loc)
	}

	if len(c.Archtype.Sources) > 0 && len(c.Powers) > 0 {
		text += fmt.Sprintf("\nPowers:\n")
		for _, p := range c.Powers {
			text += fmt.Sprintf("%s\n\n", p)
		}
	}

	return text
}

// CalculateCharacterCost sums total costs of all character elements
// Call this on each character update
func (c *Character) CalculateCharacterCost() {

	var cost int

	if len(c.Archtype.Sources) > 0 {
		c.Archtype.CalculateArchtypeCost()
		cost += c.Archtype.Cost
	}

	statistics := []*Statistic{c.Body, c.Coordination, c.Sense, c.Mind, c.Command, c.Charm}

	for _, stat := range statistics {
		stat.CalculateStatCost()
		cost += stat.Cost
	}

	// Hyperstats

	for _, skill := range c.Skills {
		skill.CalculateSkillCost()
		cost += skill.Cost
	}

	// HyperSkills

	for _, power := range c.Powers {
		// Determine power capacities
		power.DeterminePowerCapacities()
		power.CalculatePowerCost()
		cost += power.Cost
	}

	comTotal := c.Command.Dice.Normal + c.Command.Dice.Hard + c.Command.Dice.Wiggle
	charmTotal := c.Charm.Dice.Normal + c.Charm.Dice.Hard + c.Charm.Dice.Wiggle

	cost += 3*c.BaseWill - (comTotal + charmTotal)
	cost += c.Willpower - c.BaseWill

	c.PointCost = cost
}
