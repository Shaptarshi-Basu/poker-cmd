package operations

import (
	"fmt"
	"sort"
)

func fetchStrongerHand(p1, p2 Player) (string, error) {
	p1handValue, ok := cardKinds[p1.Type]
	if !ok {
		return "", fmt.Errorf("invalid hand type %s for first player", p1.Type)
	}
	p2handValue, ok := cardKinds[p2.Type]
	if !ok {
		return "", fmt.Errorf("invalid hand type %s for second player ", p1.Type)
	}
	if p1handValue == p2handValue {
		switch p1.Type {
		case FOUR_KIND:
			return strongerFourKind(p1, p2)
		case FULL_HOUSE:
			return strongerFullHouse(p1, p2)
		case TRIPLE:
			return strongerTriple(p1, p2)
		case TWO_PAIR:
			return strongerPair(p1, p2)
		case ONE_PAIR:
			return strongerPair(p1, p2)
		case HIGH_CARD:
			return strongerHighCard(p1, p2)
		}
	} else if p1handValue > p2handValue {
		return p1.Hand_Number, nil
	} else if p2handValue > p1handValue {
		return p2.Hand_Number, nil
	}
	return "", fmt.Errorf("invalid hands")
}

func strongerFourKind(p1, p2 Player) (string, error) {
	var value1, value2 int
	var ok bool
	for k, v := range p1.CardMap {
		if v == 4 {
			value1, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	for k, v := range p2.CardMap {
		if v == 4 {
			value2, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	if value1 > value2 {
		return p1.Hand_Number, nil
	} else if value2 > value1 {
		return p2.Hand_Number, nil
	} else {
		hands1, err := returnHandPointsForSingles(p1.CardMap)
		if err != nil {
			return "", err
		}

		hands2, err := returnHandPointsForSingles(p2.CardMap)
		if err != nil {
			return "", err
		}

		res := compareHands(hands1, hands2)
		return res, nil
	}
}

func strongerTriple(p1, p2 Player) (string, error) {
	var value1, value2 int
	var ok bool
	for k, v := range p1.CardMap {
		if v == 3 {
			value1, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	for k, v := range p2.CardMap {
		if v == 3 {
			value2, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	if value1 > value2 {
		return p1.Hand_Number, nil
	} else if value2 > value1 {
		return p2.Hand_Number, nil
	} else {
		hands1, err := returnHandPointsForSingles(p1.CardMap)
		if err != nil {
			return "", err
		}
		hands2, err := returnHandPointsForSingles(p2.CardMap)
		if err != nil {
			return "", err
		}
		result := compareHands(hands1, hands2)
		return result, nil
	}
}

func strongerFullHouse(p1, p2 Player) (string, error) {
	var value1, value2 int
	var ok bool
	for k, v := range p1.CardMap {
		if v == 3 {
			value1, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	for k, v := range p2.CardMap {
		if v == 3 {
			value2, ok = cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
		}
	}
	if value1 > value2 {
		return p1.Hand_Number, nil
	} else if value2 > value1 {
		return p2.Hand_Number, nil
	} else {
		result, err := strongerPair(p1, p2)
		if err != nil {
			return "", err
		}
		if result != "Tie" {
			return result, nil
		}
		hands1, err := returnHandPointsForSingles(p1.CardMap)
		if err != nil {
			return "", err
		}
		hands2, err := returnHandPointsForSingles(p2.CardMap)
		if err != nil {
			return "", err
		}
		result = compareHands(hands1, hands2)
		return result, nil
	}
}

func strongerPair(p1, p2 Player) (string, error) {
	var values1, values2 []int
	for k, v := range p1.CardMap {
		if v == 2 {
			value1, ok := cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
			values1 = append(values1, value1)
		}
	}
	for k, v := range p2.CardMap {
		if v == 2 {
			value2, ok := cardValues[k]
			if !ok {
				return "", fmt.Errorf("value for card %s could not be found", k)
			}
			values2 = append(values2, value2)
		}
	}
	result := compareHands(values1, values2)
	if result != "Tie" {
		return result, nil
	} else {
		hands1, err := returnHandPointsForSingles(p1.CardMap)
		if err != nil {
			return "", err
		}
		hands2, err := returnHandPointsForSingles(p2.CardMap)
		if err != nil {
			return "", err
		}
		result = compareHands(hands1, hands2)
		return result, nil
	}
}

func strongerHighCard(p1, p2 Player) (string, error) {
	hands1, err := returnHandPointsForSingles(p1.CardMap)
	if err != nil {
		return "", err
	}
	hands2, err := returnHandPointsForSingles(p2.CardMap)
	if err != nil {
		return "", err
	}
	result := compareHands(hands1, hands2)
	return result, nil
}
func returnHandPointsForSingles(cardMap map[string]int) ([]int, error) {
	var values []int
	for k, v := range cardMap {
		if v == 1 {
			value, ok := cardValues[k]
			if !ok {
				return []int{}, fmt.Errorf("value for the card %s could not be found", k)
			}
			values = append(values, value)
		}
	}
	return values, nil
}

func compareHands(hands1, hands2 []int) string {
	// Sort array1 in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(hands1)))

	// Sort array2 in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(hands2)))

	for i := 0; i < len(hands1); i++ {
		if hands1[i] > hands2[i] {
			return "Hand 1"
		} else if hands2[i] > hands1[i] {
			return "Hand 2"
		}
	}
	return "Tie"
}
