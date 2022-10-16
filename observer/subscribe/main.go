package subscribe

type Subject interface {
	Subscribe(o Observer) (bool, error)
	Unsubscribe(o Observer) (bool, error)
	Notify() (bool, error)
}

// Observer Interface
type Observer interface {
	Update(string)
}

func main() {
	stockMonitor := &Monitor{
		ticker: "AAPL",
		price:  0.0,
	}

	observerA := &SObserver{
		name: "Observer A",
	}
	observerB := &SObserver{
		name: "Observer B",
	}

	stockMonitor.Subscribe(observerA)
	stockMonitor.Unsubscribe(observerB)

	stockMonitor.Notify()

	stockMonitor.SetPrice(500)

	stockMonitor.Unsubscribe(observerA)

	stockMonitor.SetPrice(528)
}
