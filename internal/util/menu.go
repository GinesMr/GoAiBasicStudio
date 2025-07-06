package util

var options = []string{
	"Select model",
	"View models",
	"Install new model",
	"Manual",
}

func ReturnOptionsMenu() []string {
	return options[:]
}
