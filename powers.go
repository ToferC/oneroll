package oneroll

// Power is a non-standard ability or miracle
type Power struct {
	Name      string
	Qualities []*Quality
	Dice      *DiePool
	Effect    string
	Dud       bool
}

// Quality is either Attack, Defend or Useful
type Quality struct {
	Type        string
	Description string
	Level       int
	Capacities  []*Capacity
	Extras      []*Extra
	Flaws       []*Flaw
	CostPerDie  int
}

// Capacity is Range, Mass, Touch or Speed
type Capacity struct {
	Type    string
	Level   int
	Booster *Booster
}

// Booster multiplies a Capacity or Statistic
type Booster struct {
	Level int
}

// Extra enhances the abilities of a Power Quality
type Extra struct {
	Name     string
	Modifier int
}

// Flaw limits the abilities of a Power Quality
type Flaw struct {
	Name     string
	Modifier int
}

// NewPower generates a new empty Power
func (p *Power) NewPower(t string) *Power {
	p.Name = t
	p.Effect = ""
	p.Qualities = []*Quality{}
	p.Dice = &DiePool{}
	p.Dud = false

	// Take user input

	return p
}

// NewQuality generates a new empty Quality
func (q *Quality) NewQuality(t string) *Quality {
	q.Type = t
	q.Description = ""
	q.CostPerDie = 2
	q.Level = 0
	q.Capacities = []*Capacity{}
	q.Extras = []*Extra{}
	q.Flaws = []*Flaw{}

	// Take user input

	return q
}
