package main

import "fmt"


func main() {
	nums := []string{"first","second"}

	for index, value := range nums {
		fmt.Println(index)
		fmt.Println(value)
	}
}


// MAP STRUCTURE
// func main() {
// 	m := make(map[string]int)

// 	m["first"] = 1
// 	m["second"] = 2

// 	fmt.Println(m)

// 	m["first"] = 11
// 	fmt.Println(m)

// 	delete(m, "second"); // delete the key and value
// 	fmt.Println(m)
// }

// ------------------------------------------------------------------------

// SLICE STRUCTURE
// func main() {
// 	s := make([]int, 3,4)

// 	s[0] = 1
// 	s[1] = 2
// 	s[2] = 3

// 	fmt.Println(s);
// 	fmt.Printf("%p\n",s);
// 	fmt.Println(len(s));
// 	fmt.Println(cap(s));

// 	s = append(s, 4)
// 	fmt.Printf("%p\n",s);

// 	s = append(s, 5)
// 	fmt.Printf("%p\n",s);
// } 