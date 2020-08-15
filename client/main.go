package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var host, capath, certpath, keypath string
	flag.StringVar(&capath, "ca", "", "path of ca file")
	flag.StringVar(&certpath, "cert", "", "path of cert file")
	flag.StringVar(&keypath, "key", "", "path of key file")
	flag.StringVar(&host, "host", "", "host of https server")
	flag.Parse()

	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(capath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair(certpath, keypath)
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(host)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body) + "\n")
}
