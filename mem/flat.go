package mem


//todo list: 
//1. time support
//2. nest struct
//3. null and ORM support

import(
    //"time"
    "reflect"
    //"fmt"
    "unsafe"
)

type Row  = []byte
type Ptr  = unsafe.Pointer
type PFunc = func([]byte, Ptr, *TypeInfo, int) Ptr   //dest, nth-field,  src


const (
    STRING = iota //varchar[n]              
    INT           //int64
    FLOAT         //float64
    TIME          //
)



type TypeInfo struct {
    fields int
    Type  []uint
    Pos   []int        //[]byte pos
    Size  []int
    Offset []uintptr   //struct field offset
    Save   []PFunc 
    Dump   []PFunc
}

type Container struct {
    ti *TypeInfo
    Rows  []Row
}

func memSave(dest []byte, src Ptr, ti *TypeInfo, n int) Ptr {
   
    return Memcpy(Ptr(&dest[ti.Pos[n]]), src, ti.Size[n])
}

func memDump(src []byte, dest Ptr, ti *TypeInfo, n int) Ptr {
   
    return Memcpy(dest, Ptr(&src[ti.Pos[n]]), ti.Size[n])
}

func strSave(dest []byte, src Ptr, ti *TypeInfo, n int) Ptr {
    p := (*string)(src)
    
    size  := (*uint16)(Ptr(&dest[ti.Pos[n]]))

    *size = uint16(len(*p))
    copy(dest[ti.Pos[n]+2:], *p)
    return nil
}

func strDump(src []byte, dest Ptr, ti *TypeInfo, n int) Ptr {
    p := (*string)(dest)
    size  := (*uint16)(Ptr(&src[ti.Pos[n]]))
    *p = string(src[ti.Pos[n]+2:ti.Pos[n]+2+ int(*size)])
    
    return nil
}


func (container *Container) Save(obj Ptr) {
    container.Rows = append(container.Rows, make([]byte, container.ti.Pos[container.ti.fields-1] + container.ti.Size[container.ti.fields-1]))
    n := len(container.Rows)
    for i := 0; i < container.ti.fields; i++ {
        container.ti.Save[i](container.Rows[n-1], Ptr(uintptr(obj)+container.ti.Offset[i]), container.ti, i)
    }
}    

func (container *Container) Dump(n int, obj Ptr) {
    
    for i := 0; i < container.ti.fields; i++ {
        container.ti.Dump[i](container.Rows[n], Ptr(uintptr(obj)+container.ti.Offset[i]), container.ti, i)
    }
    
}


func Parse(intf interface {}) (* Container){
    
    fields := reflect.TypeOf(intf)
    values := reflect.ValueOf(intf)
    
    num := fields.NumField()
    ti := &TypeInfo {
        fields: num,
        Type : make([]uint, num),
        Pos : make([]int, num),
        Size : make([]int, num),
        Offset : make([]uintptr, num),
        Save : make([]PFunc, num),
        Dump : make([]PFunc, num),
    }
    
    var length int = 0
    for i := 0; i < num; i++ {
        field := fields.Field(i)
        value := values.Field(i)
        ti.Save[i] = memSave
        ti.Dump[i] = memDump
        switch value.Kind() {
            case reflect.String:
                ti.Type[i] = STRING
                ti.Size[i] = 2+64     //2bytes head for varchar len,  another maxlen should acquire from tag
                ti.Save[i] = strSave
                ti.Dump[i] = strDump
            case reflect.Int64:
                ti.Type[i] = INT
                ti.Size[i] = 8
            case reflect.Float64:
                ti.Type[i] = INT
                ti.Size[i] = 8
            
            //case default:
              //  fmt.Printf("Not support data type: Time\n")
              //  return nil
            
        }
        ti.Pos[i] = length
        ti.Offset[i] = field.Offset
        length += ti.Size[i]
    }
    container := &Container{
        ti : ti,
        Rows : make([]Row,0),
    }
    return container
}

