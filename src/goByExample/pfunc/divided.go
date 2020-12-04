package pfunc

import "awesomeProjectA/src/goByExample/perror"

func Divided(a,b int)(int ,error){
	/*if a%b!=0{
		return -1,errors.New("不能整除")
	}else {
		return a/b ,nil
	}*/
	if a%b !=0{
		return -1,&perror.DividedArgError{a,b,"不能整除"}
	}else {
		return a/b,nil
	}
}
