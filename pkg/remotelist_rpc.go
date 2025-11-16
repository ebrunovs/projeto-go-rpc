package remotelist

import (
	"errors"
	"fmt"
	"sync"
)

type RemoteList struct {
	mu   sync.Mutex
	lists map[int][]int
	//size uint32
}

func (l *RemoteList) Append(args AppendArgs, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _,ok := l.lists[args.ListID]; !ok {
		l.lists[args.ListID] = []int{}
	}

	l.lists[args.ListID] = append(l.lists[args.ListID], args.Value)
	fmt.Printf("Lista %d: %v\n", args.ListID, l.lists[args.ListID])

	//l.size++
	*reply = true
	return nil
}

func (l *RemoteList) Remove(args RemoveArgs, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	list, exists := l.lists[args.ListID]
	if !exists {
		return errors.New("Lista não existe.\n")
	}

	if len(list) == 0 {
		return errors.New("Lista vazia.\n")
	}

	last := list[len(list)-1]
	*reply = last

	l.lists[args.ListID] = list[:len(list)-1]

	fmt.Printf("Item %d removido da lista %d. Estado atual: %v\n",
		last, args.ListID, l.lists[args.ListID])

	return nil
}

func (l *RemoteList) Get(args GetArgs, reply *int) error{
	l.mu.Lock()
	defer l.mu.Unlock()

	list, exists := l.lists[args.ListID]
	if !exists {
		return errors.New("Lista não existe.\n")
	}

	if args.Index < 0 || args.Index >= len(list) {
		return errors.New("Índice fora dos limites.\n")
	}

	*reply = list[args.Index]

	fmt.Printf("Get da lista %d na posição %d → %d\n",
		args.ListID, args.Index, *reply)

	return nil
}

func (l *RemoteList) Size(args SizeArgs, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	list, exists := l.lists[args.ListID]
	if !exists {
		return errors.New("Lista não existe.\n")
	}

	*reply = len(list)

	fmt.Printf("Tamanho da lista %d: %d\n", args.ListID, *reply)

	return nil
}

func NewRemoteList() *RemoteList {
	return &RemoteList{
		lists: make(map[int][]int),
	}
}

type AppendArgs struct {
	ListID int
	Value int
}

type RemoveArgs struct {
	ListID int
}

type GetArgs struct {
	ListID int
	Index int
}

type SizeArgs struct {
	ListID int
}