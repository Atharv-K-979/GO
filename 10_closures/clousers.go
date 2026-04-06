package main

func counter() func() int{
	count := 0
	return func() int{
		count++
		return count
	}
}

func main(){
	c1 := counter()
	println(c1()) // 1
	println(c1()) // 2
	
	c2 := counter()
	println(c2())
	println(c1()) // 3
	println(c2()) // 2
}	