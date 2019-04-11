package main

import "log"

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
func main() {
	// Println写到标准日志输出
	log.Println("message")
	// Fatalln在调用Println()之后会调用os.Exit(1)
	log.Fatalln("fetal message")
	// Panicln在调用Println()之后会调用panic()
	log.Panicln("panic message")
}
