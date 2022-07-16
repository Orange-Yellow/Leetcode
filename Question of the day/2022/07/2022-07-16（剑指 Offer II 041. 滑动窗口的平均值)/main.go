package main

/*
	name: 剑指 Offer II 041. 滑动窗口的平均值
	link: https://leetcode.cn/problems/qIsx9U/
*/

func main() {
	m := Constructor(3)
	m.Next(1)
	m.Next(10)
	m.Next(3)
	m.Next(5)
}

type MovingAverage struct {
	Arr  []int
	Size int
	Sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{Size: size}
}

func (m *MovingAverage) Next(val int) float64 {
	if len(m.Arr) >= m.Size {
		m.Sum -= m.Arr[0]
		m.Arr = m.Arr[1:]
	}
	m.Sum += val
	m.Arr = append(m.Arr, val)
	return float64(m.Sum) / float64(len(m.Arr))
}
