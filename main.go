package main

import (
	"fmt"
	"minitask1/totaldue"
)

func main ()  {
	fmt.Println(fmt.Sprintln("Total harga produk Rp.", totaldue.Discount(120000, "50sale")))
	fmt.Println(fmt.Sprintln("Total ongkos kirim Rp.", totaldue.ShippingCost(3)))
}