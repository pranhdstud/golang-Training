package main
import (
	"fmt"
)

var days_in_feb int = 28
func main() {
	var day,month,year int
    fmt.Printf("Enter date (MM/DD/YYYY): ")
    fmt.Scanf("%d/%d/%d",&month,&day,&year)
    dayCount := checkDayCount(month,day,year)
    fmt.Printf("The day count for date %d/%d/%d is %d ",month,day,year,dayCount)
}


func checkDayCount (month int, day int, year int) int{

	if year%4 == 0 && year % 100 != 0 || year % 400 == 0{
        days_in_feb = 29
	}

	buffer := day
	switch month {
	case 02:
		buffer += 31
		break
	case 03:
		buffer += 31+days_in_feb
		break
	case 04:
		buffer += 31+days_in_feb+31
		break
    case 05:
		buffer += 31+days_in_feb+31+30
		break
	case 06:
		buffer += 31+days_in_feb+31+30+31
		break
	case 07:
		buffer += 31+days_in_feb+31+30+31+30
		break
	case 8:
		buffer += 31+days_in_feb+31+30+31+30+31
		break
	case 9:
		buffer += 31+days_in_feb+31+30+31+30+31+31
		break	
	case 10:
		buffer += 31+days_in_feb+31+30+31+30+31+31+30
		break
	case 11:
		buffer += 31+days_in_feb+31+30+31+30+31+31+30+31
		break
	case 12:
		buffer += 31+days_in_feb+31+30+31+30+31+31+30+31+30
		break	

    }

return buffer
}