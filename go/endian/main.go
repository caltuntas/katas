package main

import "fmt"
import "encoding/binary"
import "strconv"
//import "encoding/hex"

func main() {
  learray := []byte{0x00,0x00,0x00,0x40}
  p0 := learray[0]
  p1 := learray[1]
  p2 := learray[2]
  p3 := learray[3]
  sp0 := learray[0] << 0
  sp1 := learray[1] << 8
  sp2 := learray[2] << 16
  sp3 := learray[3] << 24
  i := uint32(learray[0])<<0 | uint32(learray[1])<<8 | uint32(learray[2])<<16 | uint32(learray[3])<<24;
  fmt.Println("value of i")
  fmt.Println(i)
  data := binary.BigEndian.Uint32(learray)
  fmt.Println("before shifting bits")
  fmt.Println(p0)
  fmt.Println(p1)
  fmt.Println(p2)
  fmt.Println(p3)
  fmt.Println("after shifting bits")
  fmt.Println(sp0)
  fmt.Println(sp1)
  fmt.Println(sp2)
  fmt.Println(sp3)
  fmt.Println(data)
  val := sp0 | sp1 | sp2 | sp3
  fmt.Println("after or operation")
  fmt.Println(val)
   //encodedString := hex.EncodeToString(val)
   //  fmt.Println("Encoded Hex String: ", encodedString)
   n := int64(123)

   fmt.Println(strconv.FormatInt(n, 2))
   fmt.Printf("%032b",i)
}
