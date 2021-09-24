package geohash

const invalid = 0xff

// encoding is store the Geohash alphabet
type encoding struct {
	encode string
	decode [256]byte
}
const base32StrLower="0123456789bcdefghjkmnpqrstuvwxyz"
const base32StrUpper="0123456789BCDEFGHJKMNPQRSTUVWXYZ"
//you can change the default str(lower or upper)
var Base32Encoding= newEncoding(base32StrLower)

func newEncoding(encoder string) *encoding {
	res:=new(encoding)
	res.encode=encoder

	for i:=0;i<len(res.decode);i++{
		res.decode[i]=invalid
	}

	for i:=0;i<len(encoder);i++{
		res.decode[encoder[i]]=byte(i)
	}

	return res
}

// EncodeWithPrecision encode a uint64 number to base32, precision from 1 to 32
func (e *encoding)EncodeWithPrecision(x uint64,precision int) string  {

	if(precision<1||precision>22||x<=0){
		return ""
	}


	lat,lng:=DecodeWithPrecision(x,precision)
	option:=MakeOptions(lng,lat,precision)
	option.useOtherBounding=true

	hash,err:= EncodeWithPrecision(option)
	if err!=nil{
		return ""
	}


	if(precision<1||precision>22||x<=0){
		return ""
	}
	b:=[12]byte{}
	for i := 0; i < 11; i++ {
		 idx:=hash>>(52-((i+1)*5))&0x1f;
		 b[i] = e.encode[idx]
	}
	return string(b[:12])
}

func (e *encoding)Encode(x uint64) string {
	return e.EncodeWithPrecision(x,26)
}

func (e *encoding)Decode( s string) uint64  {
	if len(s)<1 ||len(s)>12 {
		return 0
	}
	ans:=uint64(0)

	for i:=0;i<len(s);i++{
		ans = (ans << 5) | uint64(e.decode[s[i]])
	}
	return ans
}