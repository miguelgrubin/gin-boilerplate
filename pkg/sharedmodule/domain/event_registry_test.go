package domain_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
	"github.com/stretchr/testify/assert"
)

func TestEventRegistry(t *testing.T) {
	registry := domain.NewEventRegistry()
	eventID := "por.que.asi"
	registry.AddEvent(eventID)

	event := registry.GetEvent(eventID)

	assert.NotNil(t, event, "Event should be registered and retrievable")
	assert.Equal(t, eventID, event.ID, "Event ID should match")
}

func TestEventRegistryGetAllEvents(t *testing.T) {
	registry := domain.NewEventRegistry()
	firstEventID := "por.que.asi"
	secondEventID := "porque.si"
	registry.AddEvent(firstEventID)
	registry.AddEvent(secondEventID)

	events := registry.GetAllEvents()

	assert.NotNil(t, events, "Events should be registered and retrievable")
	assert.Len(t, events, 2, "There should be two registered events")
	assert.Equal(t, firstEventID, events[0].ID, "First event ID should match")
	assert.Equal(t, secondEventID, events[1].ID, "Second event ID should match")
}

func TestEventRegistryHasEvent(t *testing.T) {
	registry := domain.NewEventRegistry()
	eventID := "por.que.asi"
	registry.AddEvent(eventID)

	hasEvent := registry.HasEvent(eventID)
	hasNonExistentEvent := registry.HasEvent("non.existent.event")

	assert.True(t, hasEvent, "Registry should confirm existence of the added event")
	assert.False(t, hasNonExistentEvent, "Registry should confirm non-existence of a non-added event")
}
