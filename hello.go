package main

import "fmt"

const helloPrefix = "Hello, "
const helloPrefix_Spanish = "Hola, "
const lang_sp = "sp"


func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix;
	if lang == lang_sp {
		prefix = helloPrefix_Spanish
	}

  return prefix + name
}
func main()  {
	fmt.Println(Hello("World", ""))
}