package libp2pquic

import (
	"net"

	ma "gx/ipfs/QmTZBfrPJmjWsCvHEtX5FE6KimVJhsJg5sBbqEFYf4UZtL/go-multiaddr"
	manet "gx/ipfs/Qmc85NSvmSG4Frn9Vb2cBc1rMyULH6D3TNVEfCzSKoUpip/go-multiaddr-net"
)

var quicMA ma.Multiaddr

func init() {
	var err error
	quicMA, err = ma.NewMultiaddr("/quic")
	if err != nil {
		panic(err)
	}
}

func toQuicMultiaddr(na net.Addr) (ma.Multiaddr, error) {
	udpMA, err := manet.FromNetAddr(na)
	if err != nil {
		return nil, err
	}
	return udpMA.Encapsulate(quicMA), nil
}

func fromQuicMultiaddr(addr ma.Multiaddr) (net.Addr, error) {
	return manet.ToNetAddr(addr.Decapsulate(quicMA))
}
