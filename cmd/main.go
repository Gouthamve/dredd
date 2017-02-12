package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gouthamve/dredd"
	"github.com/gouthamve/dredd/judge"
	"github.com/juju/errors"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("dredd")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	if len(os.Args) != 2 {
		log.Fatalln("Usage: dredd <args>")
	}

	a := args{}
	json.Unmarshal([]byte(os.Args[1]), &a)

	if err := validateArgs(a); err != nil {
		log.Fatalln(errors.ErrorStack(err))
	}

	r, err := judge.NewRunner(a.Problem, a.Filename)
	if err != nil {
		log.Fatalln(errors.ErrorStack(err))
	}

	res, err := r.Run()
	if err != nil {
		log.Fatalln(errors.ErrorStack(err))
	}

	byt, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(errors.Trace(err))
	}

	fmt.Println(string(byt))
}

type args struct {
	Problem  dredd.Problem `json:"problem"`
	Filename string        `json:"filename"`
}

func validateArgs(a args) error {
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

func validateTestcase(t dredd.Testcase) error {

	return nil
}
