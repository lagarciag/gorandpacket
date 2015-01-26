/*Package gorandpacket is a go library for randomly generating ethernet packets
It's main purpose is for testing networking software/hardware
*/
package gorandpacket

import (
	"code.google.com/p/gopacket"
	"code.google.com/p/gopacket/layers"
	//"fmt"
	"math/rand"
	"net"
	"time"
)

//Struct RandPacket holds the gorandpacket object
type RandPacketT struct {
	MACLen int
	Seed   int64
	Rand   *rand.Rand
}

//NewGorandPacket Factory method for gorandpacket
func NewGorandPacket() RandPacketT {

	rp := RandPacketT{}
	rp.MACLen = 6
	rp.Seed = int64(time.Now().Nanosecond())
	rp.Rand = rand.New(rand.NewSource(rp.Seed))
	return rp

}

//RandIPv4Addr generates a random IPv4 address
func (r *RandPacketT) RandIPv4Addr() net.IP {
	var myIP net.IP
	myIP = make(net.IP, net.IPv4len)

	for i := 0; i < len(myIP); i++ {
		myIP[i] = r.RandByte()
	}
	return myIP
}

//RandIPv6Addr Generate a random IPv6 address
func (r *RandPacketT) RandIPv6Addr() net.IP {
	var myIP net.IP
	myIP = make(net.IP, net.IPv6len)

	for i := 0; i < len(myIP); i++ {
		myIP[i] = r.RandByte()
	}
	return myIP
}

//RandMACAddr Generate a random MAC address
func (r *RandPacketT) RandMACAddr() net.HardwareAddr {

	var myMAC net.HardwareAddr
	myMAC = make(net.HardwareAddr, r.MACLen)
	for i := 0; i < r.MACLen; i++ {
		myMAC[i] = r.RandByte()
	}
	return myMAC

}

//SetSeed Set a seed from an external source
func (r *RandPacketT) SetSeed(s int64) {

	r.Seed = s

}

//RandInt16 generates random uint16
func (r *RandPacketT) RandInt16() uint16 {
	num := r.Rand.Intn(0xFFFF)
	num16 := uint16(num)
	return num16
}

//RandInt32 generates random uint32
func (r *RandPacketT) RandInt32() uint32 {

	num := r.Rand.Intn(0xFFFFFFFF)
	num32 := uint32(num)
	return num32
}

//RandByte Generate a random byte
func (r *RandPacketT) RandByte() byte {

	num := r.Rand.Intn(255)
	mByte := uint8(num)

	return mByte
}

//RandInt generates a random int
func (r *RandPacketT) RandInt(n int) int {
	return r.Rand.Intn(n)
}

//RandPayload generates a Random Payload
func (r *RandPacketT) RandPayload() []byte {

	pSize := r.RandInt(1000)

	var myPayload []byte
	myPayload = make([]byte, pSize)

	for i := 5; i < int(pSize); i++ {
		myPayload[i] = r.RandByte()
		//println(myPayload[i])

	}
	return myPayload

}

//RandIPv4Layer generates a random IPv4 layer
func (r *RandPacketT) RandIPv4Layer() *layers.IPv4 {
	const (
		l3tcp = iota // c0 == 0
		l3udp = iota // c1 == 1
	)
	var l3protocol uint8
	//Randomly choose the l3 protocol to be used
	switch r.Rand.Intn(2) {
	case l3tcp:
		l3protocol = uint8(layers.IPProtocolTCP)
	case l3udp:
		l3protocol = uint8(layers.IPProtocolUDP)
	}

	ipv4 := layers.IPv4{
		Version:    uint8(4),
		IHL:        uint8(5),
		TOS:        uint8(0x1),
		Length:     uint16(40),
		Id:         uint16(r.RandInt16()),
		Flags:      layers.IPv4Flag(0),
		FragOffset: uint16(0),
		TTL:        uint8(0x1),
		Protocol:   layers.IPProtocol(l3protocol),
		Checksum:   uint16(0),
		SrcIP:      r.RandIPv4Addr(),
		DstIP:      r.RandIPv4Addr(),
	}
	return &ipv4
}

//RandIPUDP generates a random UDP layer
func (r *RandPacketT) RandIPUDPLayer() *layers.UDP {

	udp := layers.UDP{
		SrcPort:  layers.UDPPort(r.RandInt16()),
		DstPort:  layers.UDPPort(r.RandInt16()),
		Length:   8,
		Checksum: 0,
	}
	return &udp
}

//RandIPv4TCPLayer generates a random TCP layer
func (r *RandPacketT) RandIPTCPLayer() *layers.TCP {

	ipv4Tcp := layers.TCP{
		SrcPort:    layers.TCPPort(r.RandInt16()), //uint16
		DstPort:    layers.TCPPort(r.RandInt16()), //uint16
		Seq:        0xFFFFFFFF,                    //uint32
		Ack:        0x2,                           //uint32
		DataOffset: uint8(5),                      //must be 5                             //uint8, higer 4 bits are 0.
		FIN:        false,
		SYN:        false,
		RST:        false,
		PSH:        false,
		ACK:        false,
		URG:        false,
		ECE:        false,
		CWR:        false,
		NS:         false,
		Window:     0xff,
		Checksum:   uint16(0),
		Urgent:     0xFF, //uint16
	}
	return &ipv4Tcp

}

//RandEthernetLayer generates a random Ethernet layer
func (r *RandPacketT) RandEthernetLayer() *layers.Ethernet {
	eth := layers.Ethernet{}
	eth.EthernetType = layers.EthernetTypeIPv4
	eth.SrcMAC = r.RandMACAddr()
	eth.DstMAC = r.RandMACAddr()
	return &eth

}

//RandL3Layer generates a random L3 layer:  currently supports TCP & UDP only.
func (r *RandPacketT) RandL3Layer(l3type layers.IPProtocol) gopacket.SerializableLayer {

	var l3 gopacket.SerializableLayer

	/*********************
	Create a Random L3 layer
	**********************/
	switch l3type {
	case layers.IPProtocolTCP:
		// Generate a random TCP layer
		l3 = r.RandIPTCPLayer()
	case layers.IPProtocolUDP:
		//Generate a random UDP layer
		l3 = r.RandIPUDPLayer()
	default:
		panic("Bad l3 packet type")
	}
	return l3
}

/*RandEthernetPacket generates a random ethernet packet.
For now it only generates IPv4/TCP packets
TODO:  Generate more types/protocols
*/
func (r *RandPacketT) RandEthernetPacket() gopacket.SerializeBuffer {

	buf := gopacket.NewSerializeBuffer()

	// See gopacket SerializeOptions for more details.
	opts := gopacket.SerializeOptions{}

	// Generate a random ethernet layer
	eth := r.RandEthernetLayer()

	// Generate a random IPV4 Layer
	//TODO: randomize ip version
	l3 := r.RandIPv4Layer()

	//Generate a random IP L4 layer.
	l4 := r.RandL3Layer(l3.Protocol)

	/*****************************
	Check IP layer size
	******************************/
	l3Buf := gopacket.NewSerializeBuffer()
	l4Buf := gopacket.NewSerializeBuffer()
	err := l3.SerializeTo(l3Buf, opts)
	err = l4.SerializeTo(l4Buf, opts)

	l3.Length = uint16(len(l3Buf.Bytes())) + uint16(len(l4Buf.Bytes()))
	/*****************************
	Generate the final ethernet frame
	by serializing all generated layers
	*****************************/
	err = gopacket.SerializeLayers(buf, opts, eth, l3, l4)
	//fmt.Println(buf)
	if err != nil {
		panic(err)
	}
	return buf
}
