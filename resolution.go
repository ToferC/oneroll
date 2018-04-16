package oneroll

import (
	"fmt"
	"sort"
)

// FormDieString takes a stat and skill and creates a die pool string
func FormSkillDieString(skill *Skill) string {

	normal := skill.LinkStat.Dice.Normal + skill.Dice.Normal
	hard := skill.LinkStat.Dice.Hard + skill.Dice.Hard
	wiggle := skill.LinkStat.Dice.Wiggle + skill.Dice.Wiggle
	goFirst := Max(skill.LinkStat.Dice.GoFirst, skill.Dice.GoFirst)
	spray := Max(skill.LinkStat.Dice.Spray, skill.Dice.Spray)

	text := fmt.Sprintf("%dd+%dhd+%dwd+%dgf+%dsp", normal, hard, wiggle, goFirst, spray)

	return text
}

// OpposedRoll determines the results of an opposed roll between two or more actors
func OpposedRoll(rolls ...*Roll) {

	fmt.Println("Opposed Roll Resolution")

	var results []Match
	wd := make(map[string]int)

	for _, r := range rolls {
		wd[r.Actor.Name] = r.wiggles

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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
