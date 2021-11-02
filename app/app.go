package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the command line app ready to be executed.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line App"
	app.Usage = "Fetches IPs and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "gnu.org",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Fetches IPs of an address on the internet",
			Flags:  flags,
			Action: fetchIps,
		},
		{
			Name:   "server",
			Usage:  "Fetches server names of an address on the internet",
			Flags:  flags,
			Action: fetchServers,
		},
	}

	return app
}

func fetchIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func fetchServers(c *cli.Context) {

	host := c.String("host")
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
