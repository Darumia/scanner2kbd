package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/pelletier/go-toml/v2"
	"github.com/tarm/serial"
)

// TODO add configuration for AcceptUntil (byte) and AcceptKey (string) "enter"

type Config struct {
	Serial struct {
		Port string `toml:"port"`
		Baud int    `toml:"baud"`
	} `toml:"serial"`
}

func main() {
	config, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	serialConfig := &serial.Config{
		Name: config.Serial.Port,
		Baud: config.Serial.Baud,
	}

	port, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	reader := bufio.NewReader(port)
	for {
		line, err := reader.ReadString('\r')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}

		//robotgo.Type(line)
		i := strings.Split(line, "")
		for _, s := range i {
			robotgo.KeyTap(s)
		}
		//robotgo.Type("\r")
		robotgo.KeyTap("enter")
	}
}

func loadConfig(filename string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
