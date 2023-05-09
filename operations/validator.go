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

func checkifAllCardsAreValid(hand string) error {
	for _, c := range hand {
		if _, ok := cardValues[string(c)]; !ok {
			return fmt.Errorf("Error: Hand %s has an invalid card %c\n", hand, c)
		}
	}
	return nil
}
func checkNoOfCardsInHand(fh, sh string) {
	if len(fh) != 5 {
		fmt.Print("First hand should have five cards")
		return
	}
	if len(sh) != 5 {
		fmt.Println("Second hand should have five cards")
		return
	}
}

func handValidator(fh, sh string) {
	checkNoOfCardsInHand(fh, sh)

	if err := checkifAllCardsAreValid(fh); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := checkifAllCardsAreValid(sh); err != nil {
		fmt.Println(err.Error())
		return
	}
}
