package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the ready-to-run command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search for IPs and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search for server IPs on the internet",
			Flags: flags,
			Action: searchIPs,
		},
		{
			Name: "servers",
			Usage: "Search for server IPs on the internet",
			Flags: flags,
			Action: searchServers,	
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
