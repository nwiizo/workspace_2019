package main

import (
	"bytes"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {

	ip := "***.***.***.***"
	port := "22"
	user := "root"
	pass := "*****************"
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	//Check whoami
	var b bytes.Buffer
	session.Stdout = &b
	remote_command := "/usr/bin/whoami"
	if err := session.Run(remote_command); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	log.Println(remote_command + ":" + b.String())
}
