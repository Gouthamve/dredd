package judge

import (
	"context"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/juju/errors"
)

const (
	// ErrCompileTimeout is compile timeout error
	ErrCompileTimeout = errors.New("compile timeout")
)

// Compile compiles the program
// This statisfies the compiler interface
func (r Runner) Compile(lang, path string) (output string, err error) {
	out, err := getTempFile(lang)
	if err != nil {
		return "", errors.Annotate(err, "Temp file creation failed")
	}

	switch lang {
	case "go":
		err = compileGo(path, out)
		if err != nil {
			return "", errors.Annotate("Compilation failed")
		}

		return out, nil
	default:
		return "", errors.Errorf("Unknown Language: %s", lang)
	}
}

func getTempFile(lang string) (string, error) {
	out, err := ioutil.TempFile("", lang)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func compileGo(path, output) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "build", "-o", output, path)
	out, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		return ErrCompileTimeout
	}

	if err != nil {
		return errors.New("Compilation failed: %s", out)
	}

	return nil
}
