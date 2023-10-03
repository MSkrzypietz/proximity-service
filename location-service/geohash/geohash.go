package geohash

import (
	"strconv"
)

func CalcGeohash(latitude float64, longitude float64) string {
	const geohashLength = 6
	const base32Digits = 5

	binaryGeohash := ""
	minLongitude := -180.0
	maxLongitude := 180.0
	minLatitude := -90.0
	maxLatitude := 90.0
	for len(binaryGeohash) < geohashLength*base32Digits {
		midLongitude := (minLongitude + maxLongitude) / 2
		if longitude <= midLongitude {
			binaryGeohash += "0"
			maxLongitude = midLongitude
		} else {
			binaryGeohash += "1"
			minLongitude = midLongitude
		}

		midLatitude := (minLatitude + maxLatitude) / 2
		if latitude <= midLatitude {
			binaryGeohash += "0"
			maxLatitude = midLatitude
		} else {
			binaryGeohash += "1"
			minLatitude = midLatitude
		}
	}

	return encodeGeohash(chunk(binaryGeohash, 5))
}

func chunk(input string, size int) []string {
	result := []string{}

	for i := 0; i < len(input); i += size {
		result = append(result, input[i:i+size])
	}

	return result
}

func encodeGeohash(binaryChunks []string) string {
	const encoding = "0123456789bcdefghjkmnpqrstuvwxyz"
	result := ""

	for _, binaryChunk := range binaryChunks {
		decimalValue, _ := strconv.ParseInt(binaryChunk, 2, 6)
		result += string(encoding[decimalValue])
	}

	return result
}
