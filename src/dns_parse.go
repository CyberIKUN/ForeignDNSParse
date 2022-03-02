package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type DNSServer struct {
	DNSAddress string `yaml:"dns_address"`
	Domain     string `yaml:"domain"`
}

func (d *DNSServer) readFromConfig() {
	path, err := filepath.Abs(`config/address.yaml`)
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(file).Decode(d)
	if err != nil {
		panic(err)
	}

}
