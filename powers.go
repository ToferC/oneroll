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
	text := fmt.Sprintf("%s %s (%dpts)\n",
		p.Name,
		p.Dice,
		p.Cost,
	)
	for _, q := range p.Qualities {
		text += fmt.Sprintln(q)
	}
	text += fmt.Sprintln(p.Effect)

	return text
}

// PowerCost totals the cost of Qualites for a Power
func (p *Power) PowerCost() int {
	b := 0

	for _, q := range p.Qualities {
		b += q.CostPerDie
	}

	total := b * p.Dice.Normal
	total += b * 2 * p.Dice.Hard
	total += b * 4 * p.Dice.Wiggle

	return total
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

func (q Quality) String() string {
	text := fmt.Sprintf("%s %s (%d/die)\n",
		q.Type,
		q.Description,
		q.CostPerDie)

	for _, c := range q.Capacities {
		text += fmt.Sprintln("Capacities: ", c)
	}

	for _, e := range q.Extras {
		text += fmt.Sprintln("Extras: ", e)
	}

	for _, f := range q.Flaws {
		text += fmt.Sprintln("Flaws: ", f)
	}

	return text
}

// QualityCost determines the cost of a Power Quality
func (q *Quality) QualityCost() int {
	b := 2

	b += q.Level - 1

	for _, e := range q.Extras {
		if e.RequiresLevel {
			b += e.CostModifierPerLevel * e.Level
		} else {
			b += e.CostModifierPerLevel
		}
	}

	for _, f := range q.Flaws {
		if f.RequiresLevel {
			b += f.CostModifierPerLevel * f.Level
		} else {
			b += f.CostModifierPerLevel
		}
	}

	return b
}

// Capacity is Range, Mass, Touch or Speed
type Capacity struct {
	Type    string
	Level   int
	Value   string
	Booster *Booster
}

func (c Capacity) String() string {
	text := fmt.Sprintf("%s (%s)\n",
		c.Type,
		c.Value)

	// Modify value by level & booster

	return text
}

// Booster multiplies a Capacity or Statistic
type Booster struct {
	Level int
}

// Extra enhances the abilities of a Power Quality
type Extra struct {
	Name                 string
	Description          string
	RequiresLevel        bool
	Level                int
	RequiresInfo         bool
	Info                 string
	CostModifierPerLevel int
}

func (e Extra) String() string {
	text := fmt.Sprintf("%s", e.Name)

	if e.RequiresLevel {
		text += fmt.Sprintf(" (%d)", e.Level)
	}

	if e.RequiresInfo {
		text += fmt.Sprintf(" - %s)", e.Info)
	}

	if e.CostModifierPerLevel > 0 {
		text += fmt.Sprintf(" (+%d/die)", e.CostModifierPerLevel)
	} else {
		text += fmt.Sprintf(" (-%d/die)", e.CostModifierPerLevel)
	}

	return text
}

func (f Flaw) String() string {
	text := fmt.Sprintf("%s", f.Name)

	if f.RequiresLevel {
		text += fmt.Sprintf(" (%d)", f.Level)
	}

	if f.RequiresInfo {
		text += fmt.Sprintf(" - %s)", f.Info)
	}

	if f.CostModifierPerLevel > 0 {
		text += fmt.Sprintf(" (+%d/die)", f.CostModifierPerLevel)
	} else {
		text += fmt.Sprintf(" (-%d/die)", f.CostModifierPerLevel)
	}

	return text
}

// Flaw limits the abilities of a Power Quality
type Flaw struct {
	Name                 string
	Description          string
	RequiresLevel        bool
	Level                int
	RequiresInfo         bool
	Info                 string
	CostModifierPerLevel int
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

// Extras creates map of standard WT extras
var Extras = map[string]*Extra{
	"Area": &Extra{
		Name:                 "Area",
		Description:          "",
		RequiresLevel:        true,
		Level:                1,
		RequiresInfo:         false,
		Info:                 "",
		CostModifierPerLevel: 1,
	},
	"Augment": &Extra{
		Name:                 "Augment",
		Description:          "",
		RequiresLevel:        false,
		CostModifierPerLevel: 4,
	},
	"Booster": &Extra{
		Name:                 "Booster",
		Description:          "",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 1,
	},
	"Burn": &Extra{
		Name:                 "Burn",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Controlled Effect": &Extra{
		Name:                 "Controlled Effect",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Daze": &Extra{
		Name:                 "Daze",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Deadly": &Extra{
		Name:                 "Deadly",
		Description:          "1: Killing, 2: Shock & Killing",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 1,
	},
	"Disintigrate": &Extra{
		Name:                 "Disintigrate",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Duration": &Extra{
		Name:                 "Duration",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Electrocuting": &Extra{
		Name:                 "Electrocuting",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Endless": &Extra{
		Name:                 "Endless",
		Description:          "",
		CostModifierPerLevel: 3,
	},
	"Engulf": &Extra{
		Name:                 "Engulf",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Go First": &Extra{
		Name:                 "Go First",
		Description:          "",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 1,
	},
	"Hardened Defense": &Extra{
		Name:                 "Hardened Defense",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"High Capacity": &Extra{
		Name:                 "High Capacity",
		Description:          "",
		RequiresInfo:         true,
		Info:                 "",
		CostModifierPerLevel: 1,
	},
	"Interference": &Extra{
		Name:                 "Interference",
		Description:          "",
		CostModifierPerLevel: 3,
	},
	"Native Power": &Extra{
		Name:                 "Native Power",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"No Physics": &Extra{
		Name:                 "No Physics",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"No Upward Limit": &Extra{
		Name:                 "No Upward Limit",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Non-Physical": &Extra{
		Name:                 "Non-Physical",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"On Sight": &Extra{
		Name:                 "On Sight",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Penetration": &Extra{
		Name:                 "Penetration",
		Description:          "",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 1,
	},
	"Permanent": &Extra{
		Name:                 "Permanent",
		Description:          "",
		CostModifierPerLevel: 4,
	},
	"Radius": &Extra{
		Name:                 "Radius",
		Description:          "10m x2/level",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 2,
	},
	"Power Capacity": &Extra{
		Name:                 "Power Capacity",
		Description:          "Power Capacity Type",
		RequiresLevel:        true,
		Level:                1,
		RequiresInfo:         false,
		Info:                 "Mass, Range, Speed or Touch",
		CostModifierPerLevel: 1,
	},
	"Speeding Bullet": &Extra{
		Name:                 "Speeding Bullet",
		Description:          "",
		CostModifierPerLevel: 2,
	},
	"Spray": &Extra{
		Name:                 "Spray",
		Description:          "",
		RequiresLevel:        true,
		Level:                1,
		CostModifierPerLevel: 1,
	},
	"Subtle": &Extra{
		Name:                 "Subtle",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Traumatic": &Extra{
		Name:                 "Traumatic",
		Description:          "",
		CostModifierPerLevel: 1,
	},
	"Variable Effect": &Extra{
		Name:                 "Variable Effect",
		Description:          "",
		CostModifierPerLevel: 4,
	},
}
