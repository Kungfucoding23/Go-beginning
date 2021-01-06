package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//teamPlay --
type teamPlay struct {
	name string
	MP   int //MP: Matches Played
	W    int //W: Matches Won
	D    int //D: Matches Drawn (Tied)
	L    int //L: Matches Lost
	P    int //P: Points
}

var points map[string]int = map[string]int{
	"win":  3,
	"draw": 1,
	"loss": 0,
}

func main() {

	//Resultados de los partidos
	match1 := "Allegoric Alaskans;Blithering Badgers;win" // Esta linea por ejemplo significa que los Allegoric Alaskans ganaron contra los lithering Badgers
	match2 := "Devastating Donkeys;Courageous Californians;draw"
	match3 := "Devastating Donkeys;Allegoric Alaskans;win"
	match4 := "Courageous Californians;Blithering Badgers;loss"
	match5 := "Blithering Badgers;Devastating Donkeys;loss"
	match6 := "Allegoric Alaskans;Courageous Californians;win"

	//Ahora guardo estos resultados en un slice
	resultList := []string{match1, match2, match3, match4, match5, match6}
	fmt.Println(resultList)

	//Creo un map donde se guardaran los resultados despues de haber quitado los ; con la funcion cleanMatchResults
	mapTeams := make(map[string]teamPlay)

	//Limpia los resultados y los guarda en un slice
	for i := 0; i < len(resultList); i++ {
		cleanResultList := CleanMatchResults(resultList[i])
		//en team1 va el primer equipo, por ejemplo de la primer linea serian los Allegoric Alaskans
		team1 := mapTeams[cleanResultList[0]]
		team2 := mapTeams[cleanResultList[1]]

		team1.name = cleanResultList[0]
		team2.name = cleanResultList[1]

		//cargo el primer partido de estos primeros equipos
		team1.MP++
		team2.MP++
		//A win earns a team 3 points. A draw earns 1. A loss earns 0.
		switch cleanResultList[2] {
		case "win":
			team1.W++
			team1.P += 3
			team2.L++
		case "loss":
			team1.L++
			team2.W++
			team2.P += 3
		case "draw":
			team1.D++
			team1.P++
			team2.D++
			team2.P++
		default:
			fmt.Println("Error")
		}
		mapTeams[cleanResultList[0]] = team1
		mapTeams[cleanResultList[1]] = team2
	}

	//Guardo los resultados limpios en el slice que luego voy a ordenar
	resultSorted := make([]teamPlay, 0, len(mapTeams))
	for _, team := range mapTeams {
		resultSorted = append(resultSorted, team)
	}

	//The outcome should be ordered by points, descending. In case of a tie, teams are ordered alphabetically.
	sort.Slice(resultSorted, func(i, j int) bool {
		if resultSorted[i].P == resultSorted[j].P {
			//si es empate los ordena alfabeticamente
			return resultSorted[i].name < resultSorted[j].name
		}
		//sino los ordena por puntaje descendiente
		return resultSorted[i].P > resultSorted[j].P
	})

	const printResults = "%-31v |  %2v |  %2v |  %2v |  %2v |  %2v\n"

	fmt.Fprintf(os.Stdout, printResults, "Team", "MP", "W", "D", "L", "P")

	//Imprime el resultado
	for _, team := range resultSorted {
		fmt.Fprintf(os.Stdout, printResults, team.name, team.MP, team.W, team.D, team.L, team.P)
	}

}

//CleanMatchResults limpia los resultados sacando los ; con strings.Split
func CleanMatchResults(result string) []string {
	cleanResult := strings.Split(result, ";")
	return cleanResult
}
