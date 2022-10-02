package subscribe

type SObserver struct {
	name string
}

func (s *SObserver) Update(t string) {
	// TODO IMPLEMENT ME
	println("SObserver:", s.name, "has been updated,", "received subject string:", t)
}
