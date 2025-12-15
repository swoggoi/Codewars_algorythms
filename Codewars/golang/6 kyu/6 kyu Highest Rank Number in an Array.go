package kata

func HighestRank(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	maxCount := 0
	result := 0
	for num, count := range counts {
		if count > maxCount || (count == maxCount && num > result) {
			maxCount = count
			result = num
		}
	}
	return -1
}
