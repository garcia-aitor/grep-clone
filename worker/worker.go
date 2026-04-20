package worker

type Result struct {
	Line string
	LineNum int
	Path string
}

type Results struct {
	Inner []Result
}

func NewResult(line string, lineNum int, path string) Result {
	return Result{line, lineNum, path}
}

func FindInFile(path string, find string) * Results {
	file, err := os.Open(path)
	
}