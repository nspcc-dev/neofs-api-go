package object

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetNotification(t *testing.T) {
	o := new(Object)

	var ni = NotificationInfo{
		epoch: 10,
		topic: "test",
	}

	WriteNotificationInfo(o, ni)

	var foundEpoch, foundTopic bool

	for _, attr := range o.GetHeader().GetAttributes() {
		switch key := attr.GetKey(); key {
		case SysAttributeTickEpoch:
			require.Equal(t, false, foundEpoch)

			uEpoch, err := strconv.ParseUint(attr.GetValue(), 10, 64)
			require.NoError(t, err)

			require.Equal(t, ni.Epoch(), uEpoch)
			foundEpoch = true
		case SysAttributeTickTopic:
			require.Equal(t, false, foundTopic)
			require.Equal(t, ni.Topic(), attr.GetValue())
			foundTopic = true
		}
	}

	require.Equal(t, true, foundEpoch && foundTopic)
}

func TestGetNotification(t *testing.T) {
	o := new(Object)

	attr := []*Attribute{
		{SysAttributeTickEpoch, "10"},
		{SysAttributeTickTopic, "test"},
	}

	h := new(Header)
	h.SetAttributes(attr)

	o.SetHeader(h)

	t.Run("No error", func(t *testing.T) {
		ni, err := GetNotificationInfo(o)
		require.NoError(t, err)

		require.Equal(t, uint64(10), ni.Epoch())
		require.Equal(t, "test", ni.Topic())
	})
}

func TestIntegration(t *testing.T) {
	o := new(Object)

	var (
		ni1 = NotificationInfo{
			epoch: 10,
			topic: "",
		}
		ni2 = NotificationInfo{
			epoch: 11,
			topic: "test",
		}
	)

	WriteNotificationInfo(o, ni1)
	WriteNotificationInfo(o, ni2)

	t.Run("double set", func(t *testing.T) {
		ni, err := GetNotificationInfo(o)
		require.NoError(t, err)

		require.Equal(t, ni2.epoch, ni.Epoch())
		require.Equal(t, ni2.topic, ni.Topic())
		require.Equal(t, 2, len(o.GetHeader().GetAttributes()))
	})
}
