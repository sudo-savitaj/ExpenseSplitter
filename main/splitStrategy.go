package main

type EqualSplitStrategy struct {
}

func (this EqualSplitStrategy) split(ammount int, numberOfPeople int) int {
	return ammount/ numberOfPeople
}
