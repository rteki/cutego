package cutego

import "fmt"

type Subscribers []func(interface{})

type EventManager struct {
	Name     string
	Handlers map[string]Subscribers
}

func NewEventManager(name string) *EventManager {
	var em *EventManager = new(EventManager)

	(*em).Name = name
	(*em).Handlers = make(map[string]Subscribers)

	newEventManager(name)

	eventManagers[name] = em

	return em
}

func (em *EventManager) On(eventName string, subscriber func(interface{})) {

	if em.Handlers[eventName] == nil {
		em.Handlers[eventName] = Subscribers{}
	}

	for _, sub := range em.Handlers[eventName] {
		if &sub == &subscriber {
			fmt.Println("CuteGo Go Warning: Subscriber for " + eventName + " is already registered.")
			return
		}
	}

	em.Handlers[eventName] = append(em.Handlers[eventName], subscriber)

}

func (em *EventManager) Call(eventName string, value interface{}) {
	callQt(em.Name, eventName, value)
}
