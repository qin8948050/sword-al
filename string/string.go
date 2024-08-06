package string
import "strconv"

func Parse(input int64) string{
	
	data:=strconv.FormatInt(input,10)
	if len(data)<=3{
		return data

	}
	originalIndex:=len(data)-1
	newIndex:=originalIndex+(len(data)/3)
	num:=0
	result:=make([]byte, newIndex+1)
	for originalIndex>=0 && newIndex>=originalIndex {
		if num==3 {
			result[newIndex]=','
			num=0
		} else {
			result[newIndex]=data[originalIndex]
			num++
			originalIndex--
		}
		newIndex--
	}
	return string(result)
}