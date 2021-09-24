package main

import (
	"./geohash"
	"log"
	"testing"
)


// you can check the answer with your redis, and the result is equal with it!
func TestEncode(t *testing.T)  {
	geo, err :=geohash.Encode(  120.37582129240036011, 31.5603669915025975 )
	if err != nil {
		log.Fatal(err)
	}
	log.Print(geo)

	lat, lng := geohash.Decode(geo)
	log.Print(lat)
	log.Print(lng)


	geo, err = geohash.EncodeWithPrecision(geohash.MakeOptions(lng,lat,26))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(geo)

	lat, lng = geohash.DecodeWithPrecision(geo,26)
	log.Print(lat)
	log.Print(lng)

}



