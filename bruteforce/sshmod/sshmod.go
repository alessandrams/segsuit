package sshmod

import (
  "fmt"
  "golang.org/x/crypto/ssh"
  "time"
  "sync"
)

func SshConfig(user string, pass string, timeout time.Duration) (*ssh.ClientConfig){
  sshConfig := &ssh.ClientConfig {
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
		Timeout: timeout*time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
  return sshConfig
}

func SshDial(sshconf *ssh.ClientConfig, ip string, channel chan int) (bool, *ssh.Client) {
  c, err := ssh.Dial("tcp", ip + ":22", sshconf)
	if err != nil {
		<-channel
    //fmt.Println(err)
		return false, nil
	}
  return true, c
}

func SshConnect(wg *sync.WaitGroup, channel chan int, user, pass, ip string) {
  defer wg.Done()
  fmt.Printf("Trying to connect with %s:%s...\n", user, pass)
  sshconf := SshConfig(user, pass, 5)

  success, c := SshDial(sshconf, ip, channel)
  if success == false {
	   return
  }
	session, err := c.NewSession()
	if err != nil {
    fmt.Println("[!] Failed to create SSH Session wit host\n", ip)
    <-channel
    return
  }
  if session != nil {
    fmt.Printf("[*] Got it! %s:%s\n", user, pass)
  }
  c.Close()
  session.Close()
	<-channel
}
