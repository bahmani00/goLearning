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
	
  return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string{
	switch lang {
	case lang_French:
		return helloPrefix_French
	case lang_Spanish:
		return helloPrefix_Spanish
	}
	return helloPrefix;
}

func main()  {
	fmt.Println(Hello("World", ""))
}