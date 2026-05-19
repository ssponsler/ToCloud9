package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

//go:generate mockery --name=FriendsServiceProducer
type FriendsServiceProducer interface {
	StatusChange(payload *FriendEventStatusChangePayload) error
	FriendAdded(payload *FriendEventAddedPayload) error
	FriendRemoved(payload *FriendEventRemovedPayload) error
	NoteUpdated(payload *FriendEventNoteUpdatePayload) error
}

type friendsServiceProducerNatsJSON struct {
	conn *nats.Conn
	ver  string
}

func NewFriendsServiceProducerNatsJSON(conn *nats.Conn, ver string) FriendsServiceProducer {
	return &friendsServiceProducerNatsJSON{
		conn: conn,
		ver:  ver,
	}
}

func (s *friendsServiceProducerNatsJSON) StatusChange(payload *FriendEventStatusChangePayload) error {
	return s.publish(FriendEventStatusChange, payload)
}

func (s *friendsServiceProducerNatsJSON) FriendAdded(payload *FriendEventAddedPayload) error {
	return s.publish(FriendEventAdded, payload)
}

func (s *friendsServiceProducerNatsJSON) FriendRemoved(payload *FriendEventRemovedPayload) error {
	return s.publish(FriendEventRemoved, payload)
}

func (s *friendsServiceProducerNatsJSON) NoteUpdated(payload *FriendEventNoteUpdatePayload) error {
	return s.publish(FriendEventNoteUpdate, payload)
}

func (s *friendsServiceProducerNatsJSON) publish(e FriendsServiceEvent, payload interface{}) error {
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
