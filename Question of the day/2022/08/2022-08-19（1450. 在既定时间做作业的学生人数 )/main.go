package main

/*
	name: 1450. 在既定时间做作业的学生人数
	link: https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time/
*/

func main() {

}

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	for i, s := range startTime {
		if s <= queryTime && queryTime <= endTime[i] {
			ans++
		}
	}
	return
}
