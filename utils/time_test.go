package utils

import (
	"fmt"
	"testing"
)

func TestGetMonthFirstDay(t *testing.T) {
	re := GetCurrentMonthFirstDay()
	fmt.Println(re.Format("2006-01-02"))
}

func TestGetCurrentMonth(t *testing.T) {
	re := GetCurrentMonth()
	fmt.Println(re)
}
