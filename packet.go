package emitter

import (
	"encoding/json"
	"time"
)

type packet struct {
	PacketType string      `json:"type"`
	Time       time.Time   `json:"time"`
	Data       interface{} `json:"data"`
}

func newPacket(packetType string, data interface{}) *packet {
	return &packet{
		packetType,
		time.Now(),
		data,
	}
}

func (p *packet) ToJSON() (b []byte, err error) {
	b, err = json.Marshal(&p)
	return
}

func (p *packet) AsType(iface interface{}) error {
	b, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, iface)
	return err
}
