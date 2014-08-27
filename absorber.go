package emitter

import (
	"encoding/json"
	"net"
)

// Absorber - Object that "absorbs" as opposed to emits
type Absorber struct {
	conn *net.PacketConn
}

// NewAbsorber - Returns a new Absorber object to play with
func NewAbsorber() (*Absorber, error) {
	conn, err := net.ListenPacket("udp", "127.0.0.1:6900")
	if err != nil {
		return nil, err
	}

	return &Absorber{&conn}, nil
}

func (a *Absorber) waitForBytes() ([]byte, error) {
	buf := make([]byte, 5000)
	size, _, err := (*a.conn).ReadFrom(buf)
	if err != nil {
		return nil, err
	}
	return buf[:size], nil
}

// WaitForPacket - Blocks waiting for an emitted packet
func (a *Absorber) WaitForPacket() (*packet, error) {
	buf, err := a.waitForBytes()
	if err != nil {
		return nil, err
	}

	p := packet{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
