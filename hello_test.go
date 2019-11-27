package hello
import (
  "bytes"
  "io"
  "os"
  "testing"
)

func TestHi(t *testing.T) {
  stdout := os.Stdout

  r, w, err := os.Pipe()
  if err != nil {
    t.Error("Unable to create Pipe")
    return
  }
  os.Stdout = w

  defer func () {
    os.Stdout = stdout
  }()

  resultC := make(chan string)
  go func() {
    defer w.Close()
    Hi("Max")
  }()

  go func() {
    var buf bytes.Buffer
    io.Copy(&buf, r)
    resultC <- buf.String()
  }()

  expected := "hello Max\n"

  result := <-resultC
  if result != expected {
    t.Errorf("hello.Hi(\"Max\") => %s; expected \"hello Max\"", result)
  }
}

func TestSalut(t *testing.T) {
  stdout := os.Stdout

  r, w, err := os.Pipe()
  if err != nil {
    t.Error("Unable to create Pipe")
    return
  }
  os.Stdout = w

  defer func () {
    os.Stdout = stdout
  }()

  resultC := make(chan string)
  go func() {
    defer w.Close()
    Salut("Max")
  }()

  go func() {
    var buf bytes.Buffer
    io.Copy(&buf, r)
    resultC <- buf.String()
  }()

  expected := "bonjour Max\n"

  result := <-resultC
  if result != expected {
    t.Errorf("hello.Salut(\"Max\") => %s; expected \"%s\"", result, expected)
  }
}
