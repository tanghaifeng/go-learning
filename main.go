package main

import interface_01 "go-learning/interface-01"

func main() {
	var animal interface_01.Animal = interface_01.Cat{}
	animal.Eat()
	animal.Run()

}
