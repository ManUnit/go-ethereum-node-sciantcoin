package discv5

import "github.com/manunit/go-ethereum-node-sciantcoinmetrics"

var (
	ingressTrafficMeter = metrics.NewRegisteredMeter("discv5/InboundTraffic", nil)
	egressTrafficMeter  = metrics.NewRegisteredMeter("discv5/OutboundTraffic", nil)
)
