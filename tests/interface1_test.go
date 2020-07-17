package tests

import (
	interface_01 "go-learning/interface-01"
	"testing"
)

func TestInterface1(t *testing.T) {
	var animal interface_01.Animal = interface_01.Cat{}
	animal.Eat()
	animal.Run()
}
