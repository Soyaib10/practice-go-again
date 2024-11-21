package main

import (
	"fmt"
	"sync"
	"time"
)

var cnt int
var mutex sync.Mutex

func incremtnt() {
	// cnt++
	// if i only use cnt++ [temp := cnt, temp = cnt + 1, cnt = temp] for some go routine when that [] part happens they don't get the incremented value they sould get. thats why the count result do not get 1000 as it should be.

	// now for every time the cnt is locked for every goroutine and others have to wait for their turn
	mutex.Lock()
	cnt++
	mutex.Unlock()


}

func main() {
	for i := 0; i < 1000; i++ {
		go incremtnt()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count : ", cnt)
}