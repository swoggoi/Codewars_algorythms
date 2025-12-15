package kata

import (
	"strconv"
	"strings"
)

func ipToInt(ip string) uint32 {
	parts := strings.Split(ip, ".")
	var result uint32 = 0

	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(parts[i])
		result = result*256 + uint32(num)
	}

	return result
}
func IpsBetween(start, end string) uint32 {
	startInt := ipToInt(start)
	endInt := ipToInt(end)
	return endInt - startInt
}
