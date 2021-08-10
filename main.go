package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/speedrunsh/portal-api/go/service"

	"github.com/coreos/go-systemd/v22/dbus"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:1337"

type server struct {
	service.UnimplementedPortalServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Echo(ctx context.Context, in *service.Empty) (*service.Empty, error) {
	log.Printf("Received ping")
	return &service.Empty{}, nil
}

func (s *server) ServiceRestart(ctx context.Context, in *service.Service) (*service.Response, error) {
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	responseChan := make(chan string, 1)
	serviceName := fmt.Sprintf("%s.service", in.GetName())
	_, err = conn.RestartUnitContext(ctx, serviceName, "replace", responseChan)
	if err != nil {
		return nil, err
	}

	res := <-responseChan
	return &service.Response{Content: res}, nil

}

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
func main() {
	// log.Fatal(echoServer())
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterPortalServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Start a server that echos all data on the first stream opened by the client
// func echoServer() error {
// 	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	sess, err := listener.Accept(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	stream, err := sess.AcceptStream(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Echo through the loggingWriter
// 	_, err = io.Copy(loggingWriter{stream}, stream)
// 	return err
// }

// // A wrapper for io.Writer that also logs the message.
// type loggingWriter struct{ io.Writer }

// func (w loggingWriter) Write(b []byte) (int, error) {
// 	fmt.Printf("Server: Got '%s'\n", string(b))
// 	return w.Writer.Write(b)
// }

// // Setup a bare-bones TLS config for the server
// func generateTLSConfig() *tls.Config {
// 	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
// 	if err != nil {
// 		panic(err)
// 	}

// 	template := x509.Certificate{SerialNumber: big.NewInt(1)}
// 	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey, privateKey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: bytes})
// 	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

// 	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &tls.Config{
// 		Certificates: []tls.Certificate{tlsCert},
// 		NextProtos:   []string{"speedrun-portal-v1"},
// 	}
// }
