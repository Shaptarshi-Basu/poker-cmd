package operations

import (
	"testing"
)

func Test_createCardMap(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result map[string]int
	}{
		// the table itself
		{"Create card map for FOUR_KIND ", Player{Hand_Number: "Hand 1", Hand: "8888A"}, map[string]int{"8": 4, "A": 1}},
		{"Create card map for TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, map[string]int{"8": 2, "A": 2, "K": 1}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()

			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if !mapEqual(tt.result, tt.player.CardMap) {
				t.Errorf("Result should have been %+v but  got %+v.", tt.result, tt.player.CardMap)
			}

		})
	}
}

func Test_hasFourKind(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result bool
	}{
		// the table itself
		{"Player has hand with FOUR_KIND ", Player{Hand_Number: "Hand 1", Hand: "8888A"}, true},
		{"Player has hand with TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, false},
		{"Player has hand with Triple ", Player{Hand_Number: "Hand 1", Hand: "88A8K"}, false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()

			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			result := tt.player.hasFourKind()
			if tt.result != result {
				t.Errorf("Result should have been %t but  got %t.", tt.result, result)
			}

		})
	}
}

func Test_isTriple(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result bool
	}{
		// the table itself
		{"Player has hand with FOUR_KIND ", Player{Hand_Number: "Hand 1", Hand: "8888A"}, false},
		{"Player has hand with TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, false},
		{"Player has hand with TRIPLE ", Player{Hand_Number: "Hand 1", Hand: "88A8K"}, true},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()

			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			result := tt.player.isTriple()
			if tt.result != result {
				t.Errorf("Result should have been %t but  got %t.", tt.result, result)
			}

		})
	}
}

func Test_isFullHouse(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result bool
	}{
		// the table itself
		{"Player has hand with FULL_HOUSE ", Player{Hand_Number: "Hand 1", Hand: "88A8A"}, true},
		{"Player has hand with TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, false},
		{"Player has hand with TRIPLE ", Player{Hand_Number: "Hand 1", Hand: "88A8K"}, false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()

			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			result := tt.player.isFullHouse()
			if tt.result != result {
				t.Errorf("Result should have been %t but  got %t.", tt.result, result)
			}

		})
	}
}

func Test_hasPair(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result int
	}{
		// the table itself
		{"Player has hand with FULL_HOUSE ", Player{Hand_Number: "Hand 1", Hand: "8888A"}, 0},
		{"Player has hand with TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, 2},
		{"Player has hand with ONE_PAIR ", Player{Hand_Number: "Hand 1", Hand: "87A8K"}, 1},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()

			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			result := tt.player.hasPair()
			if tt.result != result {
				t.Errorf("Result should have been %d but  got %d.", tt.result, result)
			}

		})
	}
}

func Test_setHandType(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		player Player
		result string
	}{
		// the table itself
		{"Player has hand with FOUR_KIND ", Player{Hand_Number: "Hand 1", Hand: "8888A"}, FOUR_KIND},
		{"Player has hand with TWO_PAIR ", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, TWO_PAIR},
		{"Player has hand with ONE_PAIR ", Player{Hand_Number: "Hand 1", Hand: "87A8K"}, ONE_PAIR},
		{"Player has hand with TRIPLE ", Player{Hand_Number: "Hand 1", Hand: "88A8K"}, TRIPLE},
		{"Player has hand with HIGH_CARD ", Player{Hand_Number: "Hand 1", Hand: "87AJK"}, HIGH_CARD},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.player.createCardMap()
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}

			err = tt.player.setHandType()
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			result := tt.player.Type
			if tt.result != result {
				t.Errorf("Result should have been %s but  got %s.", tt.result, result)
			}
		})
	}
}

func mapEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}
