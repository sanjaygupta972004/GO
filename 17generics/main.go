package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{
		Suit: suit,
		Rank: rank,
	}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

type Deck[C any] struct {
	cards []C
}

func (d *Deck[C]) AddCard(card C) {
	d.cards = append(d.cards, card)

}

func (d *Deck[C]) RandomCard() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cardIndex := r.Intn(len(d.cards))
	return d.cards[cardIndex]
}

func NewPlayingDeckCard() *Deck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "6", "7", "8", "10", "j", "q", "k"}
	deck := &Deck[*PlayingCard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func main() {
	fmt.Println("Learning generics in golang")

	deck := NewPlayingDeckCard()
	fmt.Printf("--drawing playing card ---\n")
	PlayingCard := deck.RandomCard()
	fmt.Printf("draw card %s\n", PlayingCard)

	fmt.Printf("card suits : %s\n", PlayingCard.Suit)
	fmt.Printf("card rank : %s\n", PlayingCard.Rank)

}
