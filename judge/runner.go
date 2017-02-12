package judge

// Runner is the runner that uses Judge :P
// TODO: Better desc
type Runner struct {
	j    *Judge
	p    Problem
	file string
}

// NewRunner is self-explanatory
func NewRunner(p Problem, file string) Runner {
	return Runner{
		j: &Dredd{
			limits: p.limits,
		},
		p:    p,
		file: file,
	}
}

func (r Runner) Run() {

}
