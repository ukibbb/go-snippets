package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop(conn)

	}
}
func (fs *FileServer) readLoop(conn net.Conn) {
	// no streaming version
	// buf := make([]byte, 1024)
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		// n, err := conn.Read(buf)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}
		// file := buf[:n]
		fmt.Println(buf)
		fmt.Printf("recevied %d bytes over the network\n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	// to stream file instead of conn.Write
	// copy file to the connectionn
	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))

	// n, err := conn.Write(file)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}

func main() {

	go func() {
		time.Sleep(4 * time.Second)
		// buf size is 1024
		// - in result it reciving file in chunks first 1024 then 976
		sendFile(2000000)
	}()

	server := &FileServer{}
	server.start()
}
