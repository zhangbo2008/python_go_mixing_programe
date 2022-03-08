package main

import "C"


//export add
func add(left, right string) string {
    return left + right
}

func main() {
print(add("阿阿阿阿阿阿","aaaaaaaa"))
}