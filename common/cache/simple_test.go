package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	cache := NewSimple(nil)

	cache.Put("A", "Foo")
	assert.Equal(t, "Foo", cache.Get("A"))
	assert.Nil(t, cache.Get("B"))
	assert.Equal(t, 1, cache.Size())

	cache.Put("B", "Bar")
	cache.Put("C", "Cid")
	cache.Put("D", "Delt")
	assert.Equal(t, 4, cache.Size())

	assert.Equal(t, "Bar", cache.Get("B"))
	assert.Equal(t, "Cid", cache.Get("C"))
	assert.Equal(t, "Delt", cache.Get("D"))

	cache.Put("A", "Foo2")
	assert.Equal(t, "Foo2", cache.Get("A"))

	cache.Put("E", "Epsi")
	assert.Equal(t, "Epsi", cache.Get("E"))
	assert.Equal(t, "Foo2", cache.Get("A"))

	cache.Delete("A")
	assert.Nil(t, cache.Get("A"))
}

func TestSimpleGenerics(t *testing.T) {
	key := keyType{
		dummyString: "some random key",
		dummyInt:    59,
	}
	value := "some random value"

	cache := NewSimple(nil)
	cache.Put(key, value)

	assert.Equal(t, value, cache.Get(key))
	assert.Equal(t, value, cache.Get(keyType{
		dummyString: "some random key",
		dummyInt:    59,
	}))
	assert.Nil(t, cache.Get(keyType{
		dummyString: "some other random key",
		dummyInt:    56,
	}))
}

func TestSimpleCacheConcurrentAccess(t *testing.T) {
	cache := NewSimple(nil)
	values := map[string]string{
		"A": "foo",
		"B": "bar",
		"C": "zed",
		"D": "dank",
		"E": "ezpz",
	}

	for k, v := range values {
		cache.Put(k, v)
	}

	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)

		// concurrent get and put
		go func() {
			defer wg.Done()

			<-start

			for j := 0; j < 1000; j++ {
				cache.Get("A")
				cache.Put("A", "fooo")
			}
		}()

		// concurrent iteration
		go func() {
			defer wg.Done()

			<-start

			for j := 0; j < 50; j++ {
				it := cache.Iterator()
				for it.HasNext() {
					_ = it.Next()
				}
				it.Close()
			}
		}()
	}

	close(start)
	wg.Wait()
}

func TestSimpleRemoveFunc(t *testing.T) {
	ch := make(chan bool)
	cache := NewSimple(&SimpleOptions{
		RemovedFunc: func(i interface{}) {
			_, ok := i.(*testing.T)
			assert.True(t, ok)
			ch <- true
		},
	})

	cache.Put("testing", t)
	cache.Delete("testing")
	assert.Nil(t, cache.Get("testing"))

	timeout := time.NewTimer(time.Millisecond * 300)
	select {
	case b := <-ch:
		assert.True(t, b)
	case <-timeout.C:
		t.Error("RemovedFunc did not send true on channel ch")
	}
}

func TestSimpleIterator(t *testing.T) {
	expected := map[string]string{
		"A": "Alpha",
		"B": "Beta",
		"G": "Gamma",
		"D": "Delta",
	}

	cache := NewSimple(nil)

	for k, v := range expected {
		cache.Put(k, v)
	}

	actual := map[string]string{}

	it := cache.Iterator()
	for it.HasNext() {
		entry := it.Next()
		// nolint:revive
		actual[entry.Key().(string)] = entry.Value().(string)
	}
	it.Close()
	assert.Equal(t, expected, actual)

	it = cache.Iterator()
	for i := 0; i < len(expected); i++ {
		entry := it.Next()
		// nolint:revive
		actual[entry.Key().(string)] = entry.Value().(string)
	}
	it.Close()
	assert.Equal(t, expected, actual)
}
