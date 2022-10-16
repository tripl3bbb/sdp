package main

type state interface { //interface of the state
	addItem(int) error           //adding
	requestItem() error          //requesting
	insertMoney(money int) error //inserting
	dispenseItem() error         //giving item
}
