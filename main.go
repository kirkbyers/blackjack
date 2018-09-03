package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kirkbyers/deck"
)

// Hand represents a player's hand
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString prints the hand, showing only the first card and hidding the rest
func (h Hand) DealerString() string {
	return h[0].String() + ", *HIDDEN*"
}

// Score is the hand's score that acocunts for aces
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// MinScore is the smallest score of a hand
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var (
		player, dealer Hand
		card           deck.Card
	)
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Dealer shows:", dealer.DealerString())
		fmt.Println("Player hand:", player)
		fmt.Println("(h)it or (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		case "s":
		case "":
		case "\n":
		default:
			fmt.Println("***INVALID INPUT***")
			time.Sleep(1 * time.Second)
		}
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("===FINAL HANDS===")
	fmt.Println("Dealer hand:", dealer, "\nDealer score:", dScore)
	fmt.Println("Player hand:", player, "\nPlayer score:", pScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case pScore < dScore:
		fmt.Println("You lose!")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
