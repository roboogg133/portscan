package main

import (
	"fmt"
	"net"
	"time"
)

func Scan(ip string, port int) bool {

	adress := fmt.Sprintf("%s:%d", ip, port)
	connec, err := net.DialTimeout("tcp", adress, 3*time.Second)
	if err != nil {
		return false
	} else {
		connec.Close()
		return true
	}
}

func main() {

	var ip string
	var port int
	var portEnd int

	fmt.Println("Enter ip and port range, example: 127.0.0.1 1 100")
	fmt.Scanln(&ip, &port, &portEnd)

	for i := port; i <= portEnd; i++ {

		opened := Scan(ip, i)
		time.Sleep(1 * time.Millisecond)
		if opened == true {

			fmt.Printf("PORT %d OPEN\n", i)
		} else {
			fmt.Printf("PORT%d CLOSED\n", i)
		}
	}
}
