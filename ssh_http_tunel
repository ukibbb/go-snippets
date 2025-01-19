package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct {
	w      io.Writer
	donech chan struct{}
}

// communication mechanism between http and ssh handler
var tunnels = map[int]chan Tunnel{}

func main() {

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			idstr := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idstr)
			// checks if tunnel is created in ssh session
			tunnel, ok := tunnels[id]

			if !ok {
				w.Write([]byte("tunnel does not exist."))
				return
			}
			donech := make(chan struct{})
			// sending tunnel to ssh session
			tunnel <- Tunnel{
				w:      w,
				donech: donech,
			}
			<-donech

		})
		log.Fatal(http.ListenAndServe(":3000", nil))
	}()

	ssh.Handle(func(s ssh.Session) {
		// creates rand tunnel id
		id := rand.Intn(math.MaxInt)
		// initializes new channel
		tunnels[id] = make(chan Tunnel)

		fmt.Println("tunnel id -> ", id)

		// waits for tunnel from http handler
		tunnel := <-tunnels[id]

		fmt.Println("Channel is ready.")

		// once tunnel is ready
		// copies data from the ssh session to the http response writer
		_, err := io.Copy(tunnel.w, s)
		if err != nil {
			log.Fatal(err)
		}
		// singnals completion closing donech
		close(tunnel.donech)

		s.Write([]byte("We are done!"))
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))

}
