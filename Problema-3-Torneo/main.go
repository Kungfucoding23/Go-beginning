package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Define a function Tally(io.Reader, io.Writer) error.
//
// Note that unlike other tracks the Go version of the tally function
// should not ignore errors. It's not idiomatic Go to ignore errors.

var _ func(io.Reader, io.Writer) error = Tally

var happyTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "good",
		input: `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`[1:], // [1:] = strip initial readability newline
	},
	{
		description: "ignore comments and newlines",
		input: `

Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;win
# Catastrophic Loss of the Californians
Courageous Californians;Blithering Badgers;loss

Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
Devastating Donkeys;Courageous Californians;draw


`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`[1:],
	},
	{
		// A complete competition has all teams play eachother once or twice.
		description: "incomplete competition",
		input: `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Allegoric Alaskians;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  2 |  1 |  0 |  1 |  3
Devastating Donkeys            |  1 |  1 |  0 |  0 |  3
Courageous Californians        |  2 |  0 |  0 |  2 |  0
`[1:],
	},
	{
		description: "tie for first and last place",
		input: `
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskians;Courageous Californians;draw
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskians            |  3 |  2 |  1 |  0 |  7
Courageous Californians        |  3 |  2 |  1 |  0 |  7
Blithering Badgers             |  3 |  0 |  1 |  2 |  1
Devastating Donkeys            |  3 |  0 |  1 |  2 |  1
`[1:],
	},
}

// record holds data of one team

type record struct {
	team  string
	stats map[string]int
}

func newRecord(team string) *record {
	return &record{
		team: team,
		stats: map[string]int{
			"MP": 0,
			"W":  0,
			"D":  0,
			"L":  0,
			"P":  0,
		}}
}

func (r record) AddScore(score string) {
	r.stats["MP"]++
	r.stats["P"] += r.countPoints(score)
	r.stats[score]++
}

func (r record) GetPoints() int {
	return r.stats["P"]
}

func (r record) String() string {
	return fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", r.team, r.stats["MP"], r.stats["W"], r.stats["D"], r.stats["L"], r.stats["P"])
}

func (r record) countPoints(score string) int {
	switch score {
	case "W":
		return 3
	case "D":
		return 1
	default:
		return 0
	}
}

// byPoints is sorting type for records
type byPoints []record

func (a byPoints) Len() int { return len(a) }
func (a byPoints) Less(i, j int) bool {
	if a[i].GetPoints() == a[j].GetPoints() {
		return a[i].team[0:1] < a[j].team[0:1]
	}
	return a[i].GetPoints() > a[j].GetPoints()
}
func (a byPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// tallyMap holds multiple records
type tallyMap map[string]record

func (m *tallyMap) Insert(team string, score string) {

	if item, ok := (*m)[team]; ok {
		item.AddScore(score)
	} else {
		(*m)[team] = *newRecord(team)
		m.Insert(team, score)
	}
}

func (m tallyMap) String() (stringified string) {
	stringified = "Team                           | MP |  W |  D |  L |  P\n"
	sorted := m.getSortedElements()
	for _, record := range sorted {
		stringified += fmt.Sprint(record)
	}
	return
}

func (m tallyMap) getSortedElements() byPoints {
	sorted := make([]record, 0)
	for _, record := range m {
		sorted = append(sorted, record)
	}
	sort.Sort(byPoints(sorted))
	return sorted
}

// Tally parse input to table string
func Tally(r io.Reader, w io.Writer) error {
	text, err := readString(r)
	if err != nil {
		return err
	}

	parsed, err := parseToMap(text)
	if err != nil {
		return err
	}

	w.Write([]byte(fmt.Sprint(parsed)))
	return nil
}

func readString(r io.Reader) (string, error) {
	str := ""
	b := make([]byte, 1)
	for {
		if _, err := r.Read(b); err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		} else {
			str += string(b)
		}

	}
	return str, nil
}

func parseToMap(results string) (tallyMap, error) {
	result := tallyMap{}
	lines := strings.Split(results, "\n")
	for _, line := range lines {
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			parts := strings.Split(line, ";")
			err := validateInput(parts)
			if err != nil {
				return nil, err
			}

			st, nd := interpretMatchResult(parts[2])
			result.Insert(parts[0], st)
			result.Insert(parts[1], nd)
		}
	}

	return result, nil
}

func validateInput(parts []string) error {
	if len(parts) < 3 {
		return fmt.Errorf("Invalid record")
	}

	if parts[0] == parts[1] {
		return fmt.Errorf("Same team")
	}

	if parts[2] != "win" && parts[2] != "loss" && parts[2] != "draw" {
		return fmt.Errorf("Invalid match result")
	}

	return nil
}

func interpretMatchResult(result string) (string, string) {
	switch result {
	case "win":
		return "W", "L"
	case "loss":
		return "L", "W"
	}
	return "D", "D"
}

func main() {

}
