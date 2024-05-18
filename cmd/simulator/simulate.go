package main

import (
	"fmt"
	"github.com/ebauman/librmonitor"
	"github.com/spf13/cobra"
	"net"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "simulate an RMonitor tcp source",
	Short: "simulate",
	RunE:  run,
}

var (
	address    net.IP
	port       int
	sourceFile string
)

func init() {
	rootCmd.Flags().IPVar(&address, "address", net.ParseIP("127.0.0.1"), "address on which to listen")
	rootCmd.Flags().IntVar(&port, "port", 50000, "port on which to listen")
	rootCmd.Flags().StringVar(&sourceFile, "file", "", "file to read data from")
	_ = rootCmd.MarkFlagRequired("file")
}

func run(cmd *cobra.Command, args []string) error {
	if port > 65535 || port < 1024 {
		return fmt.Errorf("invalid port %d, must be 1024<port<65535", port)
	}

	f, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("error openign file: %s", err.Error())
	}
	defer f.Close()

	addr := net.TCPAddr{
		IP:   address,
		Port: port,
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	server, err := librmonitor.Simulate(&addr, f, stopCh)
	if err != nil {
		return fmt.Errorf("error simulating: %s", err.Error())
	}

	fmt.Println("simulating...")

	go server.Run()

	for {
		select {
		case n := <-server.ConnNotifs():
			fmt.Printf("client %s connected\n", n)
		case err := <-server.ConnErrors():
			fmt.Printf("client error: %s\n", err.Error())
		case <-cmd.Context().Done():
			return nil
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
