package subscribe

import (
	"errors"
	"strconv"
)

type Monitor struct {
	ticker string
	price  float64

	observers []Observer
}

func (s *Monitor) Subscribe(o Observer) (bool, error) {

	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *Monitor) Unsubscribe(o Observer) (bool, error) {

	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (s *Monitor) Notify() (bool, error) {
	for _, observer := range s.observers {
		observer.Update(s.String())
	}
	return true, nil
}

func (s *Monitor) SetPrice(price float64) {
	s.price = price
	s.Notify()
}

func (s *Monitor) String() string {
	convertFloatToString := strconv.FormatFloat(s.price, 'f', 2, 64)
	return "Monitor: " + s.ticker + " $" + convertFloatToString
}
