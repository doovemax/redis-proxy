package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"net"
)

func main() {
	l, e := net.Listen("tcp", "127.0.0.1:6378")
	if e == nil {
		for {
			fmt.Println("#######")
			con, err := l.Accept()
			if err == nil {
				go func(con net.Conn) {
					r, _ := parseRequest(con)
					fmt.Println("-----------#", con.RemoteAddr(), "#", r.Name, "#", r.Body.Read(), "#", r.ClientChan, "--------")
					defer con.Close()
				}(con)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(e)
	}

}
