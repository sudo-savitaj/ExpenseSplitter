package main

type Group struct {
	name string
	people []*Person
	expenses []Expense
}

func NewGroup(name string, people []*Person) *Group {
	return &Group{name: name, people: people}
}

func (this *Group)AddExpense(expense Expense)  {
	this.expenses = append(this.expenses, expense)
}

