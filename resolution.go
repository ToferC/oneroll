package oneroll

import (
	"fmt"
	"sort"
)

// FormDieString takes a stat and skill and creates a die pool string
func FormSkillDieString(skill *Skill, actions int) string {

	normal := skill.LinkStat.Dice.Normal + skill.Dice.Normal
	hard := skill.LinkStat.Dice.Hard + skill.Dice.Hard
	wiggle := skill.LinkStat.Dice.Wiggle + skill.Dice.Wiggle
	goFirst := Max(skill.LinkStat.Dice.GoFirst, skill.Dice.GoFirst)
	spray := Max(skill.LinkStat.Dice.Spray, skill.Dice.Spray)

	text := fmt.Sprintf("%dac+%dd+%dhd+%dwd+%dgf+%dsp",
		actions,
		normal,
		hard,
		wiggle,
		goFirst,
		spray)

	return text
}

// OpposedRoll determines the results of an opposed roll between two or more actors
func OpposedRoll(rolls ...*Roll) []Match {

	fmt.Println("Opposed Roll Resolution")

	var results []Match
	wd := make(map[string]int)

	for _, r := range rolls {

		fmt.Printf("Actor: %s, Action: %s, GoFirst: %d, Spray: %d, Wiggle Dice: %dwd\n",
			r.Actor.Name,
			r.Action,
			r.DiePool.GoFirst,
			r.DiePool.Spray,
			r.wiggles,
		)

		for _, m := range r.matches {
			results = append(results, m)
		}
		sort.Sort(ByWidthHeight(results))
	}
	return results
}

func PrintOpposed(results []Match) {
	fmt.Println("***Resolution***")

	for i, m := range results {
		fmt.Printf("***ACTION %d: Actor: %s, Match: %dx%d, Initiative: %dx%d\n",
			i+1,
			m.Actor.Name,
			m.height, m.width,
			m.height, m.initiative,
		)
	}
}
