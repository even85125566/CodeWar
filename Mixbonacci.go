func codewar

func Mixbonacci(pattern []string,length int)[]int64{
    orilength:=len(pattern)
	if (len(pattern))!=length{
  for i := 0; i < length/orilength; i++ {

  }

	}
	
	return nil
}




func Fibonacci(number int)int{
	if (number<2){
		return number
	}

	return Fibonacci(number-1)+Fibonacci(number-2)
}
func Pell(number int)int{
	if (number<2){
		return number
	}

	return 2*Pell(number-1)+Pell(number-2)
	
}
func Padovan(number int)int{

	if (number<3){
		return 1
	}
	return Padovan(number-2)+Padovan(number-3)
}

func Jacobstahi(number int)int{
if (number<2){
	return number
}

return Jacobstahi(number-1)+2*Jacobstahi(number-2)
}
func Tribonacci(number int)int{
	if (number<2){
		return 0
	}
	if (number==2){
		return 1
	}

	return Tribonacci(number-1)+Tribonacci(number-2)+Tribonacci(number-3)
	
}
func Tetranacci(number int)int{
	
   switch number {
   case 0:
	return 0
   case 1:
	return 1
   case 2:
	return 1
   case 3:
	return 2
   default:
	return Tetranacci(number-1)+Tetranacci(number-2)+Tetranacci(number-3)+Tetranacci(number-4)
	
   }


	
}