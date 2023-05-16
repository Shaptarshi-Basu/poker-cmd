package operations

import (
	"fmt"
	"testing"
)

func Test_strongerFourKind(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name    string
		player1 Player
		player2 Player
		result  string
	}{
		// the table itself
		{"Both hands have different fours but hand 1 has higher four", Player{Hand_Number: "Hand 1", Hand: "8888A"}, Player{Hand_Number: "Hand 2", Hand: "7K777"}, "Hand 1"},
		{"Both hands have different fours but hand 2 has higher four", Player{Hand_Number: "Hand 1", Hand: "8888A"}, Player{Hand_Number: "Hand 2", Hand: "7AAAA"}, "Hand 2"},
		{"Both hands have same fours and same other hand", Player{Hand_Number: "Hand 1", Hand: "88A88"}, Player{Hand_Number: "Hand 2", Hand: "8888A"}, "Tie"},
		{"Both hands have same fours but hand 2 has higher other hand", Player{Hand_Number: "Hand 1", Hand: "88K88"}, Player{Hand_Number: "Hand 2", Hand: "8888A"}, "Hand 2"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player1.createCardMap()
			tt.player2.createCardMap()

			tt.player1.setHandType()
			tt.player2.setHandType()

			fmt.Print("Player type")

			fmt.Print(tt.player1.Type)
			fmt.Print(tt.player2.Type)

			result, err := strongerFourKind(tt.player1, tt.player2)
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if result != tt.result {
				t.Errorf("Result should have been %s but  got %s.", tt.result, result)
			}

		})
	}
}

func Test_strongerTriple(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name    string
		player1 Player
		player2 Player
		result  string
	}{
		// the table itself
		{"Both hands have different three but hand 1 has higher three", Player{Hand_Number: "Hand 1", Hand: "888AK"}, Player{Hand_Number: "Hand 2", Hand: "AA777"}, "Hand 1"},
		{"Both hands have different three but hand 2 has higher three", Player{Hand_Number: "Hand 1", Hand: "888AK"}, Player{Hand_Number: "Hand 2", Hand: "72AAA"}, "Hand 2"},
		{"Both hands have same three and same other hand", Player{Hand_Number: "Hand 1", Hand: "8AK88"}, Player{Hand_Number: "Hand 2", Hand: "8K88A"}, "Tie"},
		{"Both hands have same three but hand 2 has higher other hand", Player{Hand_Number: "Hand 1", Hand: "88A28"}, Player{Hand_Number: "Hand 2", Hand: "88K8A"}, "Hand 2"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player1.createCardMap()
			tt.player2.createCardMap()

			tt.player1.setHandType()
			tt.player2.setHandType()

			fmt.Print("Player type")

			fmt.Print(tt.player1.Type)
			fmt.Print(tt.player2.Type)
			res, err := strongerTriple(tt.player1, tt.player2)
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if res != tt.result {
				fmt.Print(tt.player1.Type)
				fmt.Print(tt.player2.Type)
				t.Errorf("Result should have been %+v but  got %s", tt.result, res)
			}
		})
	}
}

func Test_strongerPair(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name    string
		player1 Player
		player2 Player
		result  string
	}{
		{"Both hands have different pairs but hand 1 has higher pair", Player{Hand_Number: "Hand 1", Hand: "88776"}, Player{Hand_Number: "Hand 2", Hand: "6677A"}, "Hand 1"},
		{"Both hands have different pair but hand 2 has higher three", Player{Hand_Number: "Hand 1", Hand: "886AK"}, Player{Hand_Number: "Hand 2", Hand: "72KAA"}, "Hand 2"},
		{"Both hands have same pairs and same other hand", Player{Hand_Number: "Hand 1", Hand: "88AAK"}, Player{Hand_Number: "Hand 2", Hand: "8A8AK"}, "Tie"},
		{"Both hands have same pairs but hand 2 has higher other hand", Player{Hand_Number: "Hand 1", Hand: "88AA5"}, Player{Hand_Number: "Hand 2", Hand: "88AA7"}, "Hand 2"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player1.createCardMap()
			tt.player2.createCardMap()

			tt.player1.setHandType()
			tt.player2.setHandType()

			res, err := strongerPair(tt.player1, tt.player2)
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if res != tt.result {
				fmt.Print(tt.player1.Type)
				fmt.Print(tt.player2.Type)
				t.Errorf("Result should have been %+v but  got %s", tt.result, res)
			}
		})
	}
}

func Test_strongerFullHouse(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name    string
		player1 Player
		player2 Player
		result  string
	}{
		{"Both hands have same Full House ", Player{Hand_Number: "Hand 1", Hand: "88778"}, Player{Hand_Number: "Hand 2", Hand: "78788"}, "Tie"},
		{"First player has bigger Full house", Player{Hand_Number: "Hand 1", Hand: "K88KK"}, Player{Hand_Number: "Hand 2", Hand: "88877"}, "Hand 1"},
		{"Second player has bigger Full house", Player{Hand_Number: "Hand 1", Hand: "A7AA7"}, Player{Hand_Number: "Hand 2", Hand: "AKAAK"}, "Hand 2"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player1.createCardMap()
			tt.player2.createCardMap()

			tt.player1.setHandType()
			tt.player2.setHandType()

			res, err := strongerFullhouse(tt.player1, tt.player2)
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if res != tt.result {
				fmt.Print(tt.player1.Type)
				fmt.Print(tt.player2.Type)
				t.Errorf("Result should have been %+v but  got %s", tt.result, res)
			}
		})
	}
}
func Test_strongerHighCard(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name    string
		player1 Player
		player2 Player
		result  string
	}{
		{"Both hands have same Full House ", Player{Hand_Number: "Hand 1", Hand: "867KA"}, Player{Hand_Number: "Hand 2", Hand: "KJQ87"}, "Hand 1"},
		{"Second player has bigger Full house", Player{Hand_Number: "Hand 1", Hand: "KJQA7"}, Player{Hand_Number: "Hand 2", Hand: "K7AJQ"}, "Tie"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.player1.createCardMap()
			tt.player2.createCardMap()

			tt.player1.setHandType()
			tt.player2.setHandType()

			res, err := strongerHighCard(tt.player1, tt.player2)
			if err != nil {
				t.Errorf("Error should have been nil but got %s", err.Error())
			}
			if res != tt.result {
				fmt.Print(tt.player1.Type)
				fmt.Print(tt.player2.Type)
				t.Errorf("Result should have been %+v but  got %s", tt.result, res)
			}
		})
	}
}
