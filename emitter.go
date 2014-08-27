package emitter

import "net"

// Emitter - An object that emits events to an event collector
type Emitter struct {
	conn *net.UDPConn
}

// NewEmitter - Creates and returns a new Emitter
func NewEmitter() *Emitter {

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6900")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	return &Emitter{conn}
}

// Emit - emits a packet to the local event collector
func (e *Emitter) Emit(packetType string, data interface{}) error {
	p := newPacket(packetType, data)
	b, err := p.ToJSON()
	if err != nil {
		return err
	}

	_, err = e.conn.Write(b)

	return err
}
