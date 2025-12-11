package kata

func QueueTime(customers []int, n int) int {
	tills := make([]int, n)
	for _, customer := range customers {
		minIndex := 0
		for i := 0; i < n; i++ {
			if tills[i] < tills[minIndex] {
				minIndex = i
			}
		}
		tills[minIndex] += customer
	}
	max := tills[0]
	for i := 1; i < n; i++ {
		if tills[i] > max {
			max = tills[i]
		}
	}
	return max
}
