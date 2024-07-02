package main

func maximumPrimeDifference(nums []int) int {
	mem := *getPrimeMap()
	i, j := 0, len(nums)-1
	findI, findJ := false, false
	for i < j {
		if !findI {
			if !mem[nums[i]] {
				findI = true
			} else {
				i++
			}
		}

		if !findJ {
			if !mem[nums[j]] {
				findJ = true
			} else {
				j--
			}
		}

		if findI && findJ {
			break
		}
	}

	return j - i
}

func getPrimeMap() *[]bool {
	mem := make([]bool, 101)
	mem[1] = true

	for i := 2; i <= 50; i++ {
		for j := 2; j*i <= 100; j++ {
			mem[j*i] = true
		}
	}
	return &mem
}

func main() {
}
