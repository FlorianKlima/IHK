package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

type userData struct {
	name     string
	password string
}

type serverData struct {
	IP  string
	ssh int
}

type wapData struct {
	MOTORBEZEICHNUNG  string
	GESCHAEFTSEINHEIT string
	GUELTIGAB         string
	GUETLTIGBIS       string
	ZUSCHLAGART       string
	ZUSCHLAG_WERT     string
	ANGELEGT_VON      string
	ANGELEGT_DATE     string
	UPDATE_VON        string
	UPDATE_DATE       string
	BEMERKUNG         string
	STATUS            bool
}

func main() {
	// Read config file
	userdata, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var user userData

	// Unmarshalling config.json
	err = json.Unmarshal(userdata, &user)
	if err != nil {
		log.Fatal(err)
	}

	// Read config.json file
	serverdata, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var server serverData

	// Unmarshalling config.json
	err = json.Unmarshal(serverdata, &server)
	if err != nil {
		log.Fatal(err)
	}

	// Control Output
	fmt.Printf("User Name: %s\n", user.name)
	fmt.Printf("User Password: %s\n", user.password)
	fmt.Printf("Server IP: %s\n", server.IP)
	fmt.Printf("Server SSH Port: %d\n", server.ssh)

	// Open CSV file
	f, err := os.Open("./test.csv")
	if err != nil {
		log.Fatal(err)
	}
	// CLose File later
	defer f.Close()

	// Read CSV Values
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Datentestausgabe
		fmt.Printf("%+v\n", rec)
	}

	clientConfig, _ := auth.PrivateKey("user.name", "./rsakey", ssh.InsecureIgnoreHostKey())

	// Create a new SCP client
	client := scp.NewClient("server.IP:server.port", &clientConfig)

	// Connect to the remote server
	conErr := client.Connect()
	if conErr != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}
	// Open a file
	tf, _ := os.Open("./test.csv")

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer tf.Close()

	// Copy the file over
	err = client.CopyFromFile(context.Background(), *tf, "./test.csv", "0655")
	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}
