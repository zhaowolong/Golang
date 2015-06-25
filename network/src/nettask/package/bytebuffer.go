package packet

import (
	"encoding/binary"
	"fmt"
	//"unsafe"
)
func IsPow2(size uint32) bool{
	return (size&(size-1)) == 0
}

func SizeofPow2(size uint32) uint32{
	if IsPow2(size){
		return size
	}
	size = size -1
	size = size-1
	size = size | (size>>1)
	size = size | (size>>2)
	size = size | (size>>4)
	size = size | (size>>8)
	size = size | (size>>16)
	return size + 1
}

func GetPow2(size uint32) uint32{
	var pow2 uint32 = 0
	if !IsPow2(size) {
		size = (size << 1)
	}
	for size > 1 {
		pow2++
	}
	return pow2
}
const (
	Max_bufsize  uint32 = 32000 
	Max_string_len  uint32 = 32000
	Max_bin_len  uint32 = 32000
)

type ByteBuffer struct {
	buffer []byte
	len uint64
	cap uint64
}


var (
	ErrMaxDataSlotsExceeded     = fmt.Errorf("bytebuffer: Max Buffer Size Exceeded")
	ErrInvaildData              = fmt.Errorf("bytebuffer: Invaild Data")
)


func NewBufferByBytes(bytes []byte)(*ByteBuffer){
	return &ByteBuffer{buffer:bytes,len:(uint64)(len(bytes)),cap:(uint64)(cap(bytes))}
}

/*func NewBufferByOther(other *bytebuffer)(*bytebuffer){
	if other == nil {
		return nil
	}
	buf := &bytebuffer{buffer:make([]byte,other.Cap()),len:other.Len(),cap:other.Cap()}
	//copy data
	copy(buf.buffer[:],other.buffer[:other.Len()])
	return buf
}*/

func NewByteBuffer(size uint32)(*ByteBuffer){
	if size == 0 {
		size = 64
	}else{
		size = SizeofPow2(size)
	}
	return &ByteBuffer{buffer:make([]byte,size),len:0,cap:uint64(size)}
}

func (this *ByteBuffer)Bytes()([]byte){
	return this.buffer
}

func (this *ByteBuffer)Len()(uint64){
	return this.len
}

func (this *ByteBuffer)Cap()(uint64){
	return this.cap
}

func (this *ByteBuffer)expand(newsize uint32)(error){
	newsize = SizeofPow2(newsize)
	if newsize > Max_bufsize {
		return ErrMaxDataSlotsExceeded
	}
	//allocate new buffer
	tmpbuf := make([]byte,newsize)
	//copy data
	copy(tmpbuf[0:], this.buffer[:this.len])
	//replace buffer
	this.buffer = tmpbuf
	this.cap = (uint64)(newsize)
	return nil
}

func (this *ByteBuffer)buffer_check(idx,size uint32)(error){
	if (uint64)(idx+size) > this.cap {
		err := this.expand(idx+size)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *ByteBuffer)PutByte(idx uint32,value byte)(error){
	err := this.buffer_check(idx,1)
	if err != nil {
		return err
	}
	this.buffer[idx] = value
	this.len += 1
	return nil
}

func (this *ByteBuffer)PutUint16(idx uint32,value uint16)(error){
	err := this.buffer_check(idx,2)
	if err != nil {
		return err
	}
	binary.LittleEndian.PutUint16(this.buffer[idx:idx+2],value)
	this.len += 2
	return nil
}

func (this *ByteBuffer)PutUint32(idx uint32,value uint32)(error){
	err := this.buffer_check(idx,4)//(uint32)(unsafe.Sizeof(value)))
	if err != nil {
		return err
	}
	binary.LittleEndian.PutUint32(this.buffer[idx:idx+4],value)
	this.len += 4
	return nil
}


func (this *ByteBuffer)PutString(idx uint32,value string)(error){
	sizeneed := (uint32)(4)
	sizeneed += (uint32)(len(value))
	err := this.buffer_check(idx,sizeneed)
	if err != nil {
		return err
	}

	//first put string len
	this.PutUint32(idx,(uint32)(len(value)))
	
	idx += 4
	//second put string
	copy(this.buffer[idx:],value[:len(value)])
	this.len += (uint64)(len(value))
	return nil
}

func (this *ByteBuffer)PutBinary(idx uint32,value []byte)(error){
	sizeneed := (uint32)(4)
	sizeneed += (uint32)(len(value))
	err := this.buffer_check(idx,sizeneed)
	if err != nil {
		return err
	}

	//first put bin len
	this.PutUint32(idx,(uint32)(len(value)))
	idx += 4
	//second put bin
	copy(this.buffer[idx:],value[:len(value)])
	this.len += (uint64)(len(value))
	return nil
}

func (this *ByteBuffer)PutRawBinary(value []byte)(error){
	sizeneed := (uint32)(len(value))
	err := this.buffer_check(uint32(this.len),sizeneed)
	if err != nil {
		return err
	}
	//second put bin
	copy(this.buffer[this.len:],value[:len(value)])
	this.len += (uint64)(len(value))
	return nil
}

func (this *ByteBuffer)Uint16(idx uint32)(ret uint16,err error){
	if (uint64)(idx + 2) > this.len {
		ret = 0
		err = ErrInvaildData
		return
	}
	ret = binary.LittleEndian.Uint16(this.buffer[idx:idx+2])
	err = nil
	return
}

func (this *ByteBuffer)Uint32(idx uint32)(ret uint32,err error){
	if (uint64)(idx + 4) > this.len {
		ret = 0
		err = ErrInvaildData
		return
	}
	ret = binary.LittleEndian.Uint32(this.buffer[idx:idx+4])
	err = nil
	return
}

func (this *ByteBuffer)String(idx uint32)(ret string,err error){
	if (uint64)(idx + 4) > this.len {
		err = ErrInvaildData
		return
	}
	//read string len
	str_len,_ := this.Uint32(idx)
	idx += 4
	if (uint64)(idx + str_len) > this.len {
		err = ErrInvaildData
		return
	}
	err = nil
	ret = string(this.buffer[idx:idx+str_len])
	return
}


func (this *ByteBuffer)Binary(idx uint32)(ret []byte,err error){
	if (uint64)(idx + 4) > this.len {
		err = ErrInvaildData
		return
	}
	//read bin len
	bin_len,_ := this.Uint32(idx)
	idx += 4
	if (uint64)(idx + bin_len) > this.len {
		err = ErrInvaildData
		return
	}
	err = nil
	ret = this.buffer[idx:idx+bin_len]
	return
}
