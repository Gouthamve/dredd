package dredd

// Judge is the interface for the online Judge
type Judge interface {
	SetLimits(Limits) error
	Judge(file, inp, output string) (bool, error)
}

// Limits are the constraints on the execution
type Limits struct {
	Memory int `json:"memory,omitempty"` // Memory in MB
	Time   int `json:"time,omitempty"`   // time in secs
}

// Problem is self-explanatory
type Problem struct {
	Lang      string     `json:"lang"`
	Limits    Limits     `json:"limits,omitempty"`
	Testcases []Testcase `json:"testcases"`
}

// Testcase is self-explanatory
type Testcase struct {
	Inp      string `json:"inp"`
	Expected string `json:"expected"`
}

// Result is the output of each testcase run
type Result struct {
	Ok       bool   `json:"ok,omitempty"`
	Err      string `json:"err,omitempty"`
	Exp      string `json:"exp"`
	Got      string `json:"got,omitempty"`
	ExitCode int    `json:"exitCode"`

	Flag int `json:"flag"`
}

// Compiler interface is for the compilers
type Compiler interface {
	Compile(lang, path string) (string, error)
}
