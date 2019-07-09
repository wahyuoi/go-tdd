package helloworld

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

func Hello(name, language string) string {

  prefix := getPrefix(language)
  if name == "" {
    name = "World"
  }
  return prefix + name
}

func getPrefix(language string) string {
  switch language {
  case "spanish":
    return spanishPrefix
  case "french":
    return frenchPrefix
  default:
    return englishPrefix
  }
}
