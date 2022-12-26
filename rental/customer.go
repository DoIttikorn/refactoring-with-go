package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}

}

func (rcvr Customer) Name() string {
	return rcvr.name
}

func (r Rental) Charge() float64 {
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}

func getPoint(r Rental) int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func getTotalAmount(rental []Rental) float64 {
	result := 0.0
	for _, r := range rental {
		result += r.Charge()
	}
	return result
}

func getTotalPoint(rentals []Rental) int {
	result := 0
	for _, r := range rentals {
		result += getPoint(r)
	}

	return result
}

func (c Customer) Statement() string {
	frequentRenterPoints := getTotalPoint(c.rentals)
	totalAmount := getTotalAmount(c.rentals)

	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), r.Charge())
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
