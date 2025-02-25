package priority_queue_test

import (
	"fmt"
	"testing"

	q "github.com/WitnessBro/priority_queue"
	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)
	queue.Put("Kakashi", 4)

	expected := []struct {
		Value    string
		Priority int
	}{
		{"Kakashi", 4},
		{"Sakura", 3},
		{"Saske", 2},
		{"Naruto", 1},
	}

	i := 0
	for !queue.IsEmpty() {
		item := queue.Get()
		assert.Equal(t, expected[i].Priority, item.Priority)
		assert.Equal(t, expected[i].Value, item.Value)
		i += 1
	}
}
func TestGetMaxItem(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)
	queue.Put("Kakashi", 4)

	expected := struct {
		Value    string
		Priority int
	}{"Kakashi", 4}

	item := queue.GetMax()
	assert.EqualValues(t, expected.Priority, item.Priority)
	assert.EqualValues(t, expected.Value, item.Value)
}

// Хз почему пушит не в конец. Или в пуше самого хипа какая-то хрень делается)
func TestPushItem(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)

	newItem := q.NewItem("Kakashi", 4)

	queue.Push(newItem)

	expected := []struct {
		Value    string
		Priority int
	}{
		{"Sakura", 3},
		{"Kakashi", 4},
		{"Saske", 2},
		{"Naruto", 1},
	}
	i := 0
	for !queue.IsEmpty() {
		item := queue.Get()
		assert.Equal(t, expected[i].Priority, item.Priority)
		assert.Equal(t, expected[i].Value, item.Value)
		i += 1
	}

}

// TODO Попает какой-то рандомный элемент
func TestPopItem(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)
	queue.Put("Kakashi", 4)
	queue.Put("Madara", 5)
	queue.Put("Obito", 6)
	queue.Put("Hinata", 7)

	popedItem := queue.Pop().(*q.Item)

	expected := struct {
		Value    string
		Priority int
	}{"Sakura", 3}
	for !queue.IsEmpty() {
		item := queue.Get()
		fmt.Println(item.Priority, item.Index, item.Value)
	}
	fmt.Println(popedItem.Value)
	assert.Equal(t, expected.Priority, popedItem.Priority)
	assert.Equal(t, expected.Value, popedItem.Value)
}
func TestUpdateItem(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	item := queue.Get()
	queue.Update(item, "Saske", 2)

	expected := struct {
		Value    string
		Priority int
	}{"Saske", 2}

	assert.EqualValues(t, expected.Priority, item.Priority)
	assert.EqualValues(t, expected.Value, item.Value)
}

func TestNotEmptyQueueLength(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)
	queue.Put("Kakashi", 4)

	expected := 4

	assert.EqualValues(t, expected, queue.Len())
}
func TestEmptyQueueLength(t *testing.T) {
	queue := q.New()

	expected := 0

	assert.EqualValues(t, expected, queue.Len())
}
func TestSwapItems(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)
	queue.Put("Saske", 2)
	queue.Put("Sakura", 3)
	queue.Put("Kakashi", 4)

	queue.Swap(1, 2)

	expected := []struct {
		Value    string
		Priority int
	}{
		{"Kakashi", 4},
		{"Sakura", 3},
		{"Saske", 2},
		{"Naruto", 1},
	}

	i := 0
	for !queue.IsEmpty() {
		item := queue.Get()
		assert.Equal(t, expected[i].Priority, item.Priority)
		assert.Equal(t, expected[i].Value, item.Value)
		i += 1
	}
}
func TestEmptyQueue(t *testing.T) {
	queue := q.New()

	expected := true

	assert.EqualValues(t, expected, queue.IsEmpty())
}

func TestNotEmptyQueue(t *testing.T) {
	queue := q.New()
	queue.Put("Naruto", 1)

	expected := false

	assert.EqualValues(t, expected, queue.IsEmpty())
}
