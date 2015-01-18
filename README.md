# gorandpacket
This library provides capabilities for random network frames in the Go programming language. It aims to create an extensible, easy-to-use, easy-to-expand library that's still relatively fast.

##PACKAGE DOCUMENTATION

###package gorandpacket
    import "github.com/lagarciag/gorandpacket"

    Gorandpacket is a go library for randomly generating ethernet packets
    It's main purpose if for testing networking software/hardware

TYPES

type RandPacketT struct {
    MACLen int
    Seed   int64
    Rand   *rand.Rand
}
    This struct holds the gorandpacket object

* func NewGorandPacket() RandPacketT
    Factory method for gorandpacket

* func (r *RandPacketT) RandByte() byte
    Generate a random byte

* func (r *RandPacketT) RandEthernetLayer() *layers.Ethernet
    Generate a random Ethernet layer

* func (r *RandPacketT) RandEthernetPacket() gopacket.SerializeBuffer
    Generate a random ethernet packet. For now it only generates IPv2/TCP
    packets TODO: Generate more types/protocols

func (r *RandPacketT) RandIPv4Addr() net.IP
    Generate a random IPv4 address

func (r *RandPacketT) RandIPv4Layer() *layers.IPv4
    Generate a random IPv4 layer

func (r *RandPacketT) RandIPv4TCPLayer() *layers.TCP
    Generate a random TCP layer

func (r *RandPacketT) RandIPv6Addr() net.IP
    Generate a random IPv6 address

func (r *RandPacketT) RandInt(n int) int

func (r *RandPacketT) RandInt16() uint16
    Generate random uint16

func (r *RandPacketT) RandInt32() uint32
    Generate random uint32

func (r *RandPacketT) RandMACAddr() net.HardwareAddr
    Generate a random MAC address

func (r *RandPacketT) RandPayload() []byte
    Generate Random Payload

func (r *RandPacketT) SetSeed(s int64)
    Set a seed from an external source


