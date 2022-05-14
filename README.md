# easybits

easybits is easy(not simple) bit operation for Go.

```go
a := FromBool([]bool{true, true, true, false, false, true})
b := FromByte(6, []byte{0b10101100})

// XOR
xor, _ := a.Xor(b)
// 01001000
fmt.Printf("%08b\n", xor.ToBytes())

// LEFT SHIFT(Rotation)
// 10011100
fmt.Printf("%08b\n", a.ShiftLeftRotation(2).ToBytes())
```
