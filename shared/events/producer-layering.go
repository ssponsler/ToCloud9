package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

//go:generate mockery --name=LayeringProducer
type LayeringProducer interface {
	CreatureStateChanged(payload *LayeringEventCreatureStateChangedPayload) error
}

type layeringProducerNatsJSON struct {
	conn *nats.Conn
	ver  string
}

func NewLayeringProducerNatsJSON(conn *nats.Conn, ver string) LayeringProducer {
	return &layeringProducerNatsJSON{
		conn: conn,
		ver:  ver,
	}
}

func (s *layeringProducerNatsJSON) CreatureStateChanged(payload *LayeringEventCreatureStateChangedPayload) error {
	return s.publish(LayeringEventCreatureStateChanged, payload)
}

func (s *layeringProducerNatsJSON) publish(e LayeringEvent, payload interface{}) error {
	msg := EventToSendGenericPayload{
		Version:   s.ver,
		EventType: int(e),
		Payload:   payload,
	}

	d, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	return s.conn.Publish(e.SubjectName(), d)
}
