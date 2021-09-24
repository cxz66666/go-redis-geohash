package geohash


import "errors"

type boundingBox struct {
	MaxLatitude  float64
	MinLatitude  float64
	MaxLongitude float64
	MinLongitude float64
}

type option struct {
	lng float64
	lat float64
	precision int
	useOtherBounding bool
}

//!!warning, this MaxLatitude is not 90 but 85.05112878
var bounds =boundingBox{MaxLatitude: 85.05112878, MaxLongitude: 180, MinLatitude: -85.05112878, MinLongitude: -180}
var boundsH=boundingBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}
// EncodeWithPrecision accept a specifical precision for encode, by default precision is 26(like redis)
// also we provide another way to calculate geohash(which is annotation)

func MakeOptions(lng,lat float64,precision int) option  {
	return option{
		lng:lng,
		lat:lat,
		precision: precision,
		useOtherBounding: false,
	}
}

func EncodeWithPrecision(option option) (uint64,error) {
	var nowBound boundingBox
	if option.useOtherBounding {
		nowBound=boundsH
	} else {
		nowBound=bounds
	}
	var lat=option.lat
	var lng=option.lng
	var precision=option.precision
	if lat < nowBound.MinLatitude || lat > nowBound.MaxLatitude || lng < nowBound.MinLongitude || lng > nowBound.MaxLongitude {
		err := errors.New("Coordinate out of bounds")
		return 0,err
	}
	var hash uint64 //ans
	var lat_bit, lng_bit uint64
	for i := 0; i < precision; i++ {
		if nowBound.MaxLatitude-lat >= lat-nowBound.MinLatitude {
			lat_bit = 0
			nowBound.MaxLatitude = (nowBound.MaxLatitude + nowBound.MinLatitude) / 2
		} else {
			lat_bit = 1
			nowBound.MinLatitude = (nowBound.MaxLatitude + nowBound.MinLatitude) / 2
		}

		if nowBound.MaxLongitude-lng >= lng-nowBound.MinLongitude {
			lng_bit = 0
			nowBound.MaxLongitude = (nowBound.MaxLongitude + nowBound.MinLongitude) / 2
		} else {
			lng_bit = 1
			nowBound.MinLongitude = (nowBound.MaxLongitude + nowBound.MinLongitude) / 2
		}
		hash <<= 1
		hash += lng_bit
		hash <<= 1
		hash += lat_bit
	}
	return hash,nil


}


// Encode use EncodeWithPrecision with precision 26 like redis algorithm
func Encode(lng,lat float64) (uint64,error)  {
	return EncodeWithPrecision(MakeOptions(lng,lat,26))
}


// EncodeWithPrecisionC is same as EncodeWithPrecision, but use another algorithm to calculate hash, which is used by redis, it must be 26 precision
// You can change it in Encode func,
func EncodeWithPrecisionC(option option) (uint64,error) {
	var nowBound boundingBox
	if option.useOtherBounding {
		nowBound=boundsH
	} else {
		nowBound=bounds
	}
	lat:=option.lat
	lng:=option.lng
	if lat < nowBound.MinLatitude || lat > nowBound.MaxLatitude || lng < nowBound.MinLongitude || lng > nowBound.MaxLongitude {
		err := errors.New("Coordinate out of bounds")
		return 0,err
	}
	var hash uint64 //ans
	var lat_offset,long_offset float64
	lat_offset=(lat-nowBound.MinLatitude)/(nowBound.MaxLatitude-nowBound.MinLatitude)
	long_offset=(lng-nowBound.MinLongitude)/(nowBound.MaxLongitude-nowBound.MinLongitude)
	lat_offset*=1<<26;
	long_offset*=1<<26;
	hash=interleave64(int32(lat_offset),int32(long_offset))
	return hash,nil
}


// interleave64 is used for two int number, and will return their binary bit one by one
// for example latOffset is 10110, and lngOffset is 00001, then the ans is 1000101001, so as the 32-bits int
func interleave64(latOffset int32, lngOffset int32) uint64{
	B :=[]uint64{0x5555555555555555, 0x3333333333333333, 0x0F0F0F0F0F0F0F0F, 0x00FF00FF00FF00FF, 0x0000FFFF0000FFFF}
	S :=[]uint8{1, 2, 4, 8, 16}
	x := uint64(latOffset)
	y := uint64(lngOffset)
	x = (x | (x << S[4])) & B[4];//第一轮
	y = (y | (y << S[4])) & B[4];

	x = (x | (x << S[3])) & B[3];//第二轮
	y = (y | (y << S[3])) & B[3];

	x = (x | (x << S[2])) & B[2];//第三轮
	y = (y | (y << S[2])) & B[2];
	x = (x | (x << S[1])) & B[1];
	y = (y | (y << S[1])) & B[1];
	x = (x | (x << S[0])) & B[0];
	y = (y | (y << S[0])) & B[0];
	return x|(y<<1)
}