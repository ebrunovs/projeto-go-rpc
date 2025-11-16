package remotelist

import (
	"errors"
	"fmt"
	"sync"
)

type RemoteList struct {
	mu   sync.Mutex
	lists map[string][]int
	//size uint32
}

func (l *RemoteList) Append(args AppendArgs, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _,ok := l.lists[args.ListID]; !ok {
		l.lists[args.ListID] = []int{}
	}

	l.lists[args.ListID] = append(l.lists[args.ListID], args.Value)
	fmt.Printf("Lista %s: %v\n", args.ListID, l.lists[args.ListID])

	//l.size++
	*reply = true
	return nil
}

func (l *RemoteList) Remove(args RemoveArgs, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	list, exists := l.lists[args.ListID]
	if !exists {
		return errors.New("Lista não existe.")
	}

	if len(list) == 0 {
		return errors.New("Lista vazia.")
	}

	last := list[len(list)-1]
	*reply = last

	l.lists[args.ListID] = list[:len(list)-1]

	fmt.Printf("Item %d removido da lista %s. Estado atual: %v\n",
		last, args.ListID, l.lists[args.ListID])

	return nil
}

func NewRemoteList() *RemoteList {
	return &RemoteList{
		lists: make(map[string][]int),
	}
}

type AppendArgs struct {
	ListID string
	Value int
}

type RemoveArgs struct {
	ListID string
}
