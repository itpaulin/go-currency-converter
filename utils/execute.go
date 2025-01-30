package utils

func Execute(funcs ...func()) {
	for _, function := range funcs {
		function()
	}
}