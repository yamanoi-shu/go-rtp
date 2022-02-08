package packet

import (
	"encoding/binary"
	"fmt"
)

type RTPHeader struct {
	Version     uint8
	Padding     bool
	Extention   bool
	CsrcCnt     uint8
	Marker      bool
	PayloadType uint8
	SequenceNum uint16
	Timestamp   uint32
	SSRC        uint32
}

type RTPPacket struct {
	Header  RTPHeader
	Payload []byte
}

func NewRTPPacket() *RTPPacket {
	return &RTPPacket{
		Header: RTPHeader{
			Version:   2,
			Padding:   true,
			Extention: false,
			Marker:    false,
		},
	}
}

func (h *RTPHeader) Marshal() (buf []byte) {
	/*
		The RTP header has the following format:

		 0                   1                   2                   3
		 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|V=2|P|X|  CC   |M|     PT      |       sequence number         |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                           timestamp                           |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|           synchronization source (SSRC) identifier            |
		+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		|            contributing source (CSRC) identifiers             |
		|                             ....                              |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	*/

	// 最初の8ビット
	// version(bit: 0 ~ 1)

	return
}

func (h *RTPHeader) Unmarshal(buf []byte) {
	/*
		The RTP header has the following format:

		 0                   1                   2                   3
		 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|V=2|P|X|  CC   |M|     PT      |       sequence number         |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                           timestamp                           |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|           synchronization source (SSRC) identifier            |
		+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		|            contributing source (CSRC) identifiers             |
		|                             ....                              |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	*/

	// (octed: 1)
	// version(bit: 0 ~ 1)
	h.Version = buf[0] >> 6

	// padding(bit: 2)
	h.Padding = buf[0]>>5&1 == 1

	// extention(bit: 3)
	h.Extention = buf[0]>>4&1 == 1

	// (bit: 4 ~ 7)
	h.CsrcCnt = buf[0] & 0xF

	// (octed: 2)
	// marker(bit: 0)
	h.Marker = buf[1]>>7 == 1

	// payload type(bit: 1 ~ 7)
	h.PayloadType = buf[1] & 0x7f

	// sequence number(octet: 3 ~ 4)
	h.SequenceNum = binary.BigEndian.Uint16(buf[2:4])

	// timestamp (octet: 4 ~ 7)
	h.Timestamp = binary.BigEndian.Uint32(buf[4:8])
	// timestamp (octet: 8 ~ 11)
	h.SSRC = binary.BigEndian.Uint32(buf[8:12])
	fmt.Println(len(buf))
}

func (h *RTPHeader) String() string {
	var out string
	out += fmt.Sprintf("Version: %v\n", h.Version)
	out += fmt.Sprintf("Padding: %v\n", h.Padding)
	out += fmt.Sprintf("Extention: %v\n", h.Extention)
	out += fmt.Sprintf("CSRC Cout: %v\n", h.CsrcCnt)
	out += fmt.Sprintf("Marker: %v\n", h.Marker)
	out += fmt.Sprintf("Payload Type: %v\n", h.PayloadType)
	out += fmt.Sprintf("Sequence Number: %v\n", h.SequenceNum)
	out += fmt.Sprintf("Timestamp: %v\n", h.Timestamp)
	out += fmt.Sprintf("SSRC: %v\n", h.SSRC)
	return out
}
