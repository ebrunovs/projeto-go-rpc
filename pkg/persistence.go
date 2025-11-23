package remotelist

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sync"
)

type PersistenceManager struct {
	logPath      string
	snapshotPath string
	mu           *sync.Mutex
}

func NewPersistenceManager(logPath, snapshotPath string, mu *sync.Mutex) *PersistenceManager {
	return &PersistenceManager{logPath, snapshotPath, mu}
}

func (pm *PersistenceManager) WriteLog(op string, listID int, value int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	f, err := os.OpenFile(pm.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func (pm *PersistenceManager) CreateSnapshot(lists map[int][]int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	fmt.Println("Criando snapshot...\n")
	data, err := json.MarshalIndent(lists, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(pm.snapshotPath, data, 0644); err != nil {
		return err
	}
	return os.WriteFile(pm.logPath, []byte{}, 0644)
}

func (pm *PersistenceManager) LoadSnapshot(lists *map[int][]int) error {
	data, err := os.ReadFile(pm.snapshotPath)
	if err != nil {
		return nil
	}
	return json.Unmarshal(data, lists)
}

func (pm *PersistenceManager) ApplyLog(lists *map[int][]int) error {
	file, err := os.Open(pm.logPath)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		op := parts[0]
		listID, _ := strconv.Atoi(parts[1])
		switch op {
		case "append":
			if len(parts) != 3 {
				continue
			}
			value, _ := strconv.Atoi(parts[2])
			(*lists)[listID] = append((*lists)[listID], value)
		case "remove":
			list := (*lists)[listID]
			if len(list) > 0 {
				(*lists)[listID] = list[:len(list)-1]
			}
		}
	}
	return scanner.Err()
}
