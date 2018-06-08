package oneroll

import "fmt"

// Statistic represents common attributes possessed by every character
type Statistic struct {
	Name    string
	Dice    *DiePool
	Booster []*Booster
}

// HyperStat is a modified version of a regular Statistic
type HyperStat struct {
	Name       string
	Dice       *DiePool
	Capacities []*Capacity
	Extras     []*Extra
	Flaws      []*Flaw
	CostPerDie int
	Booster    []*Booster
}

func (s Statistic) String() string {
	text := fmt.Sprintf("%s: %s",
		s.Name,
		s.Dice)

	return text
}
