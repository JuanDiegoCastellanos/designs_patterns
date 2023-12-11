package main

import "fmt"

type Topic interface {
	register(observer Observer)
	bradcast()
}
type Observer interface {
	getId() string
	updateValue(string)
}

// Item --> no dispoible
// Item --> Avisar --> Hay item
type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailable() {
	fmt.Printf("Item %s is available \n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}
func (ec *EmailClient) updateValue(val string) {
	fmt.Printf("Sending an Email, %s available from client %s\n", val, ec.id)
}

func mainx2() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "32ef",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)

	nvidiaItem.updateAvailable()
}
