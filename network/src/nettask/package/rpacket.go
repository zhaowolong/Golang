package packet

type Rpacket struct{
	readidx uint32
	buffer *ByteBuffer
	raw bool
}

func NewRpacket(buffer *ByteBuffer,raw bool)(*Rpacket){
	if buffer == nil {
		return nil
	}
	if raw {
		return &Rpacket{readidx:0,buffer:buffer,raw:raw}
	}else
	{
		return &Rpacket{readidx:4,buffer:buffer,raw:raw}
	}
}

func (this *Rpacket) IsRaw()(bool){
	return this.raw
}

func (this *Rpacket) Buffer()(*ByteBuffer){
	return this.buffer
}

func (this *Rpacket) Len()(uint32){
	if this.buffer == nil {
		return 0
	}
	if this.raw{
		return uint32(this.buffer.Len())
	}else{

		len,err := this.buffer.Uint32(0)
		if err != nil {
			return 0
		}
		return len
	}
}

func (this *Rpacket) Uint16()(uint16,error){
	if this.raw{
		return 0,ErrInvaildData
	}
	value,err := this.buffer.Uint16(this.readidx)
	if err != nil {
		return 0,err
	}
	this.readidx += 2
	return value,nil
}

func (this *Rpacket) Uint32()(uint32,error){
	if this.raw{
		return 0,ErrInvaildData
	}
	value,err := this.buffer.Uint32(this.readidx)
	if err != nil {
		return 0,err
	}
	this.readidx += 4
	return value,nil
}

func (this *Rpacket) String()(string,error){
	if this.raw{
		return "",ErrInvaildData
	}
	value,err := this.buffer.String(this.readidx)
	if err != nil {
		return "",err
	}
	this.readidx += (4 + (uint32)(len(value)))
	return value,nil

}

func (this *Rpacket) Binary()([]byte,error){
	if this.raw{
		if this.readidx != 0{
			return nil,ErrInvaildData
		}
		value,err := this.buffer.Binary(this.readidx)
		if err != nil {
			return nil,err
		}
		return value,nil
	}else{
		value,err := this.buffer.Binary(this.readidx)
		if err != nil {
			return nil,err
		}
		this.readidx += (4 + (uint32)(len(value)))
		return value,nil
	}
}
