package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "port to listen on")
	host = flag.String("h", "localhost", "host to listen on")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conencted")
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done // bloquea el programa hasta que la funcion anonima haya terminando de copiar el contenido.

}

func CopyContent(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy: %v\n", err)
		os.Exit(1)
	}
}
