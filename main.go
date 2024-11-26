package main

import (
	"fmt"
	"minitask1/totaldue"
)

func main ()  {
	shippingCost := totaldue.ShippingCost(3)
	productCost := totaldue.Discount(120000, "50sale")

	fmt.Println(fmt.Sprintln("Total harga produk = Rp.", productCost) + fmt.Sprintln("Total ongkos kirim = Rp.", shippingCost))
	fmt.Println("Total yang harus dibayar = Rp.", shippingCost + productCost)

}