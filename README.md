## Band Protocol assignment

ข้อมูลนี้เป็นชุด Test ของบริษัท Band Protocol ในตำแหน่ง Software Engineer

ในไฟล์ `test_1.go` จะเป็นการแก้โจทย์ `Boss Baby's Revenge` โดยจะใช้ภาษา Go ในการเขียน โดยจะทำการ for ตัวอักษรที่ได้รับมาเพื่อมาตรวจสอบว่า Boss Baby สามารถแก้แค้นได้ครบหรือไม่และไม่โจมตี the neighborhood kids ก่อน

ในไฟล์ `test_2.go` จะเป็นการแก้โจทย์ `Superman's Chicken Rescue` โดยจะใช้ภาษา Go ในการเขียน โดยจะทำการ for เพื่อวางต่ำแหน่งหลังคาและทำการตรวจสอบว่าในบริเวณขอบเขตของหลังคาสามารถปกป้องไก่ได้กี่ตัวแล้วจดจำจำนวนไก่ที่ปกป้องทำจนครบ Loop แล้วมา Return ค่ามากที่สุดออกมา

ในไฟล์ `test_3.py` จะเป็นการแก้โจทย์ `Transaction Broadcasting and Monitoring Client` โดยจะใช้ภาษา Python ในการเขียนเนื่องจากผมพึ่งฝึกภาษา Go ได้ไม่นานจึงยังไม่คล่องในการเขียนเพื่อติดต่อกับ Http Server จึงเลือกใช้ Python ที่ถนัดมากกว่าก่อนโดยจะทำการสร้างฟังก์ชัน `broadcast_transaction` ในการ Post request payload จากนั้นรับค่า `tx_hash` จาก server มาจากนั้นก็สร้างฟังก์ชัน `check_transaction_status` เพื่อขอสถานะของการทำธุรกรรมออกมา
