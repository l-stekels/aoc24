package day5

import (
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{
		Rules:  map[int][]int{},
		Update: []Update{},
	}
	stringRows := strings.Split(content, "\n")
	for _, row := range stringRows {
		if strings.Contains(row, "|") {
			rules := strings.Split(row, "|")
			before, err := strconv.Atoi(rules[0])
			if err != nil {
				return result, err
			}
			after, err := strconv.Atoi(rules[1])
			if err != nil {
				return result, err
			}
			_, exists := result.Rules[before]
			if !exists {
				result.Rules[before] = []int{}
			}
			result.Rules[before] = append(result.Rules[before], after)
			_, exists = result.Rules[after]
			if !exists {
				result.Rules[after] = []int{}
			}
		}
		if strings.Contains(row, ",") {
			pages := strings.Split(row, ",")
			update := Update{
				Pages:      []int{},
				MiddlePage: 0,
			}
			for _, page := range pages {
				page, err := strconv.Atoi(page)
				if err != nil {
					return result, err
				}
				update = update.AppendPage(page)
			}
			result.Update = append(result.Update, update)
		}
	}

	return result, nil
}

type SolutionInput struct {
	Update []Update
	Rules  map[int][]int
}

type Update struct {
	Pages      []int
	MiddlePage int
}

func (u Update) Equals(other Update) bool {
	if u.MiddlePage != other.MiddlePage {
		return false
	}
	if len(u.Pages) != len(other.Pages) {
		return false
	}
	for i, page := range u.Pages {
		if page != other.Pages[i] {
			return false
		}
	}

	return true
}

func (u Update) CreatePositionsMap() map[int]int {
	result := map[int]int{}
	for i, page := range u.Pages {
		result[page] = i
	}

	return result
}

func (u Update) IsUpdateValid(rules map[int][]int) bool {
	positions := u.CreatePositionsMap()
	for _, page := range u.Pages {
		_, exists := rules[page]
		if !exists {
			continue
		}
		mustComeAfter := rules[page]
		currentPagePosition := positions[page]
		for _, after := range mustComeAfter {
			position, pageExists := positions[after]
			if !pageExists {
				continue
			}
			if currentPagePosition >= position {
				return false
			}
		}
	}

	return true
}

func (u Update) AppendPage(page int) Update {
	result := Update{
		Pages:      append(u.Pages, page),
		MiddlePage: u.MiddlePage,
	}
	result.MiddlePage = result.Pages[len(u.Pages)/2]

	return result
}

func (u Update) ReorderUpdate(rules map[int][]int) Update {
	result := Update{
		Pages:      []int{},
		MiddlePage: 0,
	}
	// Rules that apply to the pages in the current update
	currentRules := map[int][]int{}
	inDegree := map[int]int{}
	for _, page := range u.Pages {
		currentRules[page] = []int{}
		inDegree[page] = 0
	}

	for _, page := range u.Pages {
		_, exists := rules[page]
		if !exists {
			continue
		}
		for _, after := range rules[page] {
			_, existsInSubGraph := currentRules[after]
			if !existsInSubGraph {
				continue
			}
			currentRules[page] = append(currentRules[page], after)
			inDegree[after]++
		}
	}
	var queue []int
	for _, page := range u.Pages {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}
	for len(queue) > 0 {
		// Pop the first element off the queue
		currentPage := queue[0]
		queue = queue[1:]

		result = result.AppendPage(currentPage)
		for _, after := range currentRules[currentPage] {
			inDegree[after]--
			if inDegree[after] == 0 {
				queue = append(queue, after)
			}
		}
	}

	return result
}

func (s SolutionInput) Validate() error {
	return nil
}

func (s SolutionInput) Equals(other SolutionInput) bool {
	for i, update := range s.Update {
		if !update.Equals(other.Update[i]) {
			return false
		}
	}
	if len(s.Rules) != len(other.Rules) {
		return false
	}
	for i, rule := range s.Rules {
		if len(rule) != len(other.Rules[i]) {
			return false
		}
		for j, r := range rule {
			if r != other.Rules[i][j] {
				return false
			}
		}
	}

	return true
}

func SolvePart1(input SolutionInput) int {
	result := 0

	for _, update := range input.Update {
		if update.IsUpdateValid(input.Rules) {
			result += update.MiddlePage
		}
	}

	return result
}

func SolvePart2(input SolutionInput) int {
	result := 0

	for _, update := range input.Update {
		if update.IsUpdateValid(input.Rules) {
			continue
		}
		reorderedUpdate := update.ReorderUpdate(input.Rules)
		if !reorderedUpdate.IsUpdateValid(input.Rules) {
			panic("There should be no invalid update entries at this point!")
		}
		result += reorderedUpdate.MiddlePage
	}

	return result
}
