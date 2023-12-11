package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
}

func (Database) createSingleConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creating Done")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB connection")
		db = &Database{}
		db.createSingleConnection()
	} else {
		fmt.Println("Db Already created")
	}
	return db
}
func main2() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
