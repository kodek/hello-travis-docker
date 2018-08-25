package main
import "fmt"
import "time"
import hello "github.com/kodek/hello-travis-docker"

func main() {
	for {
		fmt.Println(hello.SayHi())
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
