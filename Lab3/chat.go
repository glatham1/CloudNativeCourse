
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	name    string      //name of user with corresponding channel
	channel chan string //outgoing message channel
}

var clients []client //makes a slice from the struct

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	connected := make(map[client]bool) // all connected clients

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for outgoingMsg := range connected {
				outgoingMsg.channel <- msg
			}

		case cli := <-entering:
			connected[cli] = true //adjusts to who is in the server
			
			//checks the length of the struct to determine how many people are in the server
			if len(clients) == 1 { 
				cli.channel <- "You are the only person in the server."
			} else {
				cli.channel <- "The following people are in the server:"
				//for loop runs through everyone in the server with their corresponding name
				for j := 0; j < len(clients); j++ {
					cli.channel <- clients[j].name
				}
			}

		case cli := <-leaving:
			delete(connected, cli) //deletes specific client from map
			close(cli.channel) //close specific clients channel
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	//asks the user who joined, what their name is, and stores it in a variable
	ch <- "Who are you?"
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()

	//creates an instance of the structure with users name and channel
	cli := client{name: who, channel: ch}
	//adds the user who joined into the slice by appending to the tail
	clients = append(clients, cli)

	//outputs who joined and lets others know as well
	ch <- "You are " + cli.name
	messages <- cli.name + " has arrived"
	entering <- cli

	//sends any messages from client to everyone else in server
	input = bufio.NewScanner(conn)
	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- cli.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
