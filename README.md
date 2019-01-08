# dateutil
DateUtils for Go Lang

Goal is to impliment date related operations, starting with function to find deference between 2 dates in year,month,day,hour,minute and 
seconds.

usage 

  from := time.Date(1993, 4, 8, 22, 50, 42, 0, time.UTC)
	to := time.Date(2019, 1, 9, 00, 40, 39, 0, time.UTC)
  dif := dateutil.Difference(from, to)
	fmt.Printf("%d Years %d Months %d Days %d H %d M %d S\n", dif.Years, dif.Months, dif.Days, dif.Hours, dif.Minutes, dif.Seconds)
  
  out put 
    25 Years 9 Months 0 Days 1 H 49 M 57 S
