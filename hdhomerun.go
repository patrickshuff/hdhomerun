package main

import "bytes"
import "fmt"
import "log"
import "strings"
import "net"

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
	// The discovery binary was discovered via tcpdump 
	const discovery_bin = (
		"\x00\x02\x00\x0c\x01\x04\x00\x00\x00\x01\x02\x04\xff\xff" +
			"\xff\xff\x4e\x50\x7f\x35")

	// Setup socket that is going to send/receive discovery datagrams
	RAddr,_ := net.ResolveUDPAddr("udp","192.168.174.255:65001")
	ServerAddr, _ := net.ResolveUDPAddr("udp","192.168.174.168:4322")
	listen_conn, _ := net.ListenUDP("udp", ServerAddr)

	listen_conn.WriteTo([]byte(discovery_bin), RAddr)

	// Listen for a response
	buf := make([]byte, 1024)
	for {
		_,addr,err := listen_conn.ReadFromUDP(buf)

		msg := "hdhomerun device %x found at %s\n"
		// e.g. "hdhomerun device 1322F2F9 found at 192.168.174.249"
		// 
		hdhr_ip := strings.Split(addr.String(), ":")[0]
		fmt.Printf(msg, buf[12:16], hdhr_ip)
		
		if err != nil {
			fmt.Println("Error: ",err)
		}
		break
	}
}

func getChannels() {
	fmt.Println("http://<device ip>/lineup.json")
}
