package easybits

import (
	"errors"
	"fmt"
)

type Bits struct {
	internal []bool
}

func FromBool(bools []bool) *Bits {
	return &Bits{internal: bools}
}

// length=6, bytes=0b10101100, result=tftftt
func FromByte(length int, bytes []byte) *Bits {
	var bits []bool
	for i := 0; i < length; i++ {
		bits = append(bits, bit(i, bytes))
	}
	return FromBool(bits)
}

// get bit for index
func bit(index int, bytes []byte) bool {
	switch index % 8 {
	case 0:
		return bytes[index/8]&0b10000000 != 0
	case 1:
		return bytes[index/8]&0b01000000 != 0
	case 2:
		return bytes[index/8]&0b00100000 != 0
	case 3:
		return bytes[index/8]&0b00010000 != 0
	case 4:
		return bytes[index/8]&0b00001000 != 0
	case 5:
		return bytes[index/8]&0b00000100 != 0
	case 6:
		return bytes[index/8]&0b00000010 != 0
	case 7:
		return bytes[index/8]&0b00000001 != 0
	}
	return false
}

func (b *Bits) Length() int {
	return len(b.internal)
}

// true, true, true, true, false, false, false, false To []byte(0b11110000)
// true, true, true To []byte(0b11100000)
func (bits *Bits) ToBytes() []byte {
	var bytes []byte
	var b byte = 0b00000000
	for i, v := range bits.internal {
		if v {
			b = b | 0b00000001
		}
		if i != 0 {
			if (i+1)%8 == 0 {
				bytes = append(bytes, b)
				b = 0b00000000
			} else if len(bits.internal)-1 == i {
				b = b << (8 - (len(bits.internal))%8)
				bytes = append(bytes, b)
			}
		}
		b = b << 1
	}
	return bytes
}

func (bits *Bits) PrintBits() {
	fmt.Print("Bits:")
	for i := 0; i < bits.Length(); i++ {
		if bits.internal[i] {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if i != 0 && (i+1)%8 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func (bits *Bits) Get(index int) bool {
	return bits.internal[index]
}

// bits=tttf, shift=2, result=tftt
func (bits *Bits) ShiftLeftRotation(shift int) *Bits {
	shifted := bits.internal[shift:]
	remaining := bits.internal[:shift]
	for i := 0; i < shift; i++ {
		shifted = append(shifted, remaining[i])
	}
	b := FromBool(shifted)
	return b
}

func (a *Bits) Xor(b *Bits) (*Bits, error) {
	if a.Length() != b.Length() {
		return nil, errors.New("can't take xor between different length bits")
	}
	var xor []bool
	for i, v := range a.internal {
		if v != b.internal[i] {
			xor = append(xor, true)
		} else {
			xor = append(xor, false)
		}
	}
	bits := FromBool(xor)
	return bits, nil
}
