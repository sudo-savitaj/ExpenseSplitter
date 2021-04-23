package main

func main() {
	personA := NewPerson("A","A")
	personB := NewPerson("B","B")
	personC := NewPerson("C","C")
	personD := NewPerson("D","D")
	people := []*Person{personA, personB, personC, personD}
	group := NewGroup("trip",people)
	equalSplitStrategy := EqualSplitStrategy{}
	snacksExpense := NewExpense("Snacks",group.people,300,personA,equalSplitStrategy)
	ticketsExpense := NewExpense("Tickets",group.people,100,personB,equalSplitStrategy)
	taxiExpense := NewExpense("Taxi",group.people,200,personD,equalSplitStrategy)



}
