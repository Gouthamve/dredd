package dredd

// Judge is the interface for the online Judge
type Judge interface {
	SetLimits(limits) error
	Judge(file, inp, output string) (bool, error)
}

// Database is the database interface
type Database interface {
	GetProblem(string) (problem, error)
}

// Limits are the constraints on the execution
type Limits struct {
	memory int // Memory in MB
	time   int // time in secs
}

// Problem is self-explanatory
type Problem struct {
	id        string
	question  string
	limits    limits
	testcases []testcase
}

// Testcase is self-explanatory
type Testcase struct {
	inp      string
	expected string
}

// Compiler interface is for the compilers
type Compiler interface {
	Compile(lang, path string) (string, error)
}
