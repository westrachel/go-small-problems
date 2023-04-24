package main

import "fmt"

/*
  Write a function that takes an integer that represents seconds
  and returns a formatted version of the time (HH:MM:SS)

  Input:
   > time in seconds that's >= 0
   > Maximum time <= 359999 (99:59:59)
  Output:
   > string in the format HH:MM:SS where
    HH = hours, padded to 2 digits, range: 00 - 99
    MM = minutes, padded to 2 digits, range: 00 - 59
    SS = seconds, padded to 2 digits, range: 00 - 59

  Algorithm:
  i. declare a time string that's initially empty
  ii. divide the given integer by 3600 to find the number of hours
       > find the remainingTime := seconds % 3600
       > convert the number of hours to a string & check if its length
          is 1
           > if it is then pad the front of the string w/ a "0"
       > add the string representation of hours to the time string
          and then add ":"
  iii. divide the remainingTime by 60 to find the number of minutes
       > reassign remainingTime = remainingTime % 60
       > convert the minutes to a string and check if its length is 1
          > if it is then pad the front of the string w/ a "0"
        > add the string representation of minutes to the time string
          and then add ":"
  iv. subtract the number of mins from remaining time to find the amount of seconds
       > seconds = remainingTime - (minutes * 60)
       > convert amount of seconds into a string and pad it if its length is 1
       > add the string representation of minutes to the time string
  v. return the time string
*/

func main() {
	// Test cases:
	fmt.Println(HumanReadableTime(0))     //  "00:00:00"
	fmt.Println(HumanReadableTime(120))   //  "00:02:00"
	fmt.Println(HumanReadableTime(363))   //  "00:06:03"
	fmt.Println(HumanReadableTime(3627))  //  "01:00:27"
	fmt.Println(HumanReadableTime(21895)) //  "06:04:55"
}

func HumanReadableTime(seconds int) string {
	time := ""
	hours, remaining := seconds/3600, seconds%3600
	time += formatTime(hours, true)
	mins, seconds := remaining/60, remaining%60
	time += formatTime(mins, true)
	time += formatTime(seconds, false)
	return time
}

func formatTime(time int, addColonFlag bool) string {
	strTime, colon := fmt.Sprintf("%v", time), ""

	if len(strTime) == 1 {
		strTime = "0" + strTime
	}
	if addColonFlag {
		colon = ":"
	}
	return strTime + colon
}
