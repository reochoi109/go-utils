package lrc

func Get(message []byte) byte {
	var buf byte
	for _, v := range message {
		buf ^= v
	}
	return buf
}

func Validate(message []byte, checksum byte) bool {
	return Get(message) == checksum
}

func ValidateFrame(frame []byte) bool {
	if len(frame) == 0 {
		return false
	}
	return Validate(frame[:len(frame)-1], frame[len(frame)-1])
}
