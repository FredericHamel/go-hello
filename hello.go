package hello
import "fmt"

func exclaim(msg1 string, msg2 string) {
  fmt.Printf("%s%s\n", msg1, msg2)
}

func Hi(name string) {
  exclaim("hello ", name)
}

func Salut(name string) {
  exclaim("bonjour ", name)
}

