package oneroll

import (
	"fmt"
	"strings"
)

// Quality is either Attack, Defend or Useful
type Quality struct {
	Type        string
	Description string
	Level       int
	Capacities  []*Capacity
	Modifiers   []*Modifier
	CostPerDie  int
}

func (q Quality) String() string {
	text := fmt.Sprintf("%s (%s) (%d/die): ",
		q.Type,
		q.Description,
		q.CostPerDie)

	text += fmt.Sprint("Capacities:")

	for _, c := range q.Capacities {
		text += fmt.Sprintf(" %s", c)
	}

	text += fmt.Sprint("; Extras & Flaws:")

	for _, m := range q.Modifiers {
		if m.CostPerLevel > 0 {
			text += fmt.Sprintf(" %s,", m)
		}
	}

	for _, m := range q.Modifiers {
		if m.CostPerLevel < 0 {
			text += fmt.Sprintf(" %s,", m)
		}
	}

	text = strings.TrimSuffix(text, ",")
	return text
}

// NewQuality generates a new empty Quality
func NewQuality(t string) *Quality {

	q := new(Quality)

	q.Type = t
	q.Description = ""
	q.CostPerDie = 2
	q.Level = 1
	q.Capacities = []*Capacity{}
	q.Modifiers = []*Modifier{}

	// Take user input

	return q
}

// CalculateQualityCost determines the cost of a Power Quality
// Called from Power.PowerCost()
func (q *Quality) CalculateQualityCost() {
	b := 2

	b += q.Level - 1

	for _, m := range q.Modifiers {
		if m.RequiresLevel {
			b += m.CostPerLevel * m.Level
		} else {
			b += m.CostPerLevel
		}
	}

	q.CostPerDie = b
}

// Capacity is Range, Mass, Touch or Speed
type Capacity struct {
	Type    string
	Level   int
	Value   string
	Booster *Booster
}

func (c Capacity) String() string {
	text := fmt.Sprintf("%s (%s)",
		c.Type,
		c.Value)

	// Modify value by level & booster

	return text
}

// Booster multiplies a Capacity or Statistic
type Booster struct {
	Level int
}
