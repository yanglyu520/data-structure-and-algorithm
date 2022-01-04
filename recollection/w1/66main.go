func plusOne(digits []int) []int {
    //find the last element of the array
    n := len(digits) - 1
    //if the last digit is less than 9
    if digits[n] < 9 {
        digits[n]++
        return digits
    }else {
        for i:=n; i>=0; i-- {
           if digits[i] != 9 {
               digits[i]++
               for h :=n;h>i;h--{
                   digits[h]=0
               }
               return digits
           }
        }
      }
	  //if the digits are like 999999... plus one make it 1000000...
    if ifdigitsequalto9(digits) {
         digits[0]=1
            for j :=1;j<=n; j++{
              digits[j] = 0
            }
            digits=append(digits,0)
    }
    return digits
}

func ifdigitsequalto9(digits []int) bool {
    for k,_ = range digits{
        if digits[k] != digits[0] && digits[]!=9 {
            return false
        }
    }
    return true
}
