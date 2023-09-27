package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)
	// fmt.Println(countries)

	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"
	// fmt.Println(countries)
	// fmt.Println(countries["Argentina"])

	championship := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(championship)

	// for team, score := range championship {
	// 	fmt.Printf("Team %s, has a score of %d \n", team, score)
	// }

	delete(championship, "Real Madrid")
	fmt.Println(championship)

	score, exists := championship["Chivas"]
	fmt.Printf("Capured score is %d and the equipment exists = %t \n", score, exists)
}
