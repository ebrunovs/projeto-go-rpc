package remotelist

import (
	"errors"
	"fmt"
	"sync"
	"encoding/json"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type RemoteList struct {
	mu   sync.Mutex
	lists map[int][]int
}

func (l *RemoteList) Append(args AppendArgs, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _,ok := l.lists[args.ListID]; !ok {
		l.lists[args.ListID] = []int{}
	}

	l.lists[args.ListID] = append(l.lists[args.ListID], args.Value)
	fmt.Printf("Lista %d: %v\n", args.ListID, l.lists[args.ListID])

	l.writeLog("append", args.ListID, args.Value)

	*reply = true
	return l.save()
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

	l.writeLog("remove", args.ListID, 0)

	return l.save()
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
	l := &RemoteList{
		lists: make(map[int][]int),
	}
	l.load()
	return l
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

func (l *RemoteList) save() error {
	data, err := json.MarshalIndent(l.lists, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/lists.json", data, 0644)
}

func (l *RemoteList) loadArchive() error{
	data, err := os.ReadFile("data/lists.json")
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, &l.lists)
}

func (l *RemoteList) apllyLog() error {
	file, err := os.Open("data/operations.log")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Split(line, " ")

		op := part[0]
		listID, _ := strconv.Atoi(part[1])
		
		switch op {
		case "append":
			value, _ := strconv.Atoi(part[2])
			l.lists[listID] = append(l.lists[listID], value)

		case "remove":
			if len(l.lists[listID]) > 0 {
				l.lists[listID] = l.lists[listID][:len(l.lists[listID])-1]
			}
		}
	}

	return scanner.Err()
}

func(l *RemoteList) load() error {
	if err := l.loadArchive(); err != nil {
		return err
	}
	return l.apllyLog()
}

func (l *RemoteList) writeLog(op string, listID int, value int) error {
	f, err := os.OpenFile("data/operations.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	var line string
	if op == "append" {
		line = fmt.Sprintf("append %d %d\n", listID, value)
	} else {
		line = fmt.Sprintf("remove %d\n", listID)
	}

	_, err = f.Write([]byte(line))
	return err
}