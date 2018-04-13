package oneroll

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Roll shows all results and variables from an ORE roll
type Roll struct {
	results []int
	matches []Match
	loose   []int
	wiggles int
}

// Match shows the height and width of a specific match
type Match struct {
	height int
	width  int
}

// ByWidthHeight sorts matches in descending order of width then height
type ByWidthHeight []Match

func (a ByWidthHeight) Len() int      { return len(a) }
func (a ByWidthHeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByWidthHeight) Less(i, j int) bool {

	if a[i].width > a[j].width {
		return true
	}
	if a[i].width < a[j].width {
		return false
	}

	return a[i].height > a[j].height
}

// Provides standard string formatting for roll
func (r Roll) String() string {
	text := ""

	text += fmt.Sprintf("Dice show: %d\n\n", r.results)

	if len(r.matches) > 0 {
		sort.Sort(ByWidthHeight(r.matches))

		text += "Matches:\n"

		for _, m := range r.matches {
			text += fmt.Sprintf("%dx%d\n", m.width, m.height)
		}
	} else {
		text += "No matches\n"
	}

	if r.wiggles > 0 {
		text += fmt.Sprintf("+%d wiggle dice\n", r.wiggles)
	}

	if len(r.loose) > 0 {
		text += fmt.Sprintf("\nLoose dice %d\n", r.loose)
	}

	return text + "\n"
}

// RollDie rolls and sum dice
func RollDie(max, min, numDice int) int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	result := 0
	for i := 1; i < numDice+1; i++ {
		roll := r1.Intn(max+1-min) + min
		result += roll
	}
	return result
}

// Resolve ORE dice roll and prints results
func (r *Roll) Resolve(input string) (*Roll, error) {
	nd, hd, wd, err := parseString(input)

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

	return r, nil

}

// parses string like 5d+1hd+1wd or returns error
func parseString(input string) (int, int, int, error) {

	re := regexp.MustCompile("[0-9]+")

	var dieTypes []string

	errString := ""

	dieTypes = strings.SplitN(input, "+", 3)

	var nd, hd, wd int

	for _, s := range dieTypes {
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

		default:
			errString = "Error: Not a regular die notation"
		}
	}

	nd, hd, wd = verifyLessThan10(nd, hd, wd)

	if errString != "" {
		return 0, 0, 0, errors.New(errString)
	}

	return nd, hd, wd, nil
}

func verifyLessThan10(nd, hd, wd int) (int, int, int) {

	if nd+hd+wd > 10 {

		fmt.Println("Error: Can't roll more than 10 dice. Reducing to less than 10.")
		fmt.Printf(fmt.Sprintf("Current Dice: %dd+%dhd+%dwd.\n", nd, hd, wd))

		// Remove normal dice first
		for nd > 0 && nd+hd+wd > 10 {
			fmt.Printf("reduced Normal dice from %d to %d. \n", nd, nd-1)
			nd--
			fmt.Printf(fmt.Sprintf("Current Dice: %dd+%dhd+%dwd.\n", nd, hd, wd))
		}

		// Reduce hard dice next
		for hd > 0 && nd+hd+wd > 10 {
			fmt.Printf("reduced Hard dice from %d to %d. \n", hd, hd-1)
			hd--
			fmt.Printf(fmt.Sprintf("Current Dice: %dd+%dhd+%dwd.\n", nd, hd, wd))
		}

		// Reduce wiggle dice last
		for wd > 0 && nd+hd+wd > 10 {
			fmt.Printf("reduced Wiggle dice from %d to %d. \n", wd, wd-1)
			wd--
			fmt.Printf(fmt.Sprintf("Current Dice: %dd+%dhd+%dwd.\n", nd, hd, wd))
		}

		return nd, hd, wd

	}

	return nd, hd, wd
}

func (r *Roll) parseDieRoll() *Roll {

	matches := make(map[int]int)
	for _, d := range r.results {
		matches[d]++
	}

	for k, v := range matches {
		switch {
		case v == 1:
			r.loose = append(r.loose, k)
		case v > 1:
			r.matches = append(r.matches, Match{
				height: k,
				width:  v,
			})
		}
	}
	return r
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
