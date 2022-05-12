package main

func main() {

}

// 官方题解
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	prev := -1
	for i := range flowerbed {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			if count >= n {
				return true
			}
			prev = i
		}
	}

	l := len(flowerbed)
	if prev < 0 {
		count += (l + 1) / 2
	} else {
		count += (l - prev - 1) / 2
	}

	return count >= n
}

// 自己题解
func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	count := 0
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			count++
		}
		if count >= n {
			return true
		}
	}
	return false
}
