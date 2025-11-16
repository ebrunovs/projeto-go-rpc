package remotelist

import (
	//"errors"
	"fmt"
	"sync"
)

type RemoteList struct {
	mu   sync.Mutex
	lists map[string][]int
	//size uint32
}

func (r *RemoteList) Append(args AppendArgs, reply *bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _,ok := r.lists[args.ListID]; !ok {
		r.lists[args.ListID] = []int{}
	}

	r.lists[args.ListID] = append(r.lists[args.ListID], args.Value)
	fmt.Printf("Lista %s: %v\n", args.ListID, r.lists[args.ListID])

	//l.size++
	*reply = true
	return nil
}

// func (l *RemoteList) Remove(arg int, reply *int) error {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()
// 	if len(l.list) > 0 {
// 		*reply = l.list[len(l.list)-1]
// 		//l.list = l.list[:len(l.list)-1]
// 		fmt.Println(l.list)
// 	} else {
// 		return errors.New("empty list")
// 	}
// 	return nil
// }

func NewRemoteList() *RemoteList {
	return &RemoteList{
		lists: make(map[string][]int),
	}
}

type AppendArgs struct {
	ListID string
	Value int
}
