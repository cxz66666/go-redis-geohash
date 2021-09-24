package geohash

// DecodeWithPrecision is a reverse of Encode, just get the bit of hash, and update the left or right bounds
func DecodeWithPrecision(hash uint64, precision int) (lat float64, lng float64)  {

	var lat_bit, lng_bit uint64
	nowBound:= bounds
	for i:=0;i<precision;i++ {
		lat_bit = getBit(hash,uint8(precision-i)*2-2)
		lng_bit = getBit(hash,uint8(precision-i)*2-1)

		if lat_bit==0{
			nowBound.MaxLatitude=(nowBound.MaxLatitude+nowBound.MinLatitude)/2
		} else {
			nowBound.MinLatitude=(nowBound.MaxLatitude+nowBound.MinLatitude)/2
		}

		if lng_bit==0{
			nowBound.MaxLongitude=(nowBound.MaxLongitude+nowBound.MinLongitude)/2
		} else {
			nowBound.MinLongitude=(nowBound.MaxLongitude+nowBound.MinLongitude)/2
		}
	}

	return (nowBound.MaxLatitude+nowBound.MinLatitude)/2,(nowBound.MaxLongitude+nowBound.MinLongitude)/2
}


// Decode use DecodeWithPrecision with precision 26 like redis algorithm
func Decode(point uint64) (float64,float64) {
	return DecodeWithPrecision(point,26)
}

//getBit return the uint64(0) or uint64(1) according the bit of bits[pos]
func getBit(bits uint64, pos uint8) uint64 {
	return (bits>>pos)& 0x01
}