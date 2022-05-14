package easybits

import (
	"testing"
)

func TestFromBool(t *testing.T) {
	result := FromBool([]bool{true, true, true, false, false, true})
	if result.Length() != 6 {
		t.Error(result)
	}
	if !(result.Get(0) == true &&
		result.Get(1) == true &&
		result.Get(2) == true &&
		result.Get(3) == false &&
		result.Get(4) == false &&
		result.Get(5) == true) {
		t.Error(result)
	}
}

func TestFromByte1(t *testing.T) {
	result := FromByte(6, []byte{0b11100100})
	if result.Length() != 6 {
		t.Error(result)
	}
	if !(result.Get(0) == true &&
		result.Get(1) == true &&
		result.Get(2) == true &&
		result.Get(3) == false &&
		result.Get(4) == false &&
		result.Get(5) == true) {
		t.Error(result)
	}
}

func TestFromByte2(t *testing.T) {
	result := FromByte(14, []byte{0b11100100, 0b01011100})
	if result.Length() != 14 {
		t.Error(result)
	}
	if !(result.Get(0) == true &&
		result.Get(1) == true &&
		result.Get(2) == true &&
		result.Get(3) == false &&
		result.Get(4) == false &&
		result.Get(5) == true &&
		result.Get(6) == false &&
		result.Get(7) == false &&
		result.Get(8) == false &&
		result.Get(9) == true &&
		result.Get(10) == false &&
		result.Get(11) == true &&
		result.Get(12) == true &&
		result.Get(13) == true) {
		t.Error(result)
	}
}

func TestToBytes(t *testing.T) {
	original := []byte{0b11100100, 0b01011100}
	result := FromByte(14, original).ToBytes()
	if len(result) != 2 {
		t.Error(result)
	}
	if result[0] != original[0] {
		t.Error(result)
	}
	if result[1] != original[1] {
		t.Error(result)
	}
}

func TestShiftLeftRotation1(t *testing.T) {
	bits := FromBool([]bool{true, true, true, false, false, true})
	result := bits.ShiftLeftRotation(1)
	if result.Length() != 6 {
		t.Error(result)
	}
	if !(result.Get(0) == true &&
		result.Get(1) == true &&
		result.Get(2) == false &&
		result.Get(3) == false &&
		result.Get(4) == true &&
		result.Get(5) == true) {
		t.Error(result)
	}
}

func TestShiftLeftRotation2(t *testing.T) {
	bits := FromBool([]bool{true, true, true, false, false, true})
	result := bits.ShiftLeftRotation(2)
	if result.Length() != 6 {
		t.Error(result)
	}
	if !(result.Get(0) == true &&
		result.Get(1) == false &&
		result.Get(2) == false &&
		result.Get(3) == true &&
		result.Get(4) == true &&
		result.Get(5) == true) {
		t.Error(result)
	}
}

func TestXor(t *testing.T) {
	a := FromBool([]bool{true, true, true, false, false, true})
	b := FromBool([]bool{true, false, true, false, true, true})
	result, err := a.Xor(b)
	if err != nil {
		t.Error(err)
	}
	if !(result.Get(0) == false &&
		result.Get(1) == true &&
		result.Get(2) == false &&
		result.Get(3) == false &&
		result.Get(4) == true &&
		result.Get(5) == false) {
		t.Error(result)
	}
}
