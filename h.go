package main
  
  import (
	"fmt"
	"io"
	"log"
	"net"
	"container/list"
	//"os"
  )
  
type person struct {
	nick		string
	rname		string
	password	string
	email		string
}

type user struct {
	
	chanels []string
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

func adduser(nick, rname, password, email string) {
	fmt.Println(lst)
}


func main() {
	
 l, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go parse_input(conn, &users)
	}
}
	//adduser("name", "", "", "")

	users.PushBack(person{"name", "rname", "password", "email"})

	/*for e := users.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}*/

	/*for e := users.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	adduser("name1", "", "", "")
	for e := users.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	adduser("name", "qq", "ww", "ee")*/
	for e := users.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
	//ExampleListener();
  }