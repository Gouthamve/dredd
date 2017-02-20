package judge

import (
	"io/ioutil"

	sandbox "github.com/ggaaooppeenngg/libsandbox"
	"github.com/gouthamve/dredd"
	"github.com/juju/errors"
	minio "github.com/minio/minio-go"
	"github.com/spf13/viper"
)

// RunnerArgs encapsulates the arguments for Runner
type RunnerArgs struct {
	Problem  dredd.Problem `json:"problem"`
	Filename string        `json:"filename"`
}

// Runner is the runner that uses Judge :P
// TODO: Better desc
type Runner struct {
	j            dredd.Judge
	p            dredd.Problem
	file         string
	localFile    string
	compiledFile string
}

// NewRunner is self-explanatory
func NewRunner(ra RunnerArgs) (Runner, error) {
	if err := validateArgs(ra); err != nil {
		return Runner{}, errors.Trace(err)
	}

	r := Runner{
		j: &Dredd{
			limits: ra.Problem.Limits,
		},
		p:    ra.Problem,
		file: ra.Filename,
	}

	lFile, err := downloadFile(ra.Filename)
	if err != nil {
		return Runner{}, errors.Trace(err)
	}
	r.localFile = lFile

	return r, nil
}

// Run runs the code against the testcases
// and returns the result
func (r Runner) Run() ([]dredd.Result, error) {
	// Compile
	out, err := r.Compile(r.p.Lang, r.localFile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	r.compiledFile = out

	// Execute the testcases
	res := make([]dredd.Result, len(r.p.Testcases))
	for i, v := range r.p.Testcases {
		correct, err := r.j.Judge(r.compiledFile, v.Inp, v.Expected)
		if err != nil {
			if err == sandbox.OutOfTimeError {
				res[i] = newResult(false, err.Error(), v.Expected, "", -1, 1)
				continue
			}

			if err == sandbox.OutOfMemoryError {
				res[i] = newResult(false, err.Error(), v.Expected, "", -1, 2)
				continue
			}

			// TODO: Expose more errors
			res[i] = newResult(false, err.Error(), v.Expected, "", -1, 3)
			continue
		}

		if correct {
			res[i] = newResult(true, "", v.Expected, v.Expected, 0, 0)
			continue
		}

		// TODO: Fix output
		res[i] = newResult(false, "", v.Expected, "", 0, 0)
	}

	return res, nil
}

func downloadFile(remoteFile string) (string, error) {
	endpoint := viper.GetString("minio-endpoint")
	accessKeyID := viper.GetString("minio-access")
	secretAccessKey := viper.GetString("minio-secret")

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return "", errors.Annotate(err, "could not create client")
	}

	file, err := ioutil.TempFile("", "dredd")
	if err != nil {
		return "", errors.Trace(err)
	}
	file.Close()

	// TODO: Make bucket configurable
	err = minioClient.FGetObject("submissions", remoteFile, file.Name())
	if err != nil {
		return "", errors.Annotate(err, "cannot download object")
	}

	return file.Name(), nil
}

func newResult(ok bool, err string, exp string, got string, code int, f int) dredd.Result {
	return dredd.Result{
		Ok:       ok,
		Err:      err,
		Exp:      exp,
		Got:      got,
		ExitCode: code,
		Flag:     f,
	}
}

func validateArgs(a RunnerArgs) error {
	// validate problem
	if a.Problem.Lang == "" {
		return errors.New("Lang is a mandatory argument")
	}

	if len(a.Problem.Testcases) == 0 {
		return errors.New("Atleast 1 testcase has to be provided")
	}

	if a.Problem.Limits.Memory == 0 {
		return errors.New("Memory limit should be non-zero")
	}

	if a.Problem.Limits.Time == 0 {
		return errors.New("Time limit should be non-zero")
	}

	// validate file
	if a.Filename == "" {
		return errors.New("Remote filename is a mandatory argument")
	}

	return nil
}
