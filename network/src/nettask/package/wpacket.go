package packet

import (
	"encoding/binary"
	"fmt"
)

type Wpacket struct {
	writeidx      uint32
	buffer        *ByteBuffer
	raw           bool
	Fn_sendfinish func(interface{}, *Wpacket)
}

func NewWpacket(buffer *ByteBuffer, raw bool) *Wpacket {
	if buffer == nil {
		return nil
	}
	if raw {
		return &Wpacket{writeidx: 0, buffer: buffer, raw: raw}
	} else {
		buffer.PutUint32(0, 0)
		return &Wpacket{writeidx: 4, buffer: buffer, raw: raw}
	}
}

func (this *Wpacket) IsRaw() bool {
	return this.raw
}

func (this *Wpacket) Buffer() *ByteBuffer {
	return this.buffer
}

func (this *Wpacket) PutUint16(value uint16) error {
	if this.buffer == nil {
		return ErrInvaildData
	}
	if this.raw {
		return ErrInvaildData
	}
	size, err := this.buffer.Uint32(0)
	if err != nil {
		return err
	}
	err = this.buffer.PutUint16(this.writeidx, value)
	if err != nil {
		return err
	}
	size += 2
	this.writeidx += 2
	binary.LittleEndian.PutUint32(this.buffer.Bytes()[0:4], size)
	return nil
}

func (this *Wpacket) PutUint32(value uint32) error {
	if this.buffer == nil {
		return ErrInvaildData
	}
	if this.raw {
		return ErrInvaildData
	}
	size, err := this.buffer.Uint32(0)
	if err != nil {
		return err
	}
	err = this.buffer.PutUint32(this.writeidx, value)
	if err != nil {
		return err
	}
	size += 4
	this.writeidx += 4
	binary.LittleEndian.PutUint32(this.buffer.Bytes()[0:4], size)
	return nil
}

func (this *Wpacket) PutString(value string) error {
	if this.buffer == nil {
		return ErrInvaildData
	}
	if this.raw {
		return ErrInvaildData
	}
	size, err := this.buffer.Uint32(0)
	if err != nil {
		return err
	}
	err = this.buffer.PutString(this.writeidx, value)
	if err != nil {
		return err
	}
	size += (4 + (uint32)(len(value)))
	this.writeidx += (4 + (uint32)(len(value)))
	binary.LittleEndian.PutUint32(this.buffer.Bytes()[0:4], size)
	fmt.Println("buffsize:", size)
	return nil
}

func (this *Wpacket) PutBinary(value []byte) error {
	if this.buffer == nil {
		return ErrInvaildData
	}
	if this.raw {
		if this.writeidx != 0 {
			return ErrInvaildData
		}
		err := this.buffer.PutBinary(this.writeidx, value)
		if err != nil {
			return err
		}
		this.writeidx += uint32(len(value))
	} else {
		size, err := this.buffer.Uint32(0)
		if err != nil {
			return err
		}
		err = this.buffer.PutBinary(this.writeidx, value)
		if err != nil {
			return err
		}
		size += (4 + (uint32)(len(value)))
		this.writeidx += (4 + (uint32)(len(value)))
		binary.LittleEndian.PutUint32(this.buffer.Bytes()[0:4], size)
	}
	return nil
}
