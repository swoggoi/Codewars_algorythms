package kata

import (
	"math"
	"strconv"
)

// код не работает для больших чисел
func LastDigit(n1, n2 string) int {
	ch := make(chan int)
	go func() {
		s1, err := strconv.Atoi(n1)
		if err != nil {
			ch <- 0
			return
		}
		ch <- s1
	}()

	go func() {
		s2, err := strconv.Atoi(n2)
		if err != nil {
			ch <- 0
			return
		}
		ch <- s2
	}()
	frst_num := <-ch
	second_num := <-ch
	result := math.Pow(float64(frst_num), float64(second_num))
	return int(result) % 10
}

func SLastDigit(n1, n2 string) int {
	if n2 == "0" {
		return 1
	}
	a := int(n1[len(n1)-1] - '0')
	mod := 0
	for i := 0; i < len(n2); i++ {
		mod = (mod*10 + int(n2[i]-'0')) % 4
	}
	if mod == 0 {
		mod = 4
	}
	res := 1
	for i := 0; i < mod; i++ {
		res = (res * a) % 10
	}

	return res
}
