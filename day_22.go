package adventofcode2020

import (
	"fmt"
	"strconv"
	"strings"
)

type day22player struct {
	name  string
	head  *day22card
	tail  *day22card
	cards int
}

func (day22 *day22player) copy(numcards int) *day22player {

	next := &day22player{day22.name, nil, nil, 0}
	card := day22.head
	for card != nil && numcards > 0 {
		next.push(card.num)
		card = card.next
		numcards--
	}
	return next
}

func (day22 *day22player) pop() int {
	card := day22.head.num
	day22.head = day22.head.next
	day22.cards--
	return card
}

func (day22 *day22player) push(value int) {

	if day22.head == nil {
		// first card
		day22.head = &day22card{value, nil}
		day22.tail = day22.head
	} else {
		day22.tail.next = &day22card{value, nil}
		day22.tail = day22.tail.next
	}
	day22.cards++
}

func (day22 *day22player) countCards() int {
	count := 0
	card := day22.head
	for card != nil {
		count++
		card = card.next
	}
	return count
}

func (day22 *day22player) String() string {

	card := day22.head
	if card == nil {
		return fmt.Sprintf("%s: nil", day22.name)
	}
	result := fmt.Sprintf("%s: %d", day22.name, card.num)
	card = card.next
	for card != nil {
		result = fmt.Sprintf("%s, %d", result, card.num)
		card = card.next
	}
	return result
}

func (day22 *day22player) sum() int {
	fmt.Println(day22.String())
	count := day22.cards
	if count == 0 {
		return 0
	}
	card := day22.head
	sum := 0
	for card != nil {
		sum += (card.num * count)
		count--
		card = card.next
	}
	return sum
}

type day22card struct {
	num  int
	next *day22card
}

func day22Load(input []string) (*day22player, *day22player, error) {

	player1 := &day22player{"Player 1", nil, nil, 0}
	player2 := &day22player{"Player 2", nil, nil, 0}

	p2 := false

	// we know it always starts with player 1
	for _, line := range input[1:] {

		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Player 2") {
			p2 = true
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, err
		}
		if p2 {
			player2.push(value)
		} else {
			player1.push(value)
		}
	}
	return player1, player2, nil
}

func day22(input []string) (int, error) {

	player1, player2, err := day22Load(input)
	if err != nil {
		return 0, err
	}

	rounds := 0

	// play!
	for rounds < 1024 {

		// fmt.Println(player1.String())
		// fmt.Println(player2.String())

		// both players take a card
		p1card := player1.pop()
		p2card := player2.pop()

		if p1card > p2card {
			// p1 wins
			//fmt.Printf("Player 1 wins Round %d\n", rounds)
			player1.push(p1card)
			player1.push(p2card)
		} else {
			//fmt.Printf("Player 2 wins Round %d\n", rounds)
			player2.push(p2card)
			player2.push(p1card)
		}
		if player1.head == nil {
			// player 2 wins !
			fmt.Println("Player 2 wins")
			return player2.sum(), nil
		} else if player2.head == nil {
			// player 1 wins
			fmt.Println("Player 1 wins")
			return player1.sum(), nil
		}
		rounds++
	}

	return 0, fmt.Errorf("too many loops %d", rounds)
}

func day22Part2(input []string) (int, error) {

	player1, player2, err := day22Load(input)
	if err != nil {
		return 0, err
	}

	game := 1
	winner, err := day22Part2Play(player1, player2, game)
	if err != nil {
		return 0, err
	}
	return winner.sum(), nil
}

func day22Part2Play(player1 *day22player, player2 *day22player, game int) (*day22player, error) {

	rounds := 1

	seen := make(map[string]bool)

	// play!
	for rounds < 250000 {

		// prevent infinite recursion
		p1s := player1.String()
		p2s := player2.String()

		if _, prs := seen[p1s]; prs {
			// player 1 win by defualt
			return player1, nil
		}
		if _, prs := seen[p2s]; prs {
			// player 1 win by defualt
			return player1, nil
		}
		seen[p1s] = true
		seen[p2s] = true

		// fmt.Println(player1.String())
		// fmt.Println(player2.String())

		// both players take a card
		p1card := player1.pop()
		p2card := player2.pop()

		if p1card <= player1.cards && p2card <= player2.cards {
			// then we must recurse

			p1r := player1.copy(p1card)
			p2r := player2.copy(p2card)

			winner, err := day22Part2Play(p1r, p2r, game+1)
			// fmt.Printf("Player 2 wins Game %d in Round %d\n", game+1, rounds)
			if err != nil {
				return nil, err
			}
			if winner == p1r {
				player1.push(p1card)
				player1.push(p2card)
			} else {
				player2.push(p2card)
				player2.push(p1card)
			}
		} else {
			if p1card > p2card {
				// p1 wins
				// fmt.Printf("Player 1 wins Round %d\n", rounds)
				player1.push(p1card)
				player1.push(p2card)
			} else {
				// fmt.Printf("Player 2 wins Round %d\n", rounds)
				player2.push(p2card)
				player2.push(p1card)
			}
			if player1.head == nil {
				// player 2 wins !
				// fmt.Printf("Player 2 wins Game %d in Round %d\n", game, rounds)
				return player2, nil
			} else if player2.head == nil {
				// player 1 wins
				// fmt.Printf("Player 1 wins Game %d in Round %d\n", game, rounds)
				return player1, nil
			}
		}
		rounds++
	}

	return nil, fmt.Errorf("too many loops %d", rounds)
}
