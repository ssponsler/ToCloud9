package events

import "fmt"

// LayeringEvent is event type that worldserver emits for layering.
type LayeringEvent int

const (
	// LayeringEventCreatureStateChanged fires when a creature becomes alive/dead.
	LayeringEventCreatureStateChanged LayeringEvent = iota + 1
)

// SubjectName is key that nats uses.
func (e LayeringEvent) SubjectName() string {
	switch e {
	case LayeringEventCreatureStateChanged:
		return "layer.creature.state"
	}
	panic(fmt.Errorf("unk event %d", e))
}

// LayeringEventCreatureStateChangedPayload carries creature state updates.
type LayeringEventCreatureStateChangedPayload struct {
	RealmID       uint32
	MapID         uint32
	ZoneID        uint32
	AreaID        uint32
	LayerID       string
	SpawnID       uint32
	CreatureEntry uint32
	IsAlive       bool
	X             float32
	Y             float32
	Z             float32
	TimestampUnix int64
}
