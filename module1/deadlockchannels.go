// simple deadlock
package module1

import (
	"fmt"
	//"sync"
)

// causes deadlock
/* func main() {

c:=make(chan int) // unbuffered channel

c<-32
value := <-c

fmt.Println("valeu",value)


} */

// no deadlock
func Deadlockc() {

	c := make(chan int) // unbuffered channel

	go func() {

		c <- 42
	}()
	value := <-c
	fmt.Println("value", value)

}
