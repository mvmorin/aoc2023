package day07

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func card_to_index(c byte) int {
	if c <= '9' {
		return int(c - '2')
	}
	switch c {
	case 'T':
		return 8
	case 'J':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}

	panic("invalid card")
	return -1
}

type Hand struct {
	bid     int
	most byte
	next_most byte
	jokers byte
	cards   [5]byte
}

func (h *Hand) get_hand_rank() int {
	switch 10*(h.most+h.jokers) + h.next_most {
	case 11:
		return 0
	case 21:
		return 1
	case 22:
		return 2
	case 31:
		return 3
	case 32:
		return 4
	case 41:
		return 5
	case 50:
		return 6
	}

	panic("invalid hand")
	return -1
}

func (h *Hand) compare(t *Hand, j_wilds bool) int {
	h_rank := h.get_hand_rank()
	t_rank := t.get_hand_rank()
	if h_rank > t_rank {
		return 1
	} else if h_rank < t_rank {
		return -1
	}

	for i := 0; i < 5; i++ {
		hc := h.cards[i]
		tc := t.cards[i]
		hci := card_to_index(hc)
		tci := card_to_index(tc)

		if j_wilds {
			if hc == 'J' {
				hci = -1
			}
			if tc == 'J' {
				tci = -1
			}
		}

		if hci > tci {
			return 1
		} else if hci < tci {
			return -1
		}
	}
	return 0
}

func parse_hand(line string, j_wilds bool) Hand {
	j_idx := card_to_index('J')
	hand := Hand{}

	var summary [13]byte
	for i, r := range line {
		if i < 5 {
			hand.cards[i] = byte(r)
			summary[card_to_index(byte(r))]++
		} else if i > 5 {
			hand.bid = 10*hand.bid + int(r-'0')
		}
	}

	var most, next_most, jokers byte = 0, 0, 0
	for i, count := range summary {
		if j_wilds && i == j_idx {
			jokers = count
		} else if count >= most {
			next_most = most
			most = count
		} else if count >= next_most {
			next_most = count
		}
	}
	hand.most = most
	hand.next_most = next_most
	hand.jokers = jokers

	return hand
}


func Prob1() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	hands := make([]Hand, 0, 1000)
	for scanner.Scan() {
		hand := parse_hand(scanner.Text(), false)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		cmp := hands[i].compare(&hands[j], false)
		return cmp < 0
	})

	win_sum := 0
	for i, h := range hands {
		win := h.bid * (i+1)
		win_sum += win
	}

	fmt.Println(win_sum)
	return win_sum
}

func Prob2() int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	hands := make([]Hand, 0, 1000)
	for scanner.Scan() {
		hand := parse_hand(scanner.Text(), true)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		cmp := hands[i].compare(&hands[j], true)
		return cmp < 0
	})

	win_sum := 0
	for i, h := range hands {
		win := h.bid * (i+1)
		win_sum += win
	}

	fmt.Println(win_sum)
	return win_sum
}
