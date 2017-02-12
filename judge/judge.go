package judge

import "github.com/juju/errors"

// Dredd is the default implementation
// of Judge
type Dredd struct {
	limits Limits
}

// NewDefaultDredd returns a default dredd
func NewDefaultDredd() *Dredd {
	return &Dredd{
		limits: Limits{
			memory: 300,
			time:   10,
		},
	}
}

// SetLimits sets the program limits
func (d *Dredd) SetLimits(l Limits) {
	d.limits = l
}

// Judge judges the testcase
func (d *Dredd) Judge(file, inp, output string) (bool, error) {
	c := sandbox.Config{
		[]string{file},
		io.NewReader(inp),
		d.limits.memory * 1024,
		d.limits.time * 1000,
	}
	s, err := sandbox.NewStdSandbox(c)
	if err != nil {
		return false, errors.Annotate(err, "failed to create new sandbox")
	}
	out, err := s.Run()
	if err != nil {
		return false, errors.Trace(err)
	}

	// TODO: Compare output
	return true, nil
}
