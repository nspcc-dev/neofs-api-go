package object

import (
	"errors"
	"fmt"
	"strconv"
)

// SysAttributePrefix is a prefix of key to system attribute.
const SysAttributePrefix = "__NEOFS__"

const (
	// SysAttributeUploadID marks smaller parts of a split bigger object.
	SysAttributeUploadID = SysAttributePrefix + "UPLOAD_ID"

	// SysAttributeExpEpoch tells GC to delete object after that epoch.
	SysAttributeExpEpoch = SysAttributePrefix + "EXPIRATION_EPOCH"

	// SysAttributeTickEpoch defines what epoch must produce object
	// notification.
	SysAttributeTickEpoch = SysAttributePrefix + "TICK_EPOCH"

	// SysAttributeTickTopic defines what topic object notification
	// must be sent to.
	SysAttributeTickTopic = SysAttributePrefix + "TICK_TOPIC"
)

// NotificationInfo groups information about object notification
// that can be written to object.
//
// Topic is an optional field.
type NotificationInfo struct {
	epoch uint64
	topic string
}

// Epoch returns object notification tick
// epoch.
func (n NotificationInfo) Epoch() uint64 {
	return n.epoch
}

// SetEpoch sets object notification tick
// epoch.
func (n *NotificationInfo) SetEpoch(epoch uint64) {
	n.epoch = epoch
}

// Topic return optional object notification
// topic.
func (n NotificationInfo) Topic() string {
	return n.topic
}

// SetTopic sets optional object notification
// topic.
func (n *NotificationInfo) SetTopic(topic string) {
	n.topic = topic
}

// WriteNotificationInfo writes NotificationInfo to the Object via attributes. Object must not be nil.
//
// Existing notification attributes are expected to be key-unique, otherwise undefined behavior.
func WriteNotificationInfo(o *Object, ni NotificationInfo) {
	h := o.GetHeader()
	if h == nil {
		h = new(Header)
		o.SetHeader(h)
	}

	var (
		attrs = h.GetAttributes()

		epoch = strconv.FormatUint(ni.Epoch(), 10)
		topic = ni.Topic()

		changedEpoch bool
		changedTopic bool
		deleteIndex  = -1
	)

	for i := range attrs {
		switch attrs[i].GetKey() {
		case SysAttributeTickEpoch:
			attrs[i].SetValue(epoch)
			changedEpoch = true
		case SysAttributeTickTopic:
			changedTopic = true

			if topic == "" {
				deleteIndex = i
				break
			}

			attrs[i].SetValue(topic)
		}

		if changedEpoch && changedTopic {
			break
		}
	}

	if deleteIndex != -1 {
		// approach without allocation/waste
		// coping works since the attributes
		// order is not important
		attrs[deleteIndex] = attrs[len(attrs)-1]
		attrs = attrs[:len(attrs)-1]
	}

	if !changedEpoch {
		index := len(attrs)
		attrs = append(attrs, Attribute{})
		attrs[index].SetKey(SysAttributeTickEpoch)
		attrs[index].SetValue(epoch)
	}

	if !changedTopic && topic != "" {
		index := len(attrs)
		attrs = append(attrs, Attribute{})
		attrs[index].SetKey(SysAttributeTickTopic)
		attrs[index].SetValue(topic)
	}

	h.SetAttributes(attrs)
}

// ErrNotificationNotSet means that object does not have notification.
var ErrNotificationNotSet = errors.New("notification for object is not set")

// GetNotificationInfo looks for object notification attributes. Object must not be nil.
// Returns ErrNotificationNotSet if no corresponding attributes
// were found.
//
// Existing notification attributes are expected to be key-unique, otherwise undefined behavior.
func GetNotificationInfo(o *Object) (*NotificationInfo, error) {
	var (
		foundEpoch bool
		ni         = new(NotificationInfo)
	)

	for _, attr := range o.GetHeader().GetAttributes() {
		switch key := attr.GetKey(); key {
		case SysAttributeTickEpoch:
			epoch, err := strconv.ParseUint(attr.GetValue(), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("could not parse epoch: %w", err)
			}

			ni.SetEpoch(epoch)

			foundEpoch = true
		case SysAttributeTickTopic:
			ni.SetTopic(attr.GetValue())
		}
	}

	if !foundEpoch {
		return nil, ErrNotificationNotSet
	}

	return ni, nil
}
