package operations

import "fmt"

func GetWinner(p1, p2 Player) (string, error) {
	err := handValidator(p1.Hand, p2.Hand)
	if err != nil {
		return "", err
	}
	if err = p1.setHandType(); err != nil {
		return "", fmt.Errorf("problem setting hand type of player 1: %s", err.Error())
	}
	if err = p2.setHandType(); err != nil {
		return "", fmt.Errorf("problem setting hand type of player 2: %s", err.Error())
	}
	return fetchStrongerHand(p1, p2)
}
