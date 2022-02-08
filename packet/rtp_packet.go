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
	CSRC        []uint32
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
	h.CsrcCnt = buf[0] & 0xf // (0xf == 1111)

	// (octed: 1)
	// marker(bit: 0)
	h.Marker = buf[1]>>7 == 1

	// payload type(bit: 1 ~ 7)
	h.PayloadType = buf[1] & 0x7f // (0x7f == 01111111)

	// sequence number(octet: 2 ~ 3)
	h.SequenceNum = binary.BigEndian.Uint16(buf[2:4])

	// timestamp (octet: 4 ~ 7)
	h.Timestamp = binary.BigEndian.Uint32(buf[4:8])
	// ssrc (octet: 8 ~ 11)
	h.SSRC = binary.BigEndian.Uint32(buf[8:12])

	// csrc (octet: 12 ~ )
	h.CSRC = make([]uint32, 0, h.CsrcCnt)

	octetCnt := 12

	for i := uint8(0); i < h.CsrcCnt; i++ {
		csrc := binary.BigEndian.Uint32(buf[octetCnt : octetCnt+4])
		h.CSRC = append(h.CSRC, csrc)
		octetCnt += 4
	}
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
	out += fmt.Sprintf("CSRC: %v\n", h.CSRC)
	return out
}
