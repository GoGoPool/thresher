package protocols

import (
	"log"

	"github.com/fxamacker/cbor/v2"
	"github.com/johnthethird/thresher/network"
	"github.com/johnthethird/thresher/wallet/avmwallet"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp/config"
)

func RunKeygen(w *avmwallet.Wallet, net network.Network) error {
	selfid := w.Me.PartyID() // gets the user party id (32 byte slice)
	allids := w.AllPartyIDs() // gets all party ids (32 byte slices)
	threshold := w.Threshold
	log.Printf("Starting Keygen protocol - selfid: %v, allids: %v threshold: %v", selfid, allids, threshold)

	pl := pool.NewPool(0)
	defer pl.TearDown()

	// creates a new handler the user can use to interact with the multi-party keygen protocol
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, selfid, allids, threshold, pl), nil)
	if err != nil {
		return err
	}

	// checks for messages from the handler and sends them to the network
	// as well as checking for incoming messages from the network
	handlerLoop(selfid, h, net)

	// returns the result if successful
	// if not, returns an error
	r, err := h.Result()
	if err != nil {
		return err
	}
	
	log.Print("Keygen protocol complete")

	c := r.(*config.Config)
	cb, err := cbor.Marshal(c)
	if err != nil {
		return err
	}	
	
	w.Initialize(cb)

	return nil
}

