package main

func main() {

}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 1
	for ; fast < n; fast++ {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow
}
