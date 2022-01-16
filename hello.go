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
	if lang == lang_Spanish {
		prefix = helloPrefix_Spanish
	}	else if lang == lang_French {
		prefix = helloPrefix_French
	}
  return prefix + name
}
func main()  {
	fmt.Println(Hello("World", ""))
}