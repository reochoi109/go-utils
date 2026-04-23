package lrc

import (
	"testing"
)

func TestGet_Empty(t *testing.T) {
	if got := Get(nil); got != 0 {
		t.Fatalf("Get(nil) = %d, want 0", got)
	}

	if got := Get([]byte{}); got != 0 {
		t.Fatalf("Get(empty) = %d, want 0", got)
	}
}

func TestGet_KnownVector(t *testing.T) {
	msg := []byte{0x01, 0x02, 0x03}
	want := byte(0x00)
	if got := Get(msg); got != want {
		t.Fatalf("Get(% x) = 0x%02X, want 0x%02X", msg, got, want)
	}
}

func TestValidate_Match(t *testing.T) {
	msg := []byte{0x01, 0x02, 0x30}
	sum := Get(msg)
	if ok := Validate(msg, sum); !ok {
		t.Fatalf("validate should be true for matching checksum")
	}
}

func TestValidate_Mismatch(t *testing.T) {
	msg := []byte{0x01, 0x20, 0x30}
	if ok := Validate(msg, 0xFF); ok {
		t.Fatalf("validation should be false for mismatching checksum")
	}
}

func TestValidateFrame_Empty(t *testing.T) {
	if ok := ValidateFrame(nil); ok {
		t.Fatalf("ValidateFrame(nil) = true, want false")
	}
	if ok := ValidateFrame([]byte{}); ok {
		t.Fatalf("ValidateFrame(empty) = true, want false")
	}
}

func TestValidateFrame_Match(t *testing.T) {
	msg := []byte{0xAA, 0xBB, 0xCC}
	sum := Get(msg)
	frame := append(append([]byte(nil), msg...), sum)

	if ok := ValidateFrame(frame); !ok {
		t.Fatalf("ValidateFrame should be true for valid frame")
	}
}

func TestValidateFrame_Mismatch(t *testing.T) {
	msg := []byte{0xAA, 0xBB, 0xCC}
	sum := Get(msg)
	frame := append(append([]byte(nil), msg...), sum^0x01)

	if ok := ValidateFrame(frame); ok {
		t.Fatalf("ValidateFrame should be false for invalid frame")
	}
}
