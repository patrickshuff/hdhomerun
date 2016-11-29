package main

import "bytes"
import "fmt"
import "log"
import "net"
import "strconv"

var source_port int = 4321

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	// bufio.NewWriter(conn).WriteString(discovery_bin)
	// conn.WriteToUDP(discovery_bin)

	// status, err := bufio.NewReader(conn2).ReadByte()
	logger.Print("Hello, log file!")

	sendUDPProbes()
	//startUDPServer()
	
}

func sendUDPProbes() {
	const discovery_bin = "\x00\x02\x00\x0c\x01\x04\x00\x00\x00\x01\x02\x04\xff\xff\xff\xff\x4e\x50\x7f\x35"
	fmt.Println("Sending out discovery probes...")

	// Set up our discover UDP socket
	LAddr,_ := net.ResolveUDPAddr("udp",fmt.Sprintf(":%s", strconv.Itoa(source_port)))
	RAddr,_ := net.ResolveUDPAddr("udp","192.168.174.255:65001")
	probe_conn, _ := net.DialUDP("udp", LAddr, RAddr)

	fmt.Println("Source: ", probe_conn.LocalAddr())
	fmt.Println("Dest: ", probe_conn.RemoteAddr())

	// Setup socket that is going to receive the response
	ServerAddr, _ := net.ResolveUDPAddr("udp","192.168.174.168:4321")
	listen_conn, err := net.ListenUDP("udp", ServerAddr)

	fmt.Println(ServerAddr)
	fmt.Println(err)
	// fmt.Println("Listening on: ", listen_conn.LocalAddr())
	
	
	// Write the discovery bytes to UDP socket
	fmt.Fprintf(probe_conn, discovery_bin)

	// Listen for a response
	buf := make([]byte, 1024)
	for {
		n,addr,err := listen_conn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)
		
		if err != nil {
			fmt.Println("Error: ",err)
		} 
	}
}

func startUDPServer() {
	fmt.Println("Starting listening server probes...")
	ServerAddr, err := net.ResolveUDPAddr("udp","127.255.255.255:4321")

	conn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	buf := make([]byte, 1024)


	for {
		n,addr,err := conn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)
		
		if err != nil {
			fmt.Println("Error: ",err)
		} 
	}

}

func getChannels() {
	fmt.Println("http://<device ip>/lineup.json")
}
