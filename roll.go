package oneroll

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Roll shows all results and variables from an ORE roll
type Roll struct {
	Actor      *Character
	Action     string // type of action act, oppose, maneuver
	numActions int
	DiePool    *DiePool
	results    []int
	matches    []Match
	loose      []int
	wiggles    int
}

// DiePool represents a rollable dice set in ORE
type DiePool struct {
	Normal  int
	Hard    int
	Wiggle  int
	Spray   int
	GoFirst int
}

// Match shows the height and width of a specific match
type Match struct {
	Actor      *Character
	height     int
	width      int
	initiative int
}

// ByWidthHeight sorts matches in descending order of width then height
type ByWidthHeight []Match

func (a ByWidthHeight) Len() int      { return len(a) }
func (a ByWidthHeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByWidthHeight) Less(i, j int) bool {

	if a[i].initiative > a[j].initiative {
		return true
	}
	if a[i].initiative < a[j].initiative {
		return false
	}

	return a[i].height > a[j].height
}

// Resolve ORE dice roll and prints results
func (r *Roll) Resolve(input string) (*Roll, error) {

	nd, hd, wd, gf, sp, ac, err := r.parseString(input)

	r.numActions = ac

	r.DiePool = &DiePool{
		Normal:  nd,
		Hard:    hd,
		Wiggle:  wd,
		GoFirst: gf,
		Spray:   sp,
	}

	if err != nil {
		return r, err
	}

	r.wiggles = wd

	for x := 0; x < nd; x++ {
		r.results = append(r.results, RollDie(10, 1, 1))
	}

	for x := 0; x < hd; x++ {
		r.results = append(r.results, 10)
	}

	r.parseDieRoll()

	// Sort roll by initiative (width+GoFirst) and then height
	sort.Sort(ByWidthHeight(r.matches))

	return r, nil

}

// parses string like 5d+1hd+1wd or returns error
func (r *Roll) parseString(input string) (int, int, int, int, int, int, error) {

	re := regexp.MustCompile("[0-9]+")

	var sElements []string

	errString := ""

	sElements = strings.SplitN(input, "+", 6)

	var nd, hd, wd, gf, sp int

	ac := 1

	for _, s := range sElements {
		switch {
		case strings.Contains(s, "wd"):
			numString := re.FindString(s)
			wd, _ = strconv.Atoi(numString)

		case strings.Contains(s, "hd"):
			numString := re.FindString(s)
			hd, _ = strconv.Atoi(numString)

		case strings.Contains(s, "d"):
			numString := re.FindString(s)
			nd, _ = strconv.Atoi(numString)

		case strings.Contains(s, "gf"):
			numString := re.FindString(s)
			gf, _ = strconv.Atoi(numString)

		case strings.Contains(s, "sp"):
			numString := re.FindString(s)
			sp, _ = strconv.Atoi(numString)
			nd += sp

		case strings.Contains(s, "ac"):
			numString := re.FindString(s)
			ac, _ = strconv.Atoi(numString)

		default:
			errString = "Error: Not a regular die notation"
		}
	}

	actionCount := ac // Disposable counter

	// Check for multiple actions and spray
	// reduce die pool
	if actionCount > 1 && sp == 0 {
		// Remove hd first
		for hd > 0 && actionCount > 1 {
			hd--
			actionCount--
		}
		for nd > 0 && actionCount > 1 {
			nd--
			actionCount--
		}
		for wd > 0 && actionCount > 1 {
			wd--
			actionCount--
		}
	}

	nd, hd, wd = VerifyLessThan10(nd, hd, wd)

	if errString != "" {
		return 0, 0, 0, 0, 0, 0, errors.New(errString)
	}

	return nd, hd, wd, gf, sp, ac, nil
}

// Determine matches including width, height and initiative for a roll
func (r *Roll) parseDieRoll() *Roll {

	matches := make(map[int]int)
	for _, d := range r.results {
		matches[d]++
	}

	goFirst := 0
	if r.DiePool.GoFirst != 0 { // Error here
		goFirst = r.DiePool.GoFirst
	}

	for k, v := range matches {
		switch {
		case v == 1:
			r.loose = append(r.loose, k)
		case v > 1:
			r.matches = append(r.matches, Match{
				Actor:      r.Actor,
				height:     k,
				width:      v,
				initiative: v + goFirst,
			})
		}
	}
	return r
}

// Provides standard string formatting for roll
func (r Roll) String() string {

	text := ""
	var results []Match

	text += fmt.Sprintf("Actor: %s, Action: %s, Go First: %d, Spray: %d, Wiggle Dice: %dwd\n\n",
		r.Actor.Name,
		r.Action,
		r.DiePool.GoFirst,
		r.DiePool.Spray,
		r.wiggles,
	)

	text += fmt.Sprintf("Dice show: %d\n\n", r.results)

	if len(r.matches) > 0 {

		text += "Matches:\n"

		for _, m := range r.matches {
			results = append(results, m)
		}
		sort.Sort(ByWidthHeight(results))
	}

	text += fmt.Sprintln("***Resolution***")

	text += fmt.Sprintf("%s Actions: %d\n", r.Actor.Name, r.numActions)

	for _, m := range results {
		text += fmt.Sprintf("Match: %dx%d, Initiative: %dx%d\n",
			m.height, m.width,
			m.height, m.initiative,
		)
	}

	if len(r.loose) > 0 {
		text += fmt.Sprintf("\nLoose dice %d\n", r.loose)
	}

	return text + "\n"
}

func main() {

	diePool := flag.String("d", "4d", "a die string separated by + like 4d+2hd+1wd")
	numRolls := flag.Int("n", 1, "an int that represents the number of rolls to make")
	guiOn := flag.Bool("w", true, "Set whether to use the GUI or not (CLI).")

	flag.Parse()

	if *guiOn == true {
		GUI()
	} else {

		for x := 0; x < *numRolls; x++ {
			roll := Roll{}
			roll.Resolve(*diePool)
			fmt.Println(roll)
		}
	}
}
