package domain

import "time"

type Event struct {
	Id   string
	Date time.Time
}

type EventRegistry struct {
	Events map[string]Event
}

func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		Events: make(map[string]Event),
	}
}

func (er *EventRegistry) AddEvent(id string) {
	er.Events[id] = Event{
		Id:   id,
		Date: time.Now(),
	}
}

func (er *EventRegistry) GetEvent(id string) Event {
	return er.Events[id]
}

func (er *EventRegistry) GetAllEvents() []Event {
	events := make([]Event, 0, len(er.Events))
	for _, event := range er.Events {
		events = append(events, event)
	}
	return events
}

func (er *EventRegistry) HasEvent(id string) bool {
	ev := er.GetEvent(id)
	if ev.Id != "" {
		return true
	}
	return false
}
