package main

/*
	name: 433. 最小基因变化
	link: https://leetcode-cn.com/problems/minimum-genetic-mutation/
*/

func main() {
	// 作为小白能力有限这道题没能自己写出来

	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}

	MinMutation(start, end, bank)
}

// MinMutation 官方de题解
func MinMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	bankMap := map[string]struct{}{}

	for _, v := range bank {
		bankMap[v] = struct{}{}
	}

	if _, ok := bankMap[end]; !ok {
		return -1
	}

	nextList := []string{start}
	for step := 0; len(nextList) != 0; step++ {
		temp := nextList
		nextList = nil
		for _, cur := range temp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if x != y {
						next := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankMap[next]; ok {
							if next == end {
								return step + 1
							}
							delete(bankMap, next)
							nextList = append(nextList, next)
						}
					}
				}
			}
		}
	}

	return -1
}
