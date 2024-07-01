package dp

/*
	0-1 package

1. dp[i][j]=x means use item within 0~i, and the size is j, x is the max values under this condition
2. You can either choose current item or not choose, dp[i][j]=max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
No matter what should i-1, because current stat is from last stat
3. dp[*][0] should be all 0, dp[0][*] should  be 0 within dp[0][weights[0]]~dp[0][SIZE]
4. We normally choose the outer is item and the inner is the size
5. This step is test ...
*/
func ZeroOnePackage(weights []int, values []int, size int) int {
	n := len(weights)
	mem := make([][]int, n)

	for i := range n {
		mem[i] = make([]int, size+1)
	}

	// When idx is 0, only can use one item
	for i := weights[0]; i <= size; i++ {
		mem[0][i] = values[0]
	}

	// Dp start

	// The outer is item
	for i := 1; i < n; i++ {
		for j := 0; j <= size; j++ {
			if j < weights[i] {
				mem[i][j] = mem[i-1][j]
			} else {
				mem[i][j] = max(mem[i-1][j], mem[i-1][j-weights[i]]+values[i])
			}
		}
	}

	// The outer is weight
	// for j := 0; j <= size; j++ {
	// 	for i := 1; i < n; i++ {
	// 		if j < weights[i] {
	// 			mem[i][j] = mem[i-1][j]
	// 		} else {
	// 			mem[i][j] = max(mem[i-1][j], mem[i-1][j-weights[i]]+values[i])
	// 		}
	// 	}
	// }

	return mem[n-1][size]
}

/*
Why can just use one dimension array to replace two-dimension array?

Because the in dp[i][j]=max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
We can replace dp[i][*] with dp[i-1][*], but dp[i-1][*] is calculated stepwise and depends on some items in dp[i][*]
So we should careful about iterate direction
*/
func ZeroOnePackagePro(weights []int, values []int, size int) int {
	n := len(weights)
	mem := make([]int, size+1)

	for i := 0; i < n; i++ {
		// In order to avoid repeat computing
		// For example: weights[0]=1, values[0]=15
		// mem[1] = max(mem[1], mem[(1-1)]+15) = 15
		// mem[2] = max(mem[2], mem[(2-1)]+15) = 30...
		// The item0 is repeat calculated
		for j := size; j >= weights[i]; j-- {
			mem[j] = max(mem[j], mem[j-weights[i]]+values[i])
		}
	}
	return mem[size]
}

func CompletePackage(weights []int, values []int, size int) int {
	n := len(weights)
	mem := make([]int, size+1)

	// There is a different when get combinatio and  arrangement
	// Combinatio: {1,5} and {5,1} are same
	// Arrangement: {1,5} and {5,1} are different
	// So their iteration are different

	// Combinatio
	// The outer is items, the inner is package capacity
	for i := 0; i < n; i++ {
		// Because the item can use multiple times
		// So it can iterate form to high
		for j := weights[i]; j <= size; j++ {
			mem[j] = max(mem[j], mem[j-weights[i]]+values[i])
		}
	}

	// Arrangement
	// The outer is package capacity, the inner is items
	// The reason doing this is if you have a array [1,3...]
	// If your outer is item, the arrangement you get always 1,3
	// So to break the order of original array, the inner is array

	// for i := 0; i <= size; i++ {
	// 	for _, weight := range weights {
	// 		if i < weight {
	// 			continue
	// 		}
	// 		// Arrangement logical
	// 	}
	// }

	return mem[size]
}

// Multiple package problem just like zero-one package
// Except an item with numnber limit
// We can turn it into zero-one package by extending the number of the item
func MultiplePackage(weights []int, values []int, nums []int, size int) int {
	n := len(weights)
	mem := make([]int, size+1)

	for i := 0; i < n; i++ {
		for j := size; j >= weights[i]; j-- {
			// Add a loop to cope with the item number
			for k := 1; k <= nums[i] && j-(k*weights[i]) >= 0; k++ {
				mem[j] = max(mem[j], mem[j-(k*weights[i])]+k*values[i])
			}
		}
	}
	return mem[size]
}

// DP of the tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Leetcode 337
//  func rob(root *TreeNode) int {
//     res := helper(root)
//     return max(res[0], res[1])
// }

// // 0: not steal, 1: steal
// func helper(node *TreeNode) []int{
//     if node == nil{
//         return []int{0,0}
//     }

//     left := helper(node.Left)
//     right := helper(node.Right)

//     // Not steal this node
//     res1 := max(left[0], left[1]) + max(right[0], right[1])
//     // Steal this node
//     res2 := node.Val + left[0] + right[0]

//     return []int{res1,res2}
// }

/*
	Leetcode 123

There are five states:
- do nothing
- first have stock
- first not have stock
- second have stock
- second not have stock

for these five state:
- dp[i][0] = dp[i-1][0]
- dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
- dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
- dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
- dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])

Intialization:
- dp[0][0] should be 0
- dp[0][1] should be -prices[0]
- dp[0][2] is sold on the day bought, so should be 0
- dp[0][3] at first thought, consider should not be initialzed. But
think as that day already sold the first stock, so it should be -prices[0]
- dp[0][4] should be 0
*/
func maxProfit(prices []int) int {
	return maxProfitK(prices, 2)
}

// TODO: reduce the first state
func maxProfitK(prices []int, k int) int {
	size := 1 + 2*k
	mem := make([]int, size)
	for i := 1; i <= k; i++ {
		mem[i*2-1] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j < size; j++ {
			tmp := 0
			if j%2 == 0 { // sold
				tmp = mem[j-1] + prices[i]
			} else {
				tmp = mem[j-1] - prices[i]
			}
			mem[j] = max(mem[j], tmp)
		}
	}
	return mem[size-1]
}
