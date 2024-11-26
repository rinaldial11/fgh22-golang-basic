package totaldue

func ShippingCost (distance int) int {
	if distance > 2 {
		return 7000 + (distance - 2) * 2000
	}
	return 7000
}