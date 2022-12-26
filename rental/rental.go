package rental

type Rental struct {
	movie      Movie
	daysRented int
}

func NewRental(movie Movie, daysRented int) Rental {
	return Rental{
		movie:      movie,
		daysRented: daysRented,
	}
}
func (r Rental) DaysRented() int {
	return r.daysRented
}
func (r Rental) Movie() Movie {
	return r.movie
}

// จะเกิด Inappropriate Intimacy กับ Movie
// func (r Rental) Charge() float64 {
// 	return r.movie.Price.Charge(r.daysRented)
// }

func (r Rental) Charge() float64 {
	return r.movie.Charge(r.daysRented)
}
