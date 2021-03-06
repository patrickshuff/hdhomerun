package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

var HDHR_PORT int = 65001

func main() {
	app := cli.NewApp()
	app.Name = "hdhomerun"
	app.Usage = "Control the hdhomerun on your network"
	app.Commands = []cli.Command{
		{
			Name:    "discover",
			Aliases: []string{"d"},
			Usage:   "Discover HDHR devices on network",
			Action: func(c *cli.Context) error {
				discoverHDHR()
				return nil
			},
		},
		{
			Name:    "channels",
			Aliases: []string{"c"},
			Usage:   "Print out list of channels",
			Action: func(c *cli.Context) error {
				getChannels()
				return nil
			},
		},
		{
			Name:    "scanchannels",
			Aliases: []string{"c"},
			Usage:   "Force HDHR to rescan for new channels",
			Action: func(c *cli.Context) error {
				scanChannels()
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func discoverHDHR() string {
	// The discovery binary was discovered via tcpdump
	const discovery_bin = ("\x00\x02\x00\x0c\x01\x04\x00\x00\x00\x01" +
		"\x02\x04\xff\xff\xff\xff\x4e\x50\x7f\x35")

	// Setup socket that is going to send/receive discovery datagrams
	RAddr, _ := net.ResolveUDPAddr("udp", "192.168.174.255:65001")
	ServerAddr, _ := net.ResolveUDPAddr("udp", "192.168.174.168:")
	listen_conn, err := net.ListenUDP("udp", ServerAddr)

	if err != nil {
		fmt.Println("Error opening UDP socket: ", err)
		return ""
	}

	listen_conn.WriteTo([]byte(discovery_bin), RAddr)

	// Grab response for a buffer
	buf := make([]byte, 1024)
	for {
		_, addr, err := listen_conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Error reading from UDP socket: ", err)
			return ""
		}

		msg := "hdhomerun device %x found at %s\n"
		// e.g. "hdhomerun device 1322F2F9 found at 192.168.174.249"
		hdhr_ip := strings.Split(addr.String(), ":")[0]
		hdhr_dev_name := buf[12:16]

		fmt.Printf(msg, hdhr_dev_name, hdhr_ip)

		return hdhr_ip
	}
}

func getChannels() {
	hdhr_ip := discoverHDHR()

	lineup_url := "http://%s/lineup.json"
	resp, err := http.Get(fmt.Sprintf(lineup_url, hdhr_ip))

	if err != nil {
		fmt.Println("Error downloading lineup JSON: ", err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	type Channel struct {
		GuideNumber string
		GuideName   string
		URL         string
	}

	var channels []Channel

	err = json.Unmarshal(body, &channels)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}

	row := "%3s\t%-20s\t%s\n"
	for _, ch := range channels {
		fmt.Printf(row, ch.GuideNumber, ch.GuideName, ch.URL)
	}
}

func scanChannels() {
	scan_url := "http://%s/lineup.post?scan=start&source=Cable"
	status_url := "http://%s/lineup.html\n"
	hdhr_ip := discoverHDHR()
	_, err := http.Get(fmt.Sprintf(scan_url, hdhr_ip))
	if err != nil {
		fmt.Println("Error issues rescan command: ", err)
		return
	}
	fmt.Println("Success! Go here to see status:")
	fmt.Printf(status_url, hdhr_ip)
}
