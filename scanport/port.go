package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "", "site to scann")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "stage.prontohousing.io", port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Port", port, "is open")
		}(i)
	}
	fmt.Println(*site)
}
