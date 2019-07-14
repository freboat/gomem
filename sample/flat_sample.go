package main

import(
    //"time"
    //"reflect"
    "fmt"
    "unsafe"
    //"strconv"
    //"assert"
    "github.com/freboat/gomem/mem"
)

type Student struct {
    Id int64
    Name string
    Value float64
}
type Ptr  = unsafe.Pointer
func main() {

    v := &Student{Id: 101, Name: "Tom",  Value: 98.5}
    container := mem.Parse(*v)
    if container==nil {
        fmt.Println("SetupTI failed")
        return 
    }
    //fmt.Printf("total fields: %d, cap: %d\n", ti.fields,  ti.Pos[ti.fields-1]+ti.Size[ti.fields-1])
    /*
    container := &mem.Container{
        ti : ti,
        Rows : make([]Row, 0, 5),
    }*/
    //container.Rows[0] = make([]byte, ti.Pos[ti.fields-1]+ti.Size[ti.fields-1])
    
    container.Save(Ptr(v))
    
    
    fmt.Println(container.Rows[0])
    
    var str string
     b  := []byte{'a', 'b', 'c', 'd', 'e'}
    
    str = string(b[1:3])
    fmt.Println(str)
    
    var v2 Student
    fmt.Println(v2)
    container.Dump(0, Ptr(&v2))
    
    fmt.Println(v2)
    
}
