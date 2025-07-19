package util

var options = []string{
	"Install new model",
	"Use models",
	"Manual",
}

func ReturnOptionsMenu() []string {
	return options[:]
}
