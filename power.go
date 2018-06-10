package oneroll

import "fmt"

// Power is a non-standard ability or miracle
type Power struct {
	Name      string
	Qualities []*Quality
	Dice      *DiePool
	Effect    string
	Dud       bool
	Cost      int
}

func (p Power) String() string {
	text := fmt.Sprintf("%s %s (",
		p.Name,
		p.Dice,
	)

	for _, q := range p.Qualities {
		text += fmt.Sprintf("%s", string(q.Type[0]))
		if q.Level > 1 {
			text += fmt.Sprintf("+%d", q.Level-1)
		}
	}

	text += fmt.Sprintf(") %dpts\n", p.Cost)

	for _, q := range p.Qualities {
		text += fmt.Sprintln(q)
	}

	text += fmt.Sprintf("Effect: %s", p.Effect)

	return text
}

// CalculatePowerCost totals the cost of Qualites for a Power
func (p *Power) CalculatePowerCost() {

	b := 0

	for _, q := range p.Qualities {
		for _, m := range q.Modifiers {
			m.CalculateModifierCost()
		}
		q.CalculateQualityCost()
		b += q.CostPerDie
	}

	total := b * p.Dice.Normal
	total += b * 2 * p.Dice.Hard
	total += b * 4 * p.Dice.Wiggle

	p.Cost = total
}

// NewPower generates a new empty Power
func NewPower(t string) *Power {

	p := new(Power)

	p.Name = t
	p.Effect = ""
	p.Qualities = []*Quality{}
	p.Dice = &DiePool{}
	p.Dud = false

	// Take user input

	return p
}
