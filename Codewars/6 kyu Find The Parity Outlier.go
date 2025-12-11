package kata

func FindOutlier(integers []int) int {
	oddCount := 0
	evenCount := 0
	for i := 0; i < len(integers); i++ {
		if integers[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	lookForEven := oddCount > evenCount

	for _, num := range integers {
		if lookForEven && num%2 == 0 {
			return num
		}
		if !lookForEven && num%2 != 0 {
			return num
		}
	}

	return 0
}

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
