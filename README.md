# gorandpacket
--
    import "github.com/lagarciag/gorandpacket"

Gorandpacket is a go library for randomly generating ethernet packets It's main
purpose if for testing networking software/hardware

## Usage

#### type RandPacketT

```go
type RandPacketT struct {
	MACLen int
	Seed   int64
	Rand   *rand.Rand
}
```

RandPacket struct holds the gorandpacket object

#### func  NewGorandPacket

```go
func NewGorandPacket() RandPacketT
```
NewGorandPacket Factory method for gorandpacket

#### func (*RandPacketT) RandByte

```go
func (r *RandPacketT) RandByte() byte
```
RandByte Generate a random byte

#### func (*RandPacketT) RandEthernetLayer

```go
func (r *RandPacketT) RandEthernetLayer() *layers.Ethernet
```
RandEthernetLayer generates a random Ethernet layer

#### func (*RandPacketT) RandEthernetPacket

```go
func (r *RandPacketT) RandEthernetPacket() gopacket.SerializeBuffer
```
RandEthernetPacket generates a random ethernet packet. For now it only generates
IPv2/TCP packets TODO: Generate more types/protocols

#### func (*RandPacketT) RandIPv4Addr

```go
func (r *RandPacketT) RandIPv4Addr() net.IP
```
RandIPv4Addr generates a random IPv4 address

#### func (*RandPacketT) RandIPv4Layer

```go
func (r *RandPacketT) RandIPv4Layer() *layers.IPv4
```
RandIPv4Layer generates a random IPv4 layer

#### func (*RandPacketT) RandIPv4TCPLayer

```go
func (r *RandPacketT) RandIPv4TCPLayer() *layers.TCP
```
RandIPv4TCPLayer generates a random TCP layer

#### func (*RandPacketT) RandIPv6Addr

```go
func (r *RandPacketT) RandIPv6Addr() net.IP
```
RandIPv6Addr Generate a random IPv6 address

#### func (*RandPacketT) RandInt

```go
func (r *RandPacketT) RandInt(n int) int
```
RandInt generates a random int

#### func (*RandPacketT) RandInt16

```go
func (r *RandPacketT) RandInt16() uint16
```
RandInt16 generates random uint16

#### func (*RandPacketT) RandInt32

```go
func (r *RandPacketT) RandInt32() uint32
```
RandInt32 generates random uint32

#### func (*RandPacketT) RandMACAddr

```go
func (r *RandPacketT) RandMACAddr() net.HardwareAddr
```
RandMACAddr Generate a random MAC address

#### func (*RandPacketT) RandPayload

```go
func (r *RandPacketT) RandPayload() []byte
```
RandPayload generates a Random Payload

#### func (*RandPacketT) SetSeed

```go
func (r *RandPacketT) SetSeed(s int64)
```
SetSeed Set a seed from an external source
