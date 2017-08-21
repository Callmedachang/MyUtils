package sort


func StraightInsertionSort(array []int)[]int{
	length := len(array)
	for i:=1;i<length;i++{
		if array[i-1]>array[i] {
			j := i
			guard:= array[i]
			for array[j-1]>guard{
				array[j]  = array[j-1]
				if j==1{
					array[0] = guard
					break
				}
				j--
			}
			array[j-1]  = guard
		}
	}
	return array
}
func ShellSort(a []int,n int)[]int{
	var  i, j, k, t int
	k = n / 2
	for k > 0{
		for i = k; i < n; i++{
			t = a[i]
			j = i - k
			for j >= 0 && t < a[j] {
				a[j + k] = a[j]
				j = j - k
			}
			a[j + k] = t
		}
		k /= 2
	}
	return  a
}
func SimpleSelectionSort(a []int)[]int{
	length:= len(a)
	for i:=0;i<length-1;i++{
		var resindex int
		index := i
		var res = a[i]
		for index<length-1{
			if res>a[index+1]{
				res = a[index+1]
				resindex = index+1
			}else{
				resindex = index
			}
			index ++
		}
		temp := a[i]
		a[i] = a[resindex]
		a[resindex] = temp
	}
	return a
}
