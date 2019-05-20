package mem

import (
	"fmt"
	"testing"
	"unsafe"
)

//`$go test`   to run
func TestMemmove(t *testing.T) {

	var s []byte = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	var dst1, dst2 [20]byte

	//len <8
	Memmove(unsafe.Pointer(&dst2[0]), unsafe.Pointer(&s[0]), 5)

	if Memcmp(unsafe.Pointer(&s[0]), unsafe.Pointer(&dst2[0]), 5) != 0 {
		t.Error("test Memmove: failed")
	} else {
		t.Log("test Memmove:  pass")
	}

	copy(dst1[0:5], s[0:5])
	if Memcmp(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&dst2[0]), 5) != 0 {
		t.Error("test Memmove: failed")
	} else {
		t.Log("test Memmove:  pass")
	}

	//len > 8
	Memmove(unsafe.Pointer(&dst2[0]), unsafe.Pointer(&s[0]), len(s))

	if Memcmp(unsafe.Pointer(&s[0]), unsafe.Pointer(&dst2[0]), len(s)) != 0 {
		t.Error("test Memmove: failed")
	} else {
		t.Log("test Memmove:  pass")
	}

	copy(dst1[:], s)
	if Memcmp(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&dst2[0]), len(s)) != 0 {
		t.Error("test Memmove: failed")
	} else {
		t.Log("test Memmove:  pass")
	}

	//overlap copy
	Memmove(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&dst1[1]), len(s)-1)
	if Memcmp(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&s[1]), len(s)-1) != 0 {
		t.Error("test Memmove: failed")
	} else {
		t.Log("test Memmove:  pass")
	}

	Memmove(unsafe.Pointer(&dst2[1]), unsafe.Pointer(&dst2[0]), len(s)-1) //11234
	if Memcmp(unsafe.Pointer(&dst2[1]), unsafe.Pointer(&s[0]), len(s)-1) != 0 {
		t.Error("test Memmove: failed")
		fmt.Println(s)
		fmt.Println(dst2)
	} else {
		t.Log("test Memmove:  pass")
	}

}

func TestMemcmp(t *testing.T) {

	var s []byte = []byte{'1', '2', '1', '2'}
	if Memcmp(unsafe.Pointer(&s[0]), unsafe.Pointer(&s[2]), 2) != 0 {
		t.Error("test Memcmp: failed")
	} else {
		t.Log("test Memcmp:  pass")
	}

	var ii int32 = 0x12345678
	var s2 []int8 = []int8{0x78, 0x56, 0x34, 0x12}
	if Memcmp(unsafe.Pointer(&ii), unsafe.Pointer(&s2[0]), 4) != 0 {
		t.Error("test Memcmp 2: failed")
	} else {
		t.Log("test Memcmp 2:  pass")
	}

	var i3 []int32 = []int32{0x12345678, 0x12345678, 0x12345678, 0x12345678, 0x12345678}
	var s3 []int8 = []int8{0x78, 0x56, 0x34, 0x12, 0x78, 0x56, 0x34, 0x12, 0x78, 0x56, 0x34, 0x12, 0x78, 0x56, 0x34, 0x12, 0x78, 0x56, 0x34, 0x12}

	for i := 0; i < 5; i++ {
		if Memcmp(unsafe.Pointer(&i3[i]), unsafe.Pointer(&s3[4*i]), (5-i)*4) != 0 {
			t.Error(fmt.Sprintf("test Memcmp loop %d: failed", i))
		} else {
			t.Log(fmt.Sprintf("test Memcmp loop %d:  pass", i))
		}
	}

	i3[2] = 0x12395678

	if Memcmp(unsafe.Pointer(&i3[0]), unsafe.Pointer(&s3[0]), 8) != 0 {
		t.Error("test Memcmp 3: failed")
	} else {
		t.Log("test Memcmp 3: pass")
	}

	if Memcmp(unsafe.Pointer(&i3[0]), unsafe.Pointer(&s3[0]), 12) <= 0 {
		t.Error("test Memcmp 4: failed")
	} else {
		t.Log("test Memcmp 4: pass")
	}

	if Memcmp(unsafe.Pointer(&s3[0]), unsafe.Pointer(&i3[0]), 12) >= 0 {
		t.Error("test Memcmp 5: failed")
	} else {
		t.Log("test Memcmp 5: pass")
	}

}

func TestMemset(t *testing.T) {

	var s []byte = []byte{'2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2', '2'}
	var dst1, dst2 [20]byte

	Memset(unsafe.Pointer(&dst2[0]), 0x00, 5)

	if Memcmp(unsafe.Pointer(&dst2[0]), unsafe.Pointer(&dst1[0]), 5) != 0 {
		t.Error("test Memset: failed")
	} else {
		t.Log("test Memset:  pass")
	}

	Memset(unsafe.Pointer(&dst1[0]), '2', 5)
	if Memcmp(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&s[0]), 5) != 0 {
		t.Error("test Memset: failed")
	} else {
		t.Log("test Memset:  pass")
	}

	//len > 8
	Memset(unsafe.Pointer(&dst1[0]), '2', len(s)-1)
	if Memcmp(unsafe.Pointer(&dst1[0]), unsafe.Pointer(&s[0]), len(s)-1) != 0 {
		t.Error("test Memset: failed")
	} else {
		t.Log("test Memset:  pass")
	}

}
