package main

import (
	"bytes"
	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"sync"
	_ "time"
)

func getStatus(hosts []string) {
	wait := new(sync.WaitGroup)
	for _, host := range hosts {
		wait.Add(1)
		go func(host string) {
			user := ssh_config.Get(host, "User")
			addr := ssh_config.Get(host, "Hostname") + ":" + ssh_config.Get(host, "Port")
			auth := []ssh.AuthMethod{}
			buf, err := ioutil.ReadFile(ssh_config.Get(host, "IdentityFile"))
			if err != nil {
				log.Println(err)
			}
			key, err := ssh.ParsePrivateKey(buf)
			if err != nil {
				log.Println(err)
			}
			auth = append(auth, ssh.PublicKeys(key))
			config := &ssh.ClientConfig{
				User:            user,
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				Auth:            auth,
			}
			conn, err := ssh.Dial("tcp", addr, config)
			if err != nil {
				log.Println(err)
			}
			defer conn.Close()
			//
			session, err := conn.NewSession()
			if err != nil {
				log.Println(err)
			}
			defer session.Close()
			//Check whoami
			var b bytes.Buffer
			session.Stdout = &b
			remote_command := "w"
			if err := session.Run(remote_command); err != nil {
				log.Fatal("Failed to run: " + err.Error())
			}
			log.Println(host + ":" + remote_command + ":" + b.String())
			defer wait.Done()
		}(host)
	}
	wait.Wait()
}

func main() {
	hosts := []string{
		"****",
		"****",
	}
	getStatus(hosts)
}
