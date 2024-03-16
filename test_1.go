package main

import (
	"errors"
	"fmt"
)

func checkRevenge(shots string) (string, error) {
	// เริ่มต้นนับจำนวนการยิง
	countShots := 0
	if len(shots) == 0 {
		return "", errors.New("Parameter 'shots' Must not be empty")
	}

	// ถ้าตำแหน่งแรกของ shots เป็น 'R' แสดงว่าเป็นคนที่ยิงก่อน จึงส่งคืน "Bad boy"
	if shots[0] == 'R' {
		return "Bad boy", nil
	} else {
		// วนลูปตรวจสอบทุกการยิงใน shots
		for _, shot := range shots {
			// ถ้าเจอ 'S' ให้เพิ่มจำนวนการยิง
			if shot == 'S' {
				countShots++
			} else if shot == 'R' { // ถ้าเจอ 'R' และมีการยิงก่อนหน้าอยู่ ให้ลดจำนวนการยิง
				if countShots > 0 {
					countShots--
				}
			}
		}
	}

	// หลังจากตรวจสอบครบทุกการยิงแล้ว ถ้าจำนวนการยิงเท่ากับ 0 แสดงว่าเป็น "Good Boy" ไม่เช่นนั้นเป็น "Bad Boy"
	if countShots == 0 {
		return "Good Boy", nil
	} else {
		return "Bad Boy", nil
	}
}

func main() {
	shots := "SRSSRRR"
	result, err := checkRevenge(shots)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
