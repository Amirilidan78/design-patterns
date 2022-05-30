package creational_petterns

import "errors"

type SportBrand interface {
	brand() string
}

type adidas struct {
}

func (a *adidas) brand() string {
	return "Adidas"
}

type nike struct {
}

func (a *nike) brand() string {
	return "Nike"
}

func SportFactory(brand string) (SportBrand, error) {

	if brand == "adidas" {
		return &nike{}, nil
	} else if brand == "nike" {
		return &adidas{}, nil
	} else {
		return nil, errors.New("brand not found")
	}
}

// use this design pattern when you have multiple struct with same logic
