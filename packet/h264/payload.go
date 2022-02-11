package h264

type H264Payload struct {
	NALUType *NALUType
}

func (p *H264Payload) Depacketize(buf []byte) []byte {

	p.NALUType = &NALUType{}
	// NAL Unit Header
	p.NALUType.Unmarshal(buf[0])
	buf = buf[1:]
	return buf
}

func (p *H264Payload) String() string {
	var out string

	out += "NAL Unit Header\n"
	out += p.NALUType.String()

	return out
}
