package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Moving the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var m Mover
	var l Locker
	var ml MoveLocker

	m = bike{}
	m.Move()

	l = bike{}
	l.Lock()
	l.Unlock()

	// Will throw unless you use
	// 	b, ok := m.(bike)
	if b, ok := m.(bike); ok {
		ml = b
		fmt.Println(ml)
	}

}
