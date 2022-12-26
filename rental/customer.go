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

type PriceManager struct {
	point  int
	amount float64
}

type RentalManager struct {
	title  string
	amount float64
}

func (c Customer) renderPlainText(amount float64, point int) string {
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		rm := RentalManager{
			title:  r.Movie().Title(),
			amount: r.Charge(),
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", rm.title, rm.amount)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", amount)
	result += fmt.Sprintf("You earned %v frequent renter points", point)
	return result
}

func (c Customer) Statement() string {
	pm := PriceManager{
		point:  getTotalPoint(c.rentals),
		amount: getTotalAmount(c.rentals),
	}

	return c.renderPlainText(pm.amount, pm.point)
}
