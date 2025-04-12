// Copyright (c) 2025 <Alptekin Sünnetci - https://alptekin.sunnetci.net>

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var debug bool

func main() {
	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")
	server := flag.String("s", "", "Server address (host:port)")
	localPath := flag.String("l", "", "Local directory")
	remotePath := flag.String("r", "", "Remote directory")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")

	flag.Parse()

	if *username == "" || *password == "" || *server == "" || *localPath == "" || *remotePath == "" {
		fmt.Println("Usage: winSftpClient.exe -u <username> -p <password> -s <host:port> -l <localFolder> -r <remoteFolder> [-debug]")
		return
	}

	if debug {
		log.Printf("🔍 Establishing connection: %s@%s\n", *username, *server)
	}

	config := &ssh.ClientConfig{
		User: *username,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", *server, config)
	if err != nil {
		log.Fatalf("❌ SSH connection error: %v", err)
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("❌ Failed to initialize SFTP client: %v", err)
	}
	defer client.Close()

	if debug {
		log.Println("✅ Connection successful, starting file upload...")
	}

	err = uploadDirectory(client, *localPath, *remotePath)
	if err != nil {
		log.Fatalf("❌ Files could not be uploaded: %v", err)
	}

	log.Println("✅ All files uploaded successfully.")
}

func uploadDirectory(client *sftp.Client, localPath, remotePath string) error {
	return filepath.Walk(localPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, localPath)
		relPath = strings.TrimPrefix(relPath, string(os.PathSeparator))
		targetPath := filepath.ToSlash(filepath.Join(remotePath, relPath))

		if info.IsDir() {
			if debug {
				log.Printf("📁 Creating remote directory: %s\n", targetPath)
			}
			return client.MkdirAll(targetPath)
		} else {
			return uploadFile(client, path, targetPath)
		}
	})
}

func uploadFile(client *sftp.Client, localFile, remoteFile string) error {
	if debug {
		log.Printf("⬆️  Uploading file: %s → %s\n", localFile, remoteFile)
	}

	dir := filepath.ToSlash(filepath.Dir(remoteFile))
	if err := client.MkdirAll(dir); err != nil {
		log.Printf("❌ Failed to create target directory: %v\n", err)
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	src, err := os.Open(localFile)
	if err != nil {
		log.Printf("❌ Failed to open local file: %v\n", err)
		return fmt.Errorf("failed to open local file: %w", err)
	}
	defer src.Close()

	dst, err := client.Create(remoteFile)
	if err != nil {
		log.Printf("❌ Failed to create remote file: %v\n", err)
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("❌ Failed to copy file: %v\n", err)
		return fmt.Errorf("failed to copy file: %w", err)
	}

	if debug {
		log.Printf("✅ Uploaded: %s\n", remoteFile)
	}

	return nil
}
