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
	text := fmt.Sprintf("%s", q.Type)

	// Add formatting for additional levels of Quality
	if q.Level > 1 {
		text += fmt.Sprintf(" +%d ", q.Level-1)
	}

	text += fmt.Sprintf("(%s) (%d/die): ",
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