package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ftp struct {
	conn             net.Conn
	pasvListener     net.Listener
	currentDirectory string
	pasvMode         bool
	clientHost       string
	binary           bool
}

func (ftp *ftp) write(str string) {
	_, err := fmt.Fprintf(ftp.conn, "%s\r\n", str)
	if err != nil {
		log.Print(err)
	}
}

func (ftp *ftp) pasv() {
	var err error
	ftp.pasvListener, err = net.Listen("tcp", "")
	if err != nil {
		log.Print(err)
		return
	}

	_, portStr, err := net.SplitHostPort(ftp.pasvListener.Addr().String())
	if err != nil {
		log.Print(err)
		return
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Print(err)
		return
	}

	host, _, err := net.SplitHostPort(ftp.conn.LocalAddr().String())
	if err != nil {
		log.Print(err)
		return
	}
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		log.Print(err)
		return
	}
	ip := ipAddr.IP.To4()

	ftp.pasvMode = true

	msg := fmt.Sprintf("227 Entering Passive Mode (%d,%d,%d,%d,%d,%d).", ip[0], ip[1], ip[2], ip[3], port/256, port%256)
	ftp.write(msg)
}

func (ftp *ftp) list(str string) {
	ftp.write("150 Here comes the directory listing.")
	conn, err := ftp.dataConn()
	defer conn.Close()
	if err != nil {
		log.Print(err)
		return
	}
	output, err := exec.Command("ls", "-la", str).Output()
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Fprint(conn, strings.Replace(string(output), "\n", "\r\n", -1))
	ftp.write("226 Directory send OK.")
}

func (ftp *ftp) type_(str string) {
	switch str {
	case "I":
		ftp.binary = true
		ftp.write("200 Switching to Binary mode.")
	case "A":
		ftp.binary = false
		ftp.write("200 Switching to ASCII mode.")
	}
}

func (ftp *ftp) retr(str string) {
	ftp.write("150 Opening BINARY mode data connection for test.txt (0 bytes).")
	conn, err := ftp.dataConn()
	defer conn.Close()
	if err != nil {
		log.Print(err)
		return
	}
	file, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ftp.binary {
		io.Copy(conn, file)
	} else {
		r := bufio.NewReader(file)
		w := bufio.NewWriter(conn)
		for {
			line, isPrefix, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			w.Write(line)
			if !isPrefix {
				w.Write([]byte("\r\n"))
			}
		}
		w.Flush()
	}
	ftp.write("226 Transfer complete.")
}

func (ftp *ftp) stor(str string) {
	ftp.write("150 Ok to send data.")
	conn, err := ftp.dataConn()
	defer conn.Close()
	if err != nil {
		log.Print(err)
		return
	}
	file, err := os.Create(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(file, conn)
	ftp.write("226 Transfer complete.")
}

func (ftp *ftp) port(str string) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(str, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	ftp.clientHost = fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2)
	ftp.pasvMode = false

	ftp.write("200 PORT command successful. Consider using PASV.")
}

func (ftp *ftp) cwd(args []string) {
	var dir string
	if len(args) > 1 {
		dir = args[1]
	} else {
		dir = os.Getenv("HOME")
	}
	if err := os.Chdir(dir); err != nil {
		log.Print(err)
		return
	}
	ftp.currentDirectory = dir
	ftp.write("250 Directory successfully changed.")
}

func (ftp *ftp) pwd() {
	ftp.write(fmt.Sprintf("257 \"%s\"", ftp.currentDirectory))
}

func (ftp *ftp) dataConn() (conn net.Conn, err error) {
	if ftp.pasvMode {
		conn, err = ftp.pasvListener.Accept()
		if err != nil {
			return nil, err
		}
	} else {
		conn, err = net.Dial("tcp", ftp.clientHost)
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}

func (ftp *ftp) feat() {
	ftp.write("211-Features:")
	ftp.write(" EPRT")
	ftp.write(" EPSV")
	ftp.write("211 End")
}

func (ftp *ftp) epsv() {
	var err error
	ftp.pasvListener, err = net.Listen("tcp", "")
	if err != nil {
		log.Print(err)
		return
	}

	_, portStr, err := net.SplitHostPort(ftp.pasvListener.Addr().String())
	if err != nil {
		fmt.Println(err)
	}

	ftp.pasvMode = true

	msg := fmt.Sprintf("229 Entering Extended Passive Mode (|||%s|).", portStr)
	ftp.write(msg)
}

func (ftp *ftp) erpt(str string) {
	args := strings.Split(str, "|")
	ftp.clientHost = fmt.Sprintf("%s:%s", args[2], args[3])
	ftp.pasvMode = false
	ftp.write("200")
}

func main() {
	listener, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func newFtp(conn net.Conn) *ftp {
	return &ftp{
		conn:             conn,
		currentDirectory: os.Getenv("HOME"),
	}
}

func handleConn(conn net.Conn) {
	ftp := newFtp(conn)
	ftp.write("200 Welcome to FTP service.")

	input := bufio.NewScanner(ftp.conn)
	for input.Scan() {
		cmds := strings.Fields(input.Text())
		switch cmds[0] {
		case "USER":
			if err := os.Chdir(os.Getenv("HOME")); err != nil {
				log.Print(err)
				return
			}
			ftp.write("230 Login successful.")
		case "SYST":
			ftp.write("215 UNIX Type: L8")
		case "PASV":
			ftp.pasv()
		case "LIST":
			ftp.list(cmds[1])
		case "TYPE":
			ftp.type_(cmds[1])
		case "RETR":
			ftp.retr(cmds[1])
		case "STOR":
			ftp.stor(cmds[1])
		case "PORT":
			ftp.port(cmds[1])
		case "CWD":
			ftp.cwd(cmds)
		case "PWD":
			ftp.pwd()
		case "FEAT":
			ftp.feat()
		case "EPSV":
			ftp.epsv()
		case "EPRT":
			ftp.erpt(cmds[1])
		case "QUIT":
			ftp.write("221 Goodbye.")
			return
		}
		fmt.Println(cmds)
	}
}
