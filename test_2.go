package main

import (
	"errors"
	"fmt"
	"sort"
)

func maxChickens(n int, k int, positions []int) (int, error) {
	if n != len(positions) {
		return 0, errors.New("Try sending data again") // ส่งError
	}
	// เรียงลำดับตำแหน่งไก่
	sort.Ints(positions)

	// กำหนดค่า Max ของไก่
	maxChickens := 0

	for _, i := range positions {
		// หาตำแหน่งของหลังคาที่เริ่มต้นปกป้องไก่
		startRoof := i
		// หาตำแหน่งของหลังคาที่สิ้นสุดปกป้องไก่
		endRoof := i + k

		// นับจำนวนไก่ที่อยู่ในช่วงหลังคา
		chickensInRange := 0
		for _, p := range positions {
			if startRoof <= p && p < endRoof {
				chickensInRange++
			}
		}

		// อัปเดตค่าของ maxChickens
		if chickensInRange > maxChickens {
			maxChickens = chickensInRange
		}
	}

	return maxChickens, nil
}

func main() {
	n := 6
	k := 10
	positions := []int{1, 11, 30, 34, 35, 37}
	result, err := maxChickens(n, k, positions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
