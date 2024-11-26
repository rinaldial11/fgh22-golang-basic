package totaldue

// func Discount (total int) int {
// 	voucher1 := 30
// 	voucher2 := 50

// 	if total <= 100000 {
// 		diskon := total * voucher1/100
// 		if  diskon > 30000 {
// 			return total - 30000
// 		} else {
// 			return total -diskon
// 		}
// 	} else {
// 		diskon := total * voucher2/100
// 		if  diskon > 50000 {
// 			return total - 50000
// 		} else {
// 			return total -diskon
// 		}
// 	}
// }

func Discount (total int, code string) int {
	if code == "30sale" {
		diskon := total * 30/100
		if diskon > 15000 {
			return total - 15000
		} else {
			return total - diskon
		}
	} else if code == "50sale" {
		diskon := total * 50/100
		if diskon > 25000 {
			return total - 25000
		} else {
			return total - diskon
		}
	} else {
		return total
	}
}