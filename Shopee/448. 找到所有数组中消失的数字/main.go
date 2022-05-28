package main

func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
}

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}

	return ans
}
