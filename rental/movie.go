package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Charger interface {
	Charge(daysRented int) float64
	PriceCode() int
}

type Childrens struct {
	priceCode int
}

func (c Childrens) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}
func (c Childrens) PriceCode() int {
	return c.priceCode
}

func CreateChildrens() Childrens {
	return Childrens{
		priceCode: CHILDRENS,
	}
}

type Regular struct {
	priceCode int
}

func (r Regular) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}
func (r Regular) PriceCode() int {
	return r.priceCode
}

func CreateRegular() Regular {
	return Regular{
		priceCode: REGULAR,
	}
}

type NewRelease struct {
	priceCode int
}

func (n NewRelease) Charge(daysRented int) float64 {
	return float64(daysRented) * 3.0
}
func (n NewRelease) PriceCode() int {
	return n.priceCode
}

func CreateNewRelease() NewRelease {
	return NewRelease{
		priceCode: NEW_RELEASE,
	}
}

type Movie struct {
	title     string
	priceCode int
	Charger   Charger
}

func NewM(title string, charger Charger) Movie {
	return Movie{
		title:     title,
		priceCode: charger.PriceCode(),
		Charger:   charger,
	}
}

func NewMovie(title string, priceCode int) Movie {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}
func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
