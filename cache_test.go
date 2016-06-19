package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCaching(t *testing.T) {
	assert := assert.New(t)
	InitCache()
	AddItemToCache("item1", "value1")
	AddItemToCache("item2", "value2")
	AddItemToCache("item3", "value3")
	assert.Equal(GetItemFromCache("item1"), "value1", "should be equal")
	assert.Equal(GetItemFromCache("item2"), "value2", "should be equal")
	assert.Equal(GetItemFromCache("item3"), "value3", "should be equal")
	removeItem("item1")
	assert.Equal(GetItemFromCache("item1"), "", "should be equal")
	AddItemToCache("item2", "value2-2")
	assert.Equal(GetItemFromCache("item2"), "value2-2", "should be equal")
}

func TestCaching10000(t *testing.T) {
	t0 := time.Now()
	assert := assert.New(t)
	InitCache()
	i := 0
	for i < 10000 {
		AddItemToCache("item"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		assert.Equal(GetItemFromCache("item"+strconv.Itoa(i)), "value"+strconv.Itoa(i), "should be equal")
		AddItemToCache("item"+strconv.Itoa(i), "valuechanged"+strconv.Itoa(i))
		assert.Equal(GetItemFromCache("item"+strconv.Itoa(i)), "valuechanged"+strconv.Itoa(i), "should be equal")
		i++
	}
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
