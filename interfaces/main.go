package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreetings()
	abuseThem()
}

type englishBot struct {
}

type spanishBot struct {
}

type Console struct {
}

func (c Console) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func logWriter(p io.Writer) {
	p.Write([]byte("Welcome to India"))
}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	c := Console{}

	io.Copy(c, resp.Body)

	// buffer := make([]byte, 99999)

	// resp.Body.Read(buffer)

	// fmt.Println(string(buffer))

	// eb := englishBot{}
	// sb := spanishBot{}

	// printEGreeting(eb)
	// printSGreeting(sb)

	// printGreeting(eb)
	// printGreeting(sb)

	// x := "Its time to leave"

	// fmt.Println([]byte(x))

	// reader := strings.NewReader("hello")

	// fmt.Println(reader)

	// // create buffer of 4 bytes
	// buffer := make([]byte, 5)

	// fmt.Println(buffer)

	// copy(buffer, "hello")

	// fmt.Println(buffer)

	// // FIRST READ

	// n, err := reader.Read(buffer)

	// fmt.Println("FIRST READ")
	// fmt.Println("n:", n)
	// fmt.Println("err:", err)
	// fmt.Println("buffer bytes:", buffer)
	// fmt.Println("buffer as string:", string(buffer))
	// fmt.Println("valid bytes only:", string(buffer[:n]))

	// fmt.Println()

	// // SECOND READ
	// n, err = reader.Read(buffer)

	// fmt.Println("SECOND READ")
	// fmt.Println("n:", n)
	// fmt.Println("err:", err)
	// fmt.Println("buffer bytes:", buffer)
	// fmt.Println("buffer as string:", string(buffer))
	// fmt.Println("valid bytes only:", string(buffer[:n]))

	// fmt.Println()

	// // THIRD READ (EOF)
	// n, err = reader.Read(buffer)

	// fmt.Println("THIRD READ")
	// fmt.Println("n:", n)
	// fmt.Println("err:", err)

	// if err == io.EOF {
	// 	fmt.Println("No more data left")
	// }

}

func printGreeting(b bot) {
	b.getGreetings()
	b.abuseThem()
}

// func printEGreeting(b englishBot) {
// 	b.getGreetings()
// }

// func printSGreeting(b spanishBot) {
// 	b.getGreetings()
// }

func (eb englishBot) getGreetings() {
	fmt.Println("Hi there!")
}

func (sb spanishBot) getGreetings() {
	fmt.Println("Hola!")
}

func (eb englishBot) abuseThem() {
	fmt.Println("Hi there!")
}

func (sb spanishBot) abuseThem() {
	fmt.Println("Hola!")
}
