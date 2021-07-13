package cachemap

import (
	"cachemap/priorityqueue"
	"container/list"
	"errors"
	"time"
)

type Cachemap struct {
	maxsize int
	queue   *list.List
	data    map[string]string
	pq      priorityqueue.PQueue
}

func (cm *Cachemap) Init(maxsize int) {
	cm.maxsize = maxsize
	cm.queue = list.New()
	cm.data = make(map[string]string)
}

func (cm *Cachemap) Size() int {
	return cm.queue.Len()
}
func (cm *Cachemap) Pop(key string) {
	if cm.data[key] == "" {
		return
	}
	delete(cm.data, key)
	for i := cm.queue.Front(); i != nil; i.Next() {
		if i.Value.(string) == key {
			cm.queue.Remove(i)
			break
		}
	}
}
func (cm *Cachemap) Set(key string, value string) {
	if cm.Size() > cm.maxsize {
		top := cm.queue.Front()
		cm.queue.Remove(top)
		delete(cm.data, top.Value.(string))
	}
	cm.queue.PushBack(key)
	cm.data[key] = value
}
func (cm *Cachemap) SetValidTimd(key string, timeduration time.Duration) error {
	if cm.data[key] == "" {
		return errors.New("key not found")
	}
	return nil
}

func (cm *Cachemap) Get(key string) string {
	return cm.data[key]
}
