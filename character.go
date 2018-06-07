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
	Archtypes    map[string]*Archtype
	HyperStats   map[string]*HyperStat
	HyperSkills  map[string]*HyperSkill
	Permissions  map[string]*Permission
	Powers       map[string]*Power
	HitLocations map[string]*Location
	PointCost    int
}

// Display character
func (c *Character) String() string {

	text := fmt.Sprintf("\n%s\n\nStats:\n", c.Name)

	text += ShowSkills(c, false)

	text += fmt.Sprintf("\nBase Will:%d\n", c.BaseWill)
	text += fmt.Sprintf("Willpower: %d\n", c.Willpower)

	text += fmt.Sprintf("\nHit Locations:\n")

	for _, loc := range c.HitLocations {
		text += fmt.Sprintf("%s\n", loc)
	}
	return text
}
