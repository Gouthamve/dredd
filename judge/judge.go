package judge

import (
	"strings"

	sandbox "github.com/ggaaooppeenngg/libsandbox"
	"github.com/gouthamve/dredd"
	"github.com/juju/errors"
)

// Dredd is the default implementation
// of Judge
type Dredd struct {
	limits dredd.Limits
}

// NewDefaultDredd returns a default dredd
func NewDefaultDredd() *Dredd {
	return &Dredd{
		limits: dredd.Limits{
			Memory: 300,
			Time:   10,
		},
	}
}

// SetLimits sets the program limits
func (d *Dredd) SetLimits(l dredd.Limits) error {
	d.limits = l
	return nil
}

// Judge judges the testcase
func (d *Dredd) Judge(file, inp, output string) (bool, error) {
	c := sandbox.Config{
		Args:  []string{file},
		Input: strings.NewReader(inp),
		//TODO: Memory calculation screwd up?
		Memory: int64(d.limits.Memory * 1024),
		Time:   int64(d.limits.Time * 1000),
	}

	s, err := sandbox.NewStdSandbox(c)
	if err != nil {
		return false, errors.Annotate(err, "failed to create new sandbox")
	}
	_, err = s.Run()
	if err != nil {
		return false, err
	}

	// TODO: Compare output
	return true, nil
}
