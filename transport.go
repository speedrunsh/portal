package portal

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	qnet "github.com/speedrunsh/grpc-quic"

	"github.com/speedrunsh/speedrun/key"
	"github.com/speedrunsh/speedrun/ssh"
	"google.golang.org/grpc"
)

type Transport struct {
	Conn    *grpc.ClientConn
	Address string
	Key     key.Key
	opts    options
}

type options struct {
	insecure  bool
	useTunnel bool
}

type TransportOption interface {
	apply(*options)
}

func defaultOptions() options {
	return options{
		insecure:  false,
		useTunnel: true,
	}
}

type withInsecure bool

func (w withInsecure) apply(o *options) {
	o.insecure = bool(w)
}

func WithInsecure(enable bool) TransportOption {
	return withInsecure(enable)
}

type withTunnel bool

func (w withTunnel) apply(o *options) {
	o.useTunnel = bool(w)
}

func WithTunnel(enable bool) TransportOption {
	return withTunnel(enable)
}

func NewTransport(address string, key key.Key, opts ...TransportOption) (*Transport, error) {
	t := &Transport{
		Address: address,
		Key:     key,
		opts:    defaultOptions(),
	}
	for _, opt := range opts {
		opt.apply(&t.opts)
	}

	if t.opts.useTunnel {
		sshclient, err := ssh.Connect(address, &key)
		if err != nil {
			return nil, err
		}

		dialer := func(ctx context.Context, addr string) (net.Conn, error) {
			return sshclient.Dial("tcp", "127.0.0.1:1337")
		}

		t.Conn, err = grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func SSHTransport(address string, key *key.Key) (*grpc.ClientConn, error) {
	sshclient, err := ssh.Connect(address, key)
	if err != nil {
		return nil, err
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return sshclient.Dial("tcp", "127.0.0.1:1337")
	}

	return grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
}

func SSHTransportInsecure(address string, key *key.Key) (*grpc.ClientConn, error) {
	sshclient, err := ssh.ConnectInsecure(address, key)
	if err != nil {
		return nil, err
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return sshclient.Dial("tcp", "127.0.0.1:1337")
	}

	return grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
}

func HTTP2Transport(address string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%d", address, 1337)
	return grpc.Dial(target, grpc.WithInsecure())
}

func QUICTransport(address string) (*grpc.ClientConn, error) {
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

	target := fmt.Sprintf("%s:%d", address, 1337)
	return grpc.Dial(target, grpcOpts...)
}
