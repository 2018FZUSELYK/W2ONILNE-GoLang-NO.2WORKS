package main

import "fmt"

type Snowpea interface {
	shoot()
	retard()
}

type Repeatshoot interface {
	repeat()
}

type Repeater struct {
	name string
}

func (t Repeater) shoot(){
	fmt.Printf("%s shoots you!\n",t.name)
}

func (t Repeater) retard(){
	fmt.Printf("%s is repeater, so it can't use retard!\n",t.name)
}

func (t Repeater) repeat(){
	fmt.Printf("%s shoots twice!\n",t.name)
}

func main(){
	REPEATER := Repeater{"Repeater One"}
	REPEATER.shoot()
	REPEATER.repeat()
	REPEATER.retard()
}