package main

import (
	context "context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/speedrunsh/speedrun/key"
	"github.com/speedrunsh/speedrun/marathon"

	qnet "github.com/speedrunsh/grpc-quic"
	grpc "google.golang.org/grpc"
)

const Port int = 1337

func PortalTunnelConnection(address string, key *key.Key) (*grpc.ClientConn, error) {
	sshclient, err := marathon.Connect(address, key, true)
	if err != nil {
		return nil, err
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return sshclient.Dial("tcp", "127.0.0.1:1337")
	}

	return grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
}

func PortalConnection(address string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%d", address, Port)
	return grpc.Dial(target, grpc.WithInsecure())
}

func PortalQuicConnection(address string) (*grpc.ClientConn, error) {
	tlsConf := &tls.Config{
		MinVersion:         tls.VersionTLS13,
		CurvePreferences:   []tls.CurveID{tls.X25519},
		CipherSuites:       []uint16{tls.TLS_CHACHA20_POLY1305_SHA256},
		InsecureSkipVerify: true,
		NextProtos:         []string{"speedrun"},
	}

	creds := qnet.NewCredentials(tlsConf)

	dialer := qnet.NewQuicDialer(tlsConf)
	grpcOpts := []grpc.DialOption{
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(creds),
	}

	target := fmt.Sprintf("%s:%d", address, Port)
	return grpc.Dial(target, grpcOpts...)
}

func PortalTunnelQuicConnection(address string) (*grpc.ClientConn, error) {
	return nil, fmt.Errorf("not implemented yet")
}

// tlsConf := &tls.Config{
// 	InsecureSkipVerify: true,
// 	NextProtos:         []string{"speedrun-portal-v1"},
// }
// session, err := quic.DialAddr(instance.Address+":1337", tlsConf, &quic.Config{
// 	MaxIdleTimeout:       time.Second * 10,
// 	HandshakeIdleTimeout: time.Second * 5,
// })
// if err != nil {
// 	return fmt.Errorf("%s: %v", instance.Name, err)
// }

// stream, err := session.OpenStreamSync(context.Background())
// if err != nil {
// 	return err
// }

// fmt.Printf("Client: Sending '%s'\n", "dosomething")
// _, err = stream.Write([]byte("dosomething"))
// if err != nil {
// 	return err
// }

// buf := make([]byte, len("dosomething"))
// _, err = io.ReadFull(stream, buf)
// if err != nil {
// 	return err
// }
// fmt.Printf("Client: Got '%s'\n", buf)
