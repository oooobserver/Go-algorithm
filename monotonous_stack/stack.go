package monotonousstack

//
// Purpose: O(n) to find the next bigger element of one element, suit for one dimension array
// Detail: Use a stack to constantly store the index of element
//         There are two order, from big to small or reverse, each suit for different scenario
//         The typical logical whenever found a element that violate the order of stack
//         perfrom the logical here
// Example: Leetcode 84, Leetcode 42
//

// Function template
// Export is for avoid compiler warning
func Template(nums []int) {
	stack := []int{}

	for i, n := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < n {
			// Do thing at here
			// Here means find the first element after the stack top
			// that bigger than it
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
}
