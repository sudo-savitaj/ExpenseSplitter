package main

type Transaction struct {
	creditor *Person
	debitor *Person
	ammount int
}

func NewTransaction(creditor *Person,debitor *Person,ammount int) *Transaction {
	return &Transaction{creditor,debitor,ammount}
}