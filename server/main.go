package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!\n")
}

func main() {
	var capath, certpath, keypath, bind string
	flag.StringVar(&capath, "ca", "", "path of ca file")
	flag.StringVar(&certpath, "cert", "", "path of cert file")
	flag.StringVar(&keypath, "key", "", "path of key file")
	flag.StringVar(&bind, "bind", "", "local address to bind")
	flag.Parse()

	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(capath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	//初始化一个server 实例。
	s := &http.Server{
		Addr:    bind,
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(certpath, keypath)

	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
