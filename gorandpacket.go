/*Package gorandpacket is a go library for randomly generating ethernet packets
It's main purpose is for testing networking software/hardware 
*/
package gorandpacket

import (
	"code.google.com/p/gopacket/layers"
	"math/rand"
	"net"
	"time"
	"code.google.com/p/gopacket"
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
		println(myPayload[i])

	}
	return myPayload

}
//RandIPv4Layer generates a random IPv4 layer
func (r *RandPacketT) RandIPv4Layer() *layers.IPv4 {

	ipv4 := layers.IPv4{
		Version:    uint8(4),
		IHL:        uint8(5),
		TOS:        uint8(0x1),
		Length:     uint16(40),
		Id:         uint16(0xFFFF),
		Flags:      layers.IPv4Flag(0),
		FragOffset: uint16(0),
		TTL:        uint8(0x1),
		Protocol:   layers.IPProtocolTCP,
		Checksum:   uint16(0),
		SrcIP:      r.RandIPv4Addr(),
		DstIP:      r.RandIPv4Addr(),
	}
	return &ipv4
}

//RandIPv4TCPLayer generates a random TCP layer
func (r *RandPacketT) RandIPv4TCPLayer() *layers.TCP {

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
	ipv4 :=  r.RandIPv4Layer()
	
	// Generate a random TCP layer
	ipv4Tcp := r.RandIPv4TCPLayer()
	
	//Serilize layers.
	err := gopacket.SerializeLayers(buf, opts,eth,ipv4,ipv4Tcp,
		gopacket.Payload(r.RandPayload()),
	) 
	
	if err != nil {
		panic(err)
	}
	
	return buf
	
}
