package protocols

import (
	"github.com/johnthethird/thresher/network"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)


func handlerLoop(id party.ID, h protocol.Handler, network network.Network) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen(): // listen command gets the messages
			if !ok {
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			// message is sent over the network
			go network.Send(msg)

		// incoming messages
		case msg := <-network.Next(id):
			// parsed by the protocol handler
			h.Accept(msg)
		}
	}
}
