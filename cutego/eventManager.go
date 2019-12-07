package cutego


type EventManager struct {
	Name string
	Handlers map[string]func(map[string]interface{})
}


func NewEventManager(name string) *EventManager {
	var em *EventManager = new(EventManager)
	
	(*em).Name = name
	(*em).Handlers = make(map[string]func(map[string]interface{}))

	newEventManager(name)

	eventManagers[name] = em

	return em
}

func (em *EventManager) On(eventName string, handler func(map[string]interface{})) {
	em.Handlers[eventName] = handler
}