package orderedmap_test

import (
	"testing"

	orderedmap "github.com/IQ-tech/go-orderedmap"
	"github.com/stretchr/testify/assert"
)

func Test_OrderedMap(t *testing.T) {
	t.Parallel()

	m := orderedmap.New()

	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	t.Run("Get", func(t *testing.T) {
		t.Run("returns not found error if key is not in the map", func(t *testing.T) {
			_, err := m.Get("key_not_in_the_map")
			assert.Equal(t, orderedmap.ErrNotFound, err)
		})

		t.Run("returns value if key is in the map", func(t *testing.T) {
			value, err := m.Get("one")
			assert.NoError(t, err)
			assert.EqualValues(t, 1, value)
		})
	})

	t.Run("PrevKey", func(t *testing.T) {
		t.Run("returns the previous key if there's one", func(t *testing.T) {
			tests := []struct {
				key      string
				expected string
			}{
				{
					key:      "two",
					expected: "one",
				},
				{
					key:      "three",
					expected: "two",
				},
			}

			for _, tt := range tests {
				key, err := m.PrevKey(tt.key)

				assert.NoError(t, err)

				assert.Equal(t, tt.expected, key)
			}
		})

		t.Run("returns the empty string if there's no previous key", func(t *testing.T) {
			key, err := m.PrevKey("one")

			assert.NoError(t, err)

			assert.Equal(t, "", key)
		})

		t.Run("returns not found error if key is not in the map", func(t *testing.T) {
			_, err := m.PrevKey("key_not_in_the_map")

			assert.Equal(t, orderedmap.ErrNotFound, err)
		})
	})

	t.Run("NextKey", func(t *testing.T) {
		t.Run("returns the next key if there's one", func(t *testing.T) {
			tests := []struct {
				key      string
				expected string
			}{
				{
					key:      "one",
					expected: "two",
				},
				{
					key:      "two",
					expected: "three",
				},
			}

			for _, tt := range tests {
				key, err := m.NextKey(tt.key)

				assert.NoError(t, err)

				assert.Equal(t, tt.expected, key)
			}
		})

		t.Run("returns the empty string if there's no next key", func(t *testing.T) {
			key, err := m.NextKey("three")

			assert.NoError(t, err)

			assert.Equal(t, "", key)
		})

		t.Run("returns not found error if key is not in the map", func(t *testing.T) {
			_, err := m.NextKey("key_not_in_the_map")

			assert.Equal(t, orderedmap.ErrNotFound, err)
		})
	})

	t.Run("LastKey", func(t *testing.T) {
		t.Run("if the map is not empty, returns the key that was added last", func(t *testing.T) {
			assert.Equal(t, "three", m.LastKey())
		})

		t.Run("if the map is empty, returns the empty string", func(t *testing.T) {
			m := orderedmap.New()
			assert.Equal(t, "", m.LastKey())
		})
	})

	t.Run("GetFirstKey", func(t *testing.T) {
		t.Run("if the map is not empty, returns the key that was added first", func(t *testing.T) {
			assert.Equal(t, "one", m.GetFirstKey())
		})

		t.Run("if the map is empty, returns the empty string", func(t *testing.T) {
			m := orderedmap.New()
			assert.Equal(t, "", m.GetFirstKey())
		})
	})

	t.Run("Remove", func(t *testing.T) {
		t.Run("removes element from the map", func(t *testing.T) {
			_, err := m.Get("one")
			assert.NoError(t, err)

			key, err := m.PrevKey("two")
			assert.NoError(t, err)
			assert.Equal(t, "one", key)

			m.Remove("one")

			_, err = m.Get("one")
			assert.Equal(t, orderedmap.ErrNotFound, err)

			key, err = m.PrevKey("two")
			assert.NoError(t, err)
			assert.Equal(t, "", key)
		})
	})

	t.Run("Len", func(t *testing.T) {
		t.Run("returns the number of elements in the map", func(t *testing.T) {
			m := orderedmap.New()

			assert.Equal(t, 0, m.Len())

			m.Set("a", "a")

			assert.Equal(t, 1, m.Len())

			m.Remove("a")

			assert.Equal(t, 0, m.Len())
		})
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("checks if key is in the map", func(t *testing.T) {
			m := orderedmap.New()

			assert.False(t, m.Has("any_key"))

			m.Set("key", "value")

			assert.True(t, m.Has("key"))

			m.Set("other_key", 42)

			assert.True(t, m.Has("other_key"))
		})
	})
}
