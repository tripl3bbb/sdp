package deposit

func main() {
	product := "CS:GO"
	depWay := 3

	var payment Deposit
	switch depWay {
	case 1:
		payment = NewCardPayment("4400440044000000", "123")
	case 2:
		payment = NewBitcoinPayment()
	case 3:
		payment = NewSkinsPayment()
	}

	processOrder(product, payment)
}

func processOrder(game string, payment Deposit) {
	//
	err := payment.Pay()
	if err != nil {
		return
	}
}

// ----

type Deposit interface {
	Pay() error
}

// ----

type ByCard struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Deposit {
	return &ByCard{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *ByCard) Pay() error {
	//
	return nil
}

type ByBitcoin struct {
}

func NewBitcoinPayment() Deposit {
	return &ByBitcoin{}
}

func (p *ByBitcoin) Pay() error {
	//
	return nil
}

type BySkins struct {
}

func NewSkinsPayment() Deposit {
	return &BySkins{}
}

func (p *BySkins) Pay() error {
	//
	return nil
}
