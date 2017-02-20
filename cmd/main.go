package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gouthamve/dredd/judge"
	"github.com/juju/errors"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("dredd")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	ra := judge.RunnerArgs{}
	if err := json.NewDecoder(os.Stdin).Decode(&ra); err != nil {
		log.Fatalln(errors.ErrorStack(errors.Trace(err)))
	}

	r, err := judge.NewRunner(ra)
	if err != nil {
		log.Fatalln(errors.ErrorStack(errors.Trace(err)))

	}

	res, err := r.Run()
	if err != nil {
		log.Fatalln(errors.ErrorStack(errors.Trace(err)))
	}

	byt, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(errors.ErrorStack(errors.Trace(err)))
	}

	fmt.Println(string(byt))
}
