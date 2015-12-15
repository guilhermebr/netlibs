package dhcp

import "net"

// Packet is a DHCP Message defined in rfc2131
// Comments below are rfc2131 definitions
type Packet struct {
	opCode  byte      // Message op code / message type
	htype   byte      // Hardware address type
	hlen    byte      // Hardware address length
	hops    byte      // Hops used by relay agents
	xid     [4]byte   // Transaction ID
	secs    [2]byte   // Seconds elapsed
	flags   [2]byte   // Flags (Broadcast flag)
	ciaddr  [4]byte   // Client IP address
	yiaddr  [4]byte   // 'your' (client) IP address
	siaddr  [4]byte   // IP address of next server to use in bootstrap
	giaddr  [4]byte   // Relay agent IP address
	chaddr  [16]byte  // Client hardware address
	sname   [64]byte  // Optional server host name
	file    [128]byte // Boot file name
	options []byte    // Optional parameters field (variable size)
}

func ParsePacket(buffer []byte) *Packet {
	p := &Packet{}
	p.opCode = buffer[0]
	p.htype = buffer[1]
	p.hlen = buffer[2]
	p.hops = buffer[3]
	copy(p.xid[:], buffer[4:8])
	copy(p.secs[:], buffer[8:10])
	copy(p.flags[:], buffer[10:12])
	copy(p.ciaddr[:], buffer[12:16])
	copy(p.yiaddr[:], buffer[16:20])
	copy(p.siaddr[:], buffer[20:24])
	copy(p.giaddr[:], buffer[24:28])
	copy(p.chaddr[:], buffer[28:44])
	copy(p.sname[:], buffer[44:108])
	copy(p.file[:], buffer[108:236])
	p.options = buffer[236:]

	return p

}

//Packet Methods
func (p *Packet) Hlen() byte {
	return p.hlen
}

func (p *Packet) Xid() []byte {
	return p.xid[:]
}

func (p *Packet) Ciaddr() net.IP {
	return net.IP(p.ciaddr[0:4])
}

func (p *Packet) Yiaddr() net.IP {
	return net.IP(p.yiaddr[0:4])
}

func (p *Packet) Siaddr() net.IP {
	return net.IP(p.siaddr[0:4])
}

func (p *Packet) Giaddr() net.IP {
	return net.IP(p.giaddr[0:4])
}

func (p *Packet) Chaddr() net.HardwareAddr {
	hLen := p.Hlen()
	if hLen > 16 {
		hLen = 16
	}
	return net.HardwareAddr(p.chaddr[0:hLen])
}

func (p *Packet) Options() *Options {
	opt := &Options{}

	copy(opt.cookie[:], p.options[:4])

	return opt
}
