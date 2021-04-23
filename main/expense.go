package main

type SplitStrategy interface {
	split(ammount int, numberOfPeople int) int
}

type Expense struct {
	typeOfExpense string
	people        []*Person
	ammount       int
	paidBy        *Person
	transaction   []*Transaction
	splitStrategy SplitStrategy
}

func NewExpense(typeOfExpense string,people []*Person,ammount int,paidBy *Person,splitStrategy SplitStrategy) *Expense {
	return &Expense{typeOfExpense:typeOfExpense,people:people,ammount:ammount,paidBy:paidBy,splitStrategy:splitStrategy}
}

func (this *Expense) split(){
	ammountToBePaid := this.splitStrategy.split(this.ammount, len(this.people))
	for _, person := range this.people {
		if !this.samePerson(person) {
			this.transaction = append(this.transaction, NewTransaction(this.paidBy,person,ammountToBePaid))
		}
	}
}

func (this *Expense) samePerson(person *Person) bool {
	return person.id == this.paidBy.id
}












