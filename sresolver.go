package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
)

var (
	workers    = flag.Int("workers", 10, "number of worker goroutines")
	inputFile  = flag.String("input", "", "input file (default: stdin)")
	outputFile = flag.String("output", "", "output file (default: stdout)")
)

func main() {
	flag.Parse()

	var input *bufio.Scanner
	if *inputFile == "" {
		input = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = bufio.NewScanner(file)
	}

	var output *bufio.Writer
	if *outputFile == "" {
		output = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = bufio.NewWriter(file)
	}
	defer output.Flush()

	work := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup
	wg.Add(*workers)
	for i := 0; i < *workers; i++ {
		go func() {
			defer wg.Done()
			for address := range work {
				var ip string
				var port string

				u, err := url.Parse(address)
				if err != nil {
					ip, port, _ = net.SplitHostPort(address)
				} else {
					ip = u.Hostname()
					if u.Port() != "" {
						port = u.Port()
					} else if u.Scheme == "https" {
						port = "443"
					} else {
						port = "80"
					}
				}

				hostnames, err := net.LookupAddr(ip)
				if err != nil {
					// Ignore the IPs that don't resolve
					continue
				}

				for _, hostname := range hostnames {
					if port != "" {
						hostname = fmt.Sprintf("%s:%s", hostname, port)
					}
					results <- fmt.Sprintf("%s", hostname)
				}
			}
		}()
	}

	go func() {
		for input.Scan() {
			work <- strings.TrimSpace(input.Text())
		}
		close(work)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Fprintln(output, result)
	}
}
