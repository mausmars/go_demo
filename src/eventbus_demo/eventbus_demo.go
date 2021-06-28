package main

import (
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
)

type EventLogin struct {
	id int32
}

type EventPay struct {
	id int32
}

func loginHandle1(event *EventLogin) {
	fmt.Printf("loginHandle1 %d\n", event.id)
}

func loginHandle2(event *EventLogin) {
	fmt.Printf("loginHandle2 %d\n", event.id)
}

func payHandle(event *EventPay) {
	fmt.Printf("payHandle %d\n", event.id)
}

func main() {
	bus := eventbus.New()
	bus.Subscribe("login", loginHandle1)
	bus.Subscribe("login", loginHandle2)
	bus.Subscribe("pay", payHandle)

	e1 := &EventLogin{
		id: 1,
	}
	e2 := &EventPay{
		id: 1,
	}

	bus.Publish("login", e1)
	bus.Publish("pay", e2)

	bus.Unsubscribe("login", loginHandle1)
	bus.Publish("login", e1)

	bus.Unsubscribe("login", loginHandle2)
	bus.Unsubscribe("pay", payHandle)
}
