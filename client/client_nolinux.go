//go:build !(linux && amd64)
// +build !linux !amd64

package client

import "github.com/NHYCRaymond/trpc-go/transport"

func attemptSwitchingTransport(o *Options) transport.ClientTransport {
	if o.Transport == nil {
		return transport.DefaultClientTransport
	}
	return o.Transport
}
