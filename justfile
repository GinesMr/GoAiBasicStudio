set positional-arguments := true

default:
	just --list

@models:
	curl -s https://ollama.com/library | sed -n 's/.*href="\/library\/\([^"]*\)".*/\1/p' | sort -u

@tags model:
	curl -s https://ollama.com/library/{{model}}/tags | grep -o "{{model}}:[^\" ]*q[^\" ]*" | grep -E -v 'text|base|fp|q[45]_[01]'

