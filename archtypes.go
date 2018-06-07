package oneroll

// Archtype is a grouping of Sources, Permissions & Intrinsics that defines what powers a character can use
type Archtype struct {
	Type        string
	Sources     []*Source
	Permissions []*Permission
	Intrinsics  []*Intrinsic
}

// Source is a source of a Character's powers
type Source struct {
	Type        string
	Cost        int // First source is free
	Description string
}

// Permission is the type of powers a Character can purchase
type Permission struct {
	Type              string
	Cost              int
	Description       string
	AllowHyperSkill   bool
	AllowHyperStat    bool
	ExceedStatLimit   bool
	AllowWiggle       bool
	AllowHard         bool
	AllowMiracles     bool
	AllowGadgeteering bool
	AllowGadgets      bool
	PowerLimit        int
}

// Intrinsic is a modification from the human standard
type Intrinsic struct {
	Name        string
	Cost        int
	Description string
}

// Sources Set Wild Talents Sources
var Sources = map[string]*Source{

	"None":      &Source{Type: "None", Cost: 0, Description: "No Sources"},
	"Construct": &Source{Type: "Construct", Cost: 5, Description: ""},
	"Cyborg":    &Source{Type: "Cyborg", Cost: 5, Description: ""},
	"Divine":    &Source{Type: "Divine", Cost: 5, Description: ""},
	"Driven":    &Source{Type: "Driven", Cost: 5, Description: ""},
	"Extraterrestrial/Extradimensional": &Source{Type: "Extraterrestrial/Extradimensional", Cost: 5, Description: ""},
	"Genetic":                           &Source{Type: "Genetic", Cost: 5, Description: ""},
	"Life Force":                        &Source{Type: "Life Force", Cost: 5, Description: ""},
	"Paranormal":                        &Source{Type: "Paranormal", Cost: 5, Description: ""},
	"Power Focus":                       &Source{Type: "Power Focus", Cost: -8, Description: ""},
	"Psi":                               &Source{Type: "Psi", Cost: 5, Description: ""},
	"Technological":                     &Source{Type: "Technological", Cost: 5, Description: ""},
	"Unknown":                           &Source{Type: "Unknown", Cost: -5, Description: ""},
}

// Permissions sets Wild Talents default permissions
var Permissions = map[string]*Permission{

	"None": &Permission{
		Type:        "None",
		Cost:        0,
		Description: "",
	},
	"Hypertrained": &Permission{
		Type:            "Hypertrained",
		Cost:            5,
		Description:     "",
		AllowHyperSkill: true,
	},
	"Inhuman Stats": &Permission{
		Type:            "Inhuman Stats",
		Cost:            1,
		Description:     "",
		AllowHyperSkill: true,
	},
	"Inventor": &Permission{
		Type:              "Inventor",
		Cost:              5,
		Description:       "",
		AllowGadgeteering: true,
		AllowGadgets:      true,
	},
	"One Power": &Permission{
		Type:          "One Power",
		Cost:          1,
		Description:   "",
		AllowMiracles: true,
		PowerLimit:    1,
	},
	"Peak Performer": &Permission{
		Type:        "Peak Performer",
		Cost:        5,
		Description: "",
		AllowWiggle: true,
		AllowHard:   true,
	},
	"Power Theme": &Permission{
		Type:            "Power Theme",
		Cost:            5,
		Description:     "",
		AllowHyperSkill: true,
		AllowHyperStat:  true,
		AllowMiracles:   true,
		AllowHard:       true,
		AllowWiggle:     true,
	},
	"Prime Specimen": &Permission{
		Type:           "Prime Specimen",
		Cost:           5,
		Description:    "",
		AllowHyperStat: true,
	},
	"Super": &Permission{
		Type:            "Super",
		Cost:            15,
		Description:     "",
		AllowHyperSkill: true,
		AllowHyperStat:  true,
		AllowMiracles:   true,
		AllowHard:       true,
		AllowWiggle:     true,
	},
	"Super Equipment": &Permission{
		Type:         "Super Equipment",
		Cost:         2,
		Description:  "",
		AllowGadgets: true,
	},
}
