package models

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func runcmd(s SshServer) {
	session, err := connect(s.Username, s.Password, s.Ip, s.Port)
	if err != nil {
		cmdRet <- s.Ip + "--" + "args has error" + "--" + s.Cmd
		<-goroutineNum
		wg.Done()
	} else {
		defer session.Close()
		//直接输出到屏幕
		//session.Stdout = os.Stdout
		//session.Stderr = os.Stderr
		//session.Run(cmd)
		ret, err := session.Output(s.Cmd)
		if err != nil {
			cmdRet <- s.Ip + "--" + "Please check command" + "," + err.Error() + "--" + s.Cmd
		} else {
			cmdRet <- s.Ip + "--" + string(ret) + "--" + s.Cmd
		}
		<-goroutineNum
		wg.Done()
	}
}

var goroutineNum = make(chan int, 10)
var cmdRet = make(chan string, 1)
var wg sync.WaitGroup

type SshServer struct {
	Ip       string
	Username string
	Password string
	Cmd      string
	Port     int
}

type SshServers []SshServer

type CmdResult struct {
	Ip      string
	Command string
	Result  string
}

func Run(servers SshServers) []*CmdResult {
	aa := 0
	for i, server := range servers {
		goroutineNum <- i
		aa++
		wg.Add(1)
		go runcmd(server)
	}
	var result []*CmdResult
	for a := 0; a < aa; a++ {
		rr := <-cmdRet
		ip := strings.Split(rr, "--")[0]
		ret := strings.Split(rr, "--")[1]
		cmd := strings.Split(rr, "--")[2]
		result = append(result, &CmdResult{ip, cmd, ret})
	}
	wg.Wait()
	return result
}
