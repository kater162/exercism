package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	storeUnits := map[string]int{}
	storeUnits["quarter_of_a_dozen"] = 3
	storeUnits["half_of_a_dozen"] = 6
	storeUnits["dozen"] = 12
	storeUnits["small_gross"] = 120
	storeUnits["gross"] = 144
	storeUnits["great_gross"] = 1728

	return storeUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	emptyBill := make(map[string]int)
	return emptyBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	_, e := bill[item]
	if e {
		bill[item] = bill[item] + units[unit]
		return true
	} else {
		bill[item] = units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, a := bill[item]
	_, b := units[unit]

	if !a || !b {
		return false
	}

	switch {
	case (bill[item] - units[unit]) < 0:
	return false
	case (bill[item] - units[unit]) == 0:
	delete(bill, item)
	return true
	default:
	bill[item] = bill[item] - units[unit]
	return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
