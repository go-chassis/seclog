package syslog

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)
//Writer is a struct
type Writer struct {
	conn net.Conn
}

var syslogHeader string
//New is a function
func New() (*Writer, error) {
	return Dial("", "", "", "")
}
//Dial is a function
func Dial(component, appguid, network, raddr string) (*Writer, error) {

	hostname, _ := os.Hostname()
	// construct syslog header the same to rsyslog's,
	// origin, node_id, app_guid, instance_id, loglevel
	syslogHeader = fmt.Sprintf("%s  %s  %s  %s  %s", component, component, appguid, hostname, "all")

	var conn net.Conn
	var err error
	if network == "" {
		conn, err = unixSyslog()
	} else {
		conn, err = net.Dial(network, raddr)
	}
	return &Writer{
		conn: conn,
	}, err
}
//Write is a function used to write
func (r *Writer) Write(b []byte) (int, error) {
	nl := ""
	if len(b) == 0 || b[len(b)-1] != '\n' {
		nl = "\n"
	}

	r.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))

	_, err := fmt.Fprintf(r.conn, "  %s  %s%s", syslogHeader, b, nl)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
//Close is a function used to close connection
func (r *Writer) Close() error {
	return r.conn.Close()
}
//unixSyslog is a function
func unixSyslog() (net.Conn, error) {
	networks := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog"}
	var addr string
	for _, n := range networks {
		for _, p := range logPaths {
			addr = p
			if conn, err := net.Dial(n, addr); err == nil {
				return conn, nil
			}
		}
	}
	return nil, errors.New("Could not connect to local syslog socket")
}
