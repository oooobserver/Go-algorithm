package dp

import (
	"fmt"

	"shenye.com/util"
)

func win(nums []int) int {
	dealer := nums[0] + nums[1]
	player := 0
	for i := 2; i < len(nums); i++ {
		player += nums[i]
	}
	if player <= 21 && player > dealer {
		return 1
	}
	return 0
}

func black_jack(cards []int, pointer int, dp []int, m *int) {
	l := len(cards)
	if pointer+4 > l {
		return
	}

	// Can perform 5 cards
	if pointer+5 < l {
		dp[pointer+4] = win(cards[pointer:pointer+5]) + dp[pointer]
		*m = max(*m, dp[pointer+4])
		black_jack(cards, pointer+5, dp, m)
	}

	dp[pointer+3] = win(cards[pointer:pointer+4]) + dp[pointer]
	*m = max(*m, dp[pointer+3])
	black_jack(cards, pointer+4, dp, m)
}

func play() {
	size := util.RandomInt(10, 20)
	cards := []int{}
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		card := util.RandomInt(1, 10)
		cards = append(cards, card)
	}
	fmt.Println(cards)

	m := 0
	black_jack(cards, 0, dp, &m)

	fmt.Println(dp)
	fmt.Println(m)
}
