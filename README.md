Sresolver

sresolver is a command-line tool that resolves IP addresses to hostnames. It takes a list of IP addresses (in the format of IP:port or http(s)://IP) as input and outputs the corresponding hostnames.


To install the latest version of sresolver from GitHub using go, you can use the following command:


go install github.com/s1n6h/sresolver@latest


This command will download the latest version of the sresolver tool from GitHub and install it to your $GOPATH/bin directory. You can then run sresolver from the command line.

Usage

sresolver [options]



Options
Option	                   Description
-input <file>	            input file (default: stdin)
-output <file>          	output file (default: stdout)
-workers <n>	            number of worker goroutines (default: 10)


Examples

To resolve a list of IP addresses from a file and output the hostnames to another file:
$ sresolver -input ips.txt -output hostnames.txt


To resolve a list of IP addresses from standard input and output the hostnames to standard output:
$ cat ips.txt | sresolver


Building
To build the sresolver tool, run the following command:

$ go build

This will create a binary named sresolver in the current directory.

Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/yourusername/sresolver. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the Contributor Covenant code of conduct.

License
The sresolver tool is available as open source under the terms of the MIT License.
