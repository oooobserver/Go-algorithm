package concurrent

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// A channel only can out the message
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func action() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle_conn(conn)
	}
}

func delete_item(slice []string, target string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}

	return slice
}

func broadcaster() {
	clients := make(map[client]struct{})

	// Record names
	var cli_names []string

	for {
		select {
		// If is the message, send it to the writer of every connection
		case msg := <-messages: // The new connection send its name through messages channel
			name := strings.Fields(msg)
			if strings.Contains(msg, "arrived") {
				cli_names = append(cli_names, name[0])
			} else {
				cli_names = delete_item(cli_names, name[0])
			}

			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering: // The new connection send its receive channel
			clients[cli] = struct{}{}
			cli <- "current clients: "
			for _, name := range cli_names {
				cli <- name
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}

func handle_conn(conn net.Conn) {
	// Create the channel for this connection routine and its writer routine
	ch := make(chan string)
	go client_writer(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Address is " + who
	// Send the message to broadcaster, let it record this new connection
	// Send the writer channel to the broadcaster
	messages <- who + " has arrived"
	entering <- ch

	// Constantly read from this connection and send the message to the broadcaster
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// Take the message form the broadcaster and its connection
func client_writer(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
