package operations

import "fmt"

var cardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func checkIfAllCardsAreValid(hand string) error {
	for _, c := range hand {
		if _, ok := cardValues[string(c)]; !ok {
			return fmt.Errorf("error: hand %s has an invalid card %c", hand, c)
		}
	}
	return nil
}
func checkNoOfCardsInHand(fh, sh string) error {
	if len(fh) != 5 {
		return fmt.Errorf("first hand should have five cards")
	}
	if len(sh) != 5 {
		return fmt.Errorf("second hand should have five cards")
	}
	return nil
}

func handValidator(fh, sh string) error {
	if err := checkNoOfCardsInHand(fh, sh); err != nil {
		return err
	}
	if err := checkIfAllCardsAreValid(fh); err != nil {
		return err
	}

	if err := checkIfAllCardsAreValid(sh); err != nil {
		return err
	}
	return nil
}
