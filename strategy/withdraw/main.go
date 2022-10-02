package withdraw

func main() {
	WithdrawWay := 3

	var withdraw Withdrawal
	switch WithdrawWay {
	case 1:
		withdraw = NewCardCashout("4400440044000000", "123")
	case 2:
		withdraw = NewCryptoCashout()
	case 3:
		withdraw = NewSkinsCashout()
	}

	processCashOut(withdraw)
}

func processCashOut(withdraw Withdrawal) {
	//
	err := withdraw.Cashout()
	if err != nil {
		return
	}
}

// ----

type Withdrawal interface {
	Withdraw() error
}

// ----

type ToCard struct {
	cardNumber, cvv string
}

func NewCardCashout(cardNumber, cvv string) Withdrawal {
	return &ToCard{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *ToCard) Pay() error {
	//
	return nil
}

type ToCrypto struct {
}

func NewCryptoCashout() Withdrawal {
	return &ToCrypto{}
}

func (p *ToCrypto) Pay() error {
	//
	return nil
}

type ToSkins struct {
}

func NewSkinsCashout() Withdrawal {
	return &ToSkins{}
}

func (p *ToSkins) Withdrawal() error {
	//
	return nil
}
