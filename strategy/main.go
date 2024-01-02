package main

import "fmt"

/*
*
StategyInterfaces
*/
type FlyBehaviorInterface interface {
	performFly()
}

type QuackBehaviorInterface interface {
	performQuack()
}

/*
*
Concrete strategies implementing FlyBehaviorInterface
*/
type FlyWithWings struct {
}

func (fww FlyWithWings) performFly() {
	fmt.Println("I know how to fly")
}

type FlyNoWay struct {
}

func (fnw FlyNoWay) performFly() {
	fmt.Println("I can't fly")
}

/*
*
Concrete strategies implementing QuackBehaviorInterface
*/
type Quack struct {
}

func (q Quack) performQuack() {
	fmt.Println("Quack")
}

type Squeack struct {
}

func (sq Squeack) performQuack() {
	fmt.Println("Squeack")
}

/*
*
Class that calls the concrete strategies
*/
type Duck struct {
	FlyBehavior   FlyBehaviorInterface
	QuackBehavior QuackBehaviorInterface
}

// Set strategy
func (d *Duck) SetFlyBehavior(flyBehavior FlyBehaviorInterface) {
	d.FlyBehavior = flyBehavior
}

// Set strategy
func (d *Duck) SetFlyQuackBehavior(quackBehavior QuackBehaviorInterface) {
	d.QuackBehavior = quackBehavior
}

// Calls strategy
func (d Duck) Fly() {
	d.FlyBehavior.performFly()
}

// Calls strategy
func (d Duck) Quack() {
	d.QuackBehavior.performQuack()
}

func (d Duck) Eat() {
	fmt.Println("nhac, nhac, nhac.... delicious")
}

func (d Duck) Swin() {
	fmt.Println("swining is easy for ducks")
}

func main() {

	Duck := Duck{}
	// I`m a dumb duck
	Duck.SetFlyBehavior(FlyNoWay{})
	Duck.Fly()

	// Now i`ve learned to fly
	Duck.SetFlyBehavior(FlyWithWings{})
	Duck.Fly()

	Duck.Eat()
	Duck.Swin()
}
