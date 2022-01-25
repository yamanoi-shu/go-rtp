package main

type RTPHeader struct {
	Version     uint8
	Padding     bool
	Extention   bool
	CsrcCnt     uint8
	Marker      bool
	PayloadType uint8
	SequenceNum uint16
	TimeStamp   uint32
	SSRC        uint32
}

type RTPPacket struct {
	Header  RTPHeader
	Payload []byte
}

func NewRTPPacket() *RTPPacket {
	return &RTPPacket{
		Version:   2,
		Padding:   true,
		Extention: false,
		Marker:    false,
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

}
