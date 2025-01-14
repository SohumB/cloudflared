//go:build linux

package ingress

import (
	"testing"

	"github.com/sohumb/cloudflared/packet"
)

func getFunnel(t *testing.T, proxy *icmpProxy, tuple flow3Tuple) (packet.Funnel, bool) {
	return proxy.srcFunnelTracker.Get(tuple)
}
