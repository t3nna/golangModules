package golangModules

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "Bark"
}

func Barks() string {
	return Bark() + " " + Bark()
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func ILikeGo() string {
	return "I like Go lang"
}
