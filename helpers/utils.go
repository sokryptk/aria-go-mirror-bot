package helpers

import (
	"fmt"
	"strconv"
)

func convertBytesToReadableFormat(bytesString string) string {
	intSize, _  := strconv.Atoi(bytesString)
	if intSize/(1 << 10) <= (1 << 10) {
		return fmt.Sprintf("%v kB", intSize/(1 << 10))
	} else if intSize/1 << 20 <= 1 << 20 {
		return fmt.Sprintf("%v mB", intSize/(1 << 20))
	} else {
		return fmt.Sprintf("%v gB", intSize/(1 << 30))
	}
}

func GetDownloadFormat(fileName string, rec string, total string, speed string) string {
	convertedRes := convertBytesToReadableFormat(rec)
	convertedTotal := convertBytesToReadableFormat(total)
	convertedSpeed := convertBytesToReadableFormat(speed)

	return fmt.Sprintf("%s => Downloaded %s of %s - at speed : %sps", fileName, convertedRes, convertedTotal, convertedSpeed)
}