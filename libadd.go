package main
// 编译命令 go build -buildmode=c-shared -o libadd.so libadd.go
import "C"
//import "strings"
//export parsecode
func parsecode(s *C.char) (*C.char){  //上面那一行注释必须写, 用来提供接口

	code := C.GoString(s)
println("go接受的数据",code)

	return C.CString(code)
}
//export add


func main() {
}