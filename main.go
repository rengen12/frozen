package main
  
 import (
	"fmt"
	"io"
	"log"
	"net"
	//"os"
	 "bufio"
	 //"io/ioutil"
	// "net/textproto"
	 "strings"
 )
 
type server struct {
	host		string
	
	users		[]string
}

type user struct {
	nick		string
	rname		string
	password	string
	email		string
	conn 		string
	chanels		[]string

}

  func ExampleListener() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
	  //fmt.Println(conn)
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			fmt.Println(c)
		io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
  }

  func handle_conn(conn net.Conn, srv server) {
  	srv.host = "localhost"

	  for {
		  // will listen for message to process ending in newline (\n)
		  message, _ := bufio.NewReader(conn).ReadString('\n')
		  // output message received
		  fmt.Print("Message Received:", string(message))
		  // sample process for string received
		  newmessage := strings.ToUpper(message)
		  // send new string back to client
		  conn.Write([]byte(newmessage + "\n"))
	  }
	  //io.Copy(conn, conn)
	  //status, _ := bufio.NewReader(conn).ReadString(' ')
	  //fmt.Print("status", status)

  }

func main() {
	
 l, err := net.Listen("tcp", ":6667")
 var srv server
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle_conn(conn, srv)
	}
}