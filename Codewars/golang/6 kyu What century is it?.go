package kata

import (
    "strconv"
    "fmt"
)

func WhatCentury(year string) string {
    yearInt, err := strconv.Atoi(year)
    if err != nil {
        return "0th"
    }
    century := (yearInt + 99) / 100

    suffix := "th"
    if century%100 != 11 && century%10 == 1 {
        suffix = "st"
    } else if century%100 != 12 && century%10 == 2 {
        suffix = "nd"
    } else if century%100 != 13 && century%10 == 3 {
        suffix = "rd"
    }

    return fmt.Sprintf("%d%s", century, suffix)
}
