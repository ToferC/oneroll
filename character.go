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

// Statistic represents common attributes possessed by every character
type Statistic struct {
	Name    string
	Dice    *DiePool
	Booster []*Booster
}

// Skill represents specific training
type Skill struct {
	Name           string
	LinkStat       *Statistic
	Dice           *DiePool
	ReqSpec        bool
	Specialization string
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

// HyperSkill is a modified version of a regular Skill
type HyperSkill struct {
	Name       string
	LinkStat   *Statistic
	Dice       *DiePool
	Capacities []*Capacity
	Extras     []*Extra
	Flaws      []*Flaw
	CostPerDie int
}

// Archtype is a grouping of Sources, Permissions & Intrinsics that defines what powers a character can use
type Archtype struct {
	Sources     []*Source
	Permissions []*Permission
	Intrinsics  []*Intrinsic
}

// Source is a source of a Character's powers
type Source struct {
	Type        string
	Cost        int
	Description string
}

// Permission is the type of powers a Character can purchase
type Permission struct {
	Type        string
	Cost        int
	Description string
}

// Intrinsic is a modification from the human standard
type Intrinsic struct {
	Name        string
	Cost        int
	Description string
}

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
	Type       string
	Level      int
	Capacities []*Capacity
	Extras     []*Extra
	Flaws      []*Flaw
	CostPerDie int
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

// Location represents a body area that can take damage
type Location struct {
	Name     string
	HitLoc   []int
	Boxes    int
	Stun     int
	Kill     int
	LAR      int
	HAR      int
	Disabled bool
}

// Strings
func (l Location) String() string {
	text := fmt.Sprintf("%s - %s: Boxes: %d",
		TrimSliceBrackets(l.HitLoc),
		l.Name,
		l.Boxes,
	)

	if l.LAR > 0 {
		text += fmt.Sprintf(" LAR %d", l.LAR)
	}

	if l.HAR > 0 {
		text += fmt.Sprintf(" HAR %d", l.HAR)
	}

	if l.Kill > 0 {
		text += fmt.Sprintf(" Kill %d", l.Kill)
	}

	if l.Stun > 0 {
		text += fmt.Sprintf(" Stun %d", l.Stun)
	}
	return text
}

func (s Statistic) String() string {
	text := fmt.Sprintf("%s: %dd",
		s.Name,
		s.Dice.Normal,
	)

	if s.Dice.Hard > 0 {
		text += fmt.Sprintf("+%dhd", s.Dice.Hard)
	}

	if s.Dice.Wiggle > 0 {
		text += fmt.Sprintf("+%dwd", s.Dice.Wiggle)
	}

	if s.Dice.GoFirst > 0 {
		text += fmt.Sprintf(" Go First %d", s.Dice.GoFirst)
	}

	if s.Dice.Spray > 0 {
		text += fmt.Sprintf(" Spray %d", s.Dice.Spray)
	}

	return text
}

func (s Skill) String() string {

	text := fmt.Sprintf("%s ",
		s.Name)

	if s.ReqSpec {
		text += fmt.Sprintf("[%s] ", s.Specialization)
	}

	text += fmt.Sprintf("(%s): %dd",
		s.LinkStat.Name,
		s.Dice.Normal,
	)

	if s.Dice.Hard > 0 {
		text += fmt.Sprintf("+%dhd", s.Dice.Hard)
	}

	if s.Dice.Wiggle > 0 {
		text += fmt.Sprintf("+%dwd", s.Dice.Wiggle)
	}

	if s.Dice.GoFirst > 0 {
		text += fmt.Sprintf(" Go First %d", s.Dice.GoFirst)
	}

	if s.Dice.Spray > 0 {
		text += fmt.Sprintf(" Spray %d", s.Dice.Spray)
	}

	return text
}

// NewWTCharacter generates an ORE WT character
func NewWTCharacter(name string) *Character {

	c := Character{
		Name: name,
	}

	c.Body = &Statistic{
		Name: "Body",
		Dice: &DiePool{
			Normal:  2,
			Hard:    0,
			GoFirst: 0,
		},
	}

	c.Coordination = &Statistic{
		Name: "Coordination",
		Dice: &DiePool{
			Normal: 2,
		},
	}
	c.Sense = &Statistic{
		Name: "Sense",
		Dice: &DiePool{
			Normal: 2,
		},
	}
	c.Mind = &Statistic{
		Name: "Mind",
		Dice: &DiePool{
			Normal: 2,
		},
	}
	c.Command = &Statistic{
		Name: "Command",
		Dice: &DiePool{
			Normal: 2,
		},
	}
	c.Charm = &Statistic{
		Name: "Charm",
		Dice: &DiePool{
			Normal: 2,
		},
	}

	c.HitLocations = map[string]*Location{
		"Head": &Location{
			Name:     "Head",
			HitLoc:   []int{10},
			Boxes:    4,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
		"Body": &Location{
			Name:     "Body",
			HitLoc:   []int{7, 8, 9},
			Boxes:    10,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
		"Left Arm": &Location{
			Name:     "Left Arm",
			HitLoc:   []int{5, 6},
			Boxes:    6,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
		"Right Arm": &Location{
			Name:     "Right Arm",
			HitLoc:   []int{3, 4},
			Boxes:    6,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
		"Left Leg": &Location{
			Name:     "Left Leg",
			HitLoc:   []int{2},
			Boxes:    6,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
		"Right Leg": &Location{
			Name:     "Right Leg",
			HitLoc:   []int{1},
			Boxes:    6,
			Stun:     0,
			Kill:     0,
			LAR:      0,
			HAR:      0,
			Disabled: false,
		},
	}

	c.BaseWill = c.Command.Dice.Normal + c.Charm.Dice.Normal
	c.Willpower = c.BaseWill

	c.Skills = map[string]*Skill{
		// Body Skills
		"Athletics": &Skill{
			Name:     "Athletics",
			LinkStat: c.Body,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
				Wiggle: 0,
			},
		},
		"Block": &Skill{
			Name:     "Block",
			LinkStat: c.Body,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
				Wiggle: 0,
			},
		},
		"Brawling": &Skill{
			Name:     "Brawling",
			LinkStat: c.Body,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
				Wiggle: 0,
			},
		},
		"Endurance": &Skill{
			Name:     "Endurance",
			LinkStat: c.Body,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
				Wiggle: 0,
			},
		},
		"Melee Weapon": &Skill{
			Name:     "Melee Weapon",
			LinkStat: c.Body,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
				Wiggle: 0,
			},
			ReqSpec:        true,
			Specialization: "Sword",
		},
		// Coordination Skills
		"Dodge": &Skill{
			Name:     "Dodge",
			LinkStat: c.Coordination,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
			},
		},
		"Driving": &Skill{
			Name:     "Driving",
			LinkStat: c.Coordination,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
			},
			ReqSpec:        true,
			Specialization: "Ground",
		},
		"Ranged Weapon": &Skill{
			Name:     "Ranged Weapon",
			LinkStat: c.Coordination,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
			},
			ReqSpec:        true,
			Specialization: "Pistol",
		},
		"Stealth": &Skill{
			Name:     "Stealth",
			LinkStat: c.Coordination,
			Dice: &DiePool{
				Normal: 0,
				Hard:   0,
			},
		},
		// Sense Skills
		"Empathy": &Skill{
			Name:     "Empathy",
			LinkStat: c.Sense,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Perception": &Skill{
			Name:     "Perception",
			LinkStat: c.Sense,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Scrutiny": &Skill{
			Name:     "Scrutiny",
			LinkStat: c.Sense,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		// Mind Skills
		"First Aid": &Skill{
			Name:     "First Aid",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Knowledge": &Skill{
			Name:     "Knowledge",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
			ReqSpec:        true,
			Specialization: "Alchemy",
		},
		"Languages": &Skill{
			Name:     "Languages",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
			ReqSpec:        true,
			Specialization: "Chinese",
		},
		"Medicine": &Skill{
			Name:     "Medicine",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Navigation": &Skill{
			Name:     "Navigation",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Research": &Skill{
			Name:     "Research",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Security Systems": &Skill{
			Name:     "Security Systems",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Streetwise": &Skill{
			Name:     "Streetwise",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Survival": &Skill{
			Name:     "Survival",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Tactics": &Skill{
			Name:     "Tactics",
			LinkStat: c.Mind,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		// Charm Skills
		"Lie": &Skill{
			Name:     "Lie",
			LinkStat: c.Charm,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Performance": &Skill{
			Name:     "Performance",
			LinkStat: c.Charm,
			Dice: &DiePool{
				Normal: 0,
			},
			ReqSpec:        true,
			Specialization: "Standup",
		},
		"Persuasion": &Skill{
			Name:     "Persuasion",
			LinkStat: c.Charm,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		// Command Skills
		"Interrogation": &Skill{
			Name:     "Interrogation",
			LinkStat: c.Command,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Intimidation": &Skill{
			Name:     "Intimidation",
			LinkStat: c.Command,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Leadership": &Skill{
			Name:     "Leadership",
			LinkStat: c.Command,
			Dice: &DiePool{
				Normal: 0,
			},
		},
		"Stability": &Skill{
			Name:     "Stability",
			LinkStat: c.Command,
			Dice: &DiePool{
				Normal: 0,
			},
		},
	}

	return &c
}

// Display character
func (c *Character) Display() {

	fmt.Println(c.Name)

	stats := []*Statistic{c.Body, c.Coordination, c.Sense,
		c.Mind, c.Command, c.Charm}

	for _, stat := range stats {
		fmt.Println(stat)
	}

	fmt.Println("\nBase Will: ", c.BaseWill)
	fmt.Println("Willpower: ", c.Willpower)

	fmt.Println("\nSkills:")

	for _, skill := range c.Skills {
		if SkillRated(skill) {
			fmt.Println(skill)
		}
	}

	fmt.Println("\nHit Locations:")

	for _, loc := range c.HitLocations {
		fmt.Println(loc)
	}
}
