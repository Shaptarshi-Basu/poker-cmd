# Poker commad line

Command line program written in Go that evaluates two poker hands and determines the winner based on simplified poker rules.

## Poker Rules

### Cards
- Each player is dealt a hand of 5 cards.
- Available symbols for cards are 23456789TJQKA, in order of value, with A being the most valuable.

### Combinations in Order of Value
1. Four of a Kind: Four cards of the same rank, such as 77377.
2. Full House: A combination of three of a kind and a pair in the same hand, such as KK2K2.
3. Triple: Three cards of the same rank, such as 32666.
4. Two Pairs: Two pairs of cards of the same rank, such as 77332.
5. A Pair: Two cards of the same rank, such as 43K9K.
6. High Card: When none of the above combinations exist, the highest-ranked card in the hand determines the value, such as 297QJ.

### Determining the Winner
- If two hands have the same combination value, the hand with higher-ranked cards wins.
- In the case of two pairs, the higher pair is compared first.
- When comparing full houses, the three-of-a-kind rank is compared first, followed by the pair.
- If the combinations and ranks are identical, the remaining cards are compared from highest to lowest.

## Usage

**Note:** The input hands must be in the format of a string containing the card symbols with no spaces. For example, "77377" represents a hand with four sevens and one three.

## How to RUN

1. Install Docker on your system if it is not already installed.
2. Clone or download the repository to your local machine.
3. Navigate to the root directory of the project where the Dockerfile is located.
4. Build the Docker image by running the following command:
