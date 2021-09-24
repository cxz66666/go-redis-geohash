# go-redis-geohash

A Go implementation of integer redis geohashing algorithm

maybe sometimes you need to calculate the `geohash` on your backend, not on redis-server. So maybe you will find this [repo](https://github.com/mmcloughlin/geohash), but you find they provide the result of geohash is not equal with redis-server! 


### Usage

1. Encode
~~~go
package main

import "fmt"
import "./geohash"

func main()  {
	hash,err:=geohash.Encode(123.4,44.123)
	if err!=nil{
		panic(err) 
	}
	fmt.Println(hash)
	
	//this precision means use how many bits to represent lat or lng, not total,
	// so redis use 26 bits to represent them for each.
	option:=geohash.MakeOptions(11.1111,22.2222,26)
	//this can use your own precision
	hash,err=geohash.EncodeWithPrecision(option)
	if err!=nil{
		panic(err)
	}
	fmt.Println(hash)
	
	//use the algorithm in redis, but most time ans is equal
	hash,err=geohash.EncodeWithPrecisionC(option)
	if err!=nil{
		panic(err)
	}
	fmt.Println(hash)

}


~~~

2. Decode
~~~go
package main

import "fmt"
import "./geohash"

func main()  {
	hash,err:=geohash.Encode(123.4,44.123)
	if err!=nil{
		panic(err)
	}
	fmt.Println(hash)
	//use redis-algorithm to decode
	lat,lng:=geohash.Decode(hash)
	fmt.Println(lat,lng)
	
	//if you use EncodeWithPrecision to encode, so you need use this to decode
	lat,lng=geohash.DecodeWithPrecision(hash,26)
}

~~~


3. Base32 \
(maybe has bug, because there is a lot of version to choose, and I'm not sure which to use)
~~~go
package main

import "fmt"
import "./geohash"

func main()  {
	hash,_:=geohash.Encode(123.4,44.123)
    base32Str:= geohash.Base32Encoding.Encode(hash)
	fmt.Println(base32Str)
}

~~~