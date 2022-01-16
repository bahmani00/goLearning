package main

import "fmt"

const helloPrefix = "Hello, "
const helloPrefix_Spanish = "Hola, "
const helloPrefix_French = "Bonjoure, "
const lang_Spanish = "sp"
const lang_French = "fr"


func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix;
	switch lang {
	case lang_French:
		prefix = helloPrefix_French
	case lang_Spanish:
		prefix = helloPrefix_Spanish
	}
	
  return prefix + name
}
func main()  {
	fmt.Println(Hello("World", ""))
}