package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	containerv2 "github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

func TestAttribute(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *container.Attribute

		require.Nil(t, x.ToV2())
	})

	t.Run("default values", func(t *testing.T) {
		attr := container.NewAttribute()

		// check initial values
		require.Empty(t, attr.Key())
		require.Empty(t, attr.Value())

		// convert to v2 message
		attrV2 := attr.ToV2()
		require.Empty(t, attrV2.GetKey())
		require.Empty(t, attrV2.GetValue())
	})

	const (
		key   = "key"
		value = "value"
	)

	attr := container.NewAttribute()
	attr.SetKey(key)
	attr.SetValue(value)

	require.Equal(t, key, attr.Key())
	require.Equal(t, value, attr.Value())

	t.Run("test v2", func(t *testing.T) {
		const (
			newKey   = "newKey"
			newValue = "newValue"
		)

		v2 := attr.ToV2()
		require.Equal(t, key, v2.GetKey())
		require.Equal(t, value, v2.GetValue())

		v2.SetKey(newKey)
		v2.SetValue(newValue)

		newAttr := container.NewAttributeFromV2(v2)

		require.Equal(t, newKey, newAttr.Key())
		require.Equal(t, newValue, newAttr.Value())
	})
}

func TestAttributes(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x container.Attributes

		require.Nil(t, x.ToV2())

		require.Nil(t, container.NewAttributesFromV2(nil))
	})

	var (
		keys = []string{"key1", "key2", "key3"}
		vals = []string{"val1", "val2", "val3"}
	)

	attrs := make(container.Attributes, 0, len(keys))

	for i := range keys {
		attr := container.NewAttribute()
		attr.SetKey(keys[i])
		attr.SetValue(vals[i])

		attrs = append(attrs, attr)
	}

	t.Run("test v2", func(t *testing.T) {
		const postfix = "x"

		v2 := attrs.ToV2()
		require.Len(t, v2, len(keys))

		for i := range v2 {
			k := v2[i].GetKey()
			v := v2[i].GetValue()

			require.Equal(t, keys[i], k)
			require.Equal(t, vals[i], v)

			v2[i].SetKey(k + postfix)
			v2[i].SetValue(v + postfix)
		}

		newAttrs := container.NewAttributesFromV2(v2)
		require.Len(t, newAttrs, len(keys))

		for i := range newAttrs {
			require.Equal(t, keys[i]+postfix, newAttrs[i].Key())
			require.Equal(t, vals[i]+postfix, newAttrs[i].Value())
		}
	})
}

func TestNewAttributeFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *containerv2.Attribute

		require.Nil(t, container.NewAttributeFromV2(x))
	})
}
