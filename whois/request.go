package whois

import (
	"io/ioutil"
	"net"
	"strings"
	"time"
)

type Request struct {
	Domain  string
	Port    string
	Server  string
	Timeout time.Duration
}

func (r Request) Query(domain string) (string, error) {
	connection, err := r.connection(r.Timeout, r.Server, r.Port)

	if err != nil {
		return "", err
	}

	defer connection.Close()

	connection.Write([]byte(domain + "\r\n"))
	buffer, err := ioutil.ReadAll(connection)
	if err != nil {
		return "", err
	}

	return string(buffer[:]), nil
}

func (r Request) connection(timeout time.Duration, server string, port string) (net.Conn, error) {
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, port), timeout)
	if err != nil {
		return connection, err
	}
	return connection, err
}

func NewRequest(domain string, server string, port string, timeout time.Duration) Request {
	return Request{
		Server:  server,
		Port:    port,
		Timeout: timeout,
	}
}

func getTldFromDomain(domain string) string {
	parts := strings.Split(domain, ".")
	tld := parts[len(parts)-1]
	return tld
}
