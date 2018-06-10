package oneroll

import "fmt"

// Statistic represents common attributes possessed by every character
type Statistic struct {
	Name string
	Dice *DiePool
	Cost int
}

// HyperStat is a modified version of a regular Statistic
type HyperStat struct {
	Stat       *Statistic
	Capacities []*Capacity
	Modifiers  []*Modifier
	CostPerDie int
	Booster    []*Booster
	Cost       int
}

func (s Statistic) String() string {
	text := fmt.Sprintf("%s: %s",
		s.Name,
		s.Dice)

	return text
}

// CalculateStatCost determines the cost of a Power Quality
// Called from Character.CalculateCharacterCost()
func (s *Statistic) CalculateStatCost() {
	b := 5

	// Temp solution

	b += s.Dice.GoFirst
	b += s.Dice.Spray

	total := b * s.Dice.Normal
	total += b * 2 * s.Dice.Hard
	total += b * 4 * s.Dice.Wiggle

	s.Cost = total
}
