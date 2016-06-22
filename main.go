package main

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"

	"fmt"
	"net/http"
	"github.com/gengo/grpc-gateway/runtime"
	"net"
	"strings"
	"crypto/tls"
	"crypto/x509"
	"log"
	"google.golang.org/grpc/credentials"
	"github.com/andreasmaier/cimon_jobs/jobs"
	"github.com/philips/grpc-gateway-example/insecure"
	"github.com/andreasmaier/cimon_jobs/handlers"
)

const (
	port = 10000
)

var (
	demoKeyPair *tls.Certificate
	demoCertPool *x509.CertPool
	demoAddr string
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			fmt.Println("Handling GRPC")
			grpcServer.ServeHTTP(w, r)
		} else {

			if(r.Method == "OPTIONS") {
				fmt.Printf("Handling HTTP with method %s\n", r.Method)
				w.WriteHeader(http.StatusOK)
			} else {
				fmt.Println("Handling HTTP")
				otherHandler.ServeHTTP(w, r)
			}
		}
	})
}

func init() {
	pair, err := tls.X509KeyPair([]byte(insecure.Cert), []byte(insecure.Key))
	if err != nil {
		panic(err)
	}
	demoKeyPair = &pair
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(insecure.Cert))
	if !ok {
		panic("bad certs")
	}
	demoAddr = fmt.Sprintf("localhost:10000")
}

func main() {
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:50051"))}

	grpcServer := grpc.NewServer(opts...)
	jobs.RegisterJobsApiServer(grpcServer, new(handlers.JobsServer))
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs: demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()

	gwmux := runtime.NewServeMux()
	if err := jobs.RegisterJobsApiHandlerFromEndpoint(ctx, gwmux, demoAddr, dopts); err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr: demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
		},
	}

	fmt.Printf("grpc on port: %d\n", port)
	if err = srv.Serve(tls.NewListener(conn, srv.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return
}