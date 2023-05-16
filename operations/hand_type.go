package operations

import "fmt"

const (
	FOUR_KIND  = "fourKind"
	FULL_HOUSE = "fullHourse"
	TRIPLE     = "triple"
	TWO_PAIR   = "twoPair"
	ONE_PAIR   = "onePair"
	HIGH_CARD  = "highCard"
)

type Player struct {
	Hand_Number string
	Hand        string
	Type        string
	CardMap     map[string]int
}

var cardKinds = map[string]int{
	FOUR_KIND:  6,
	FULL_HOUSE: 5,
	TRIPLE:     4,
	TWO_PAIR:   3,
	ONE_PAIR:   2,
	HIGH_CARD:  1,
}

func (p *Player) createCardMap() error {
	if len(p.Hand) != 5 {
		return fmt.Errorf("Player should have five cards in his hand")
	}
	cards := make(map[string]int)
	for _, c := range p.Hand {
		cards[string(c)]++
	}
	p.CardMap = cards
	return nil
}

func (p *Player) setHandType() error {
	err := p.createCardMap()
	if err != nil {
		return err
	}
	if p.hasFourKind() {
		p.Type = FOUR_KIND
	} else if p.isFullHouse() {
		p.Type = FULL_HOUSE
	} else if p.isTriple() {
		p.Type = TRIPLE
	} else {
		pairCount := p.hasPair()
		if pairCount == 1 {
			p.Type = ONE_PAIR
		} else if pairCount == 2 {
			p.Type = TWO_PAIR
		} else {
			p.Type = HIGH_CARD
		}
	}
	return nil
}

func (p *Player) hasFourKind() bool {
	for _, count := range p.CardMap {
		if count == 4 {
			return true
		}
	}
	return false
}

func (p *Player) isFullHouse() bool {
	var hasThree, hasTwo bool
	for _, count := range p.CardMap {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}

	return hasThree && hasTwo
}
func (p *Player) isTriple() bool {
	var hasThree, hasTwo bool
	for _, count := range p.CardMap {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}

	return hasThree && !hasTwo
}
func (p *Player) hasPair() int {
	pairCount := 0
	for _, count := range p.CardMap {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount
}
