// Package testutil is utility for test
package testutil

import (
	"flag"
	"fmt"
	"os"
	"testing"

	enc "github.com/hiromaily/golibs/cipher/encryption"
	conf "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	r "github.com/hiromaily/golibs/runtimes"
)

var (
	//LogFlg is for switch to output log or not
	LogFlg = flag.Uint("log", 1, "Log Flg: 0:OFF, 1:ON")
	//ConfFile is toml file path
	ConfFile = flag.String("fp", "", "Config File Path")
	//JSONFile is json file path
	JSONFile = flag.String("jfp", "", "Json File Path")
	//YAMLFile is json file path
	YAMLFile = flag.String("yfp", "", "YAML File Path")

	//KafkaIP is IP for kafka server
	KafkaIP = flag.Int("kip", 9092, "Json File Path")

	//BenchFlg is when benchmark test, value is true
	BenchFlg = false
)

// InitializeTest is to run common initial code for test
func InitializeTest(prefix string) {
	flag.Parse()

	//log
	//lg.InitializeLog(lg.LogStatus(*LogFlg), lg.NoDateNoFile, prefix, "/var/log/go/test.log", "hiromaily")
	lg.InitializeLog(lg.LogStatus(*LogFlg), lg.NoDateNoFile, prefix, "", "hiromaily")

	//-v : to show Logs.(-test.v=true)
	if o.FindParam("-test.v") {
		lg.Debug("This test can show log in detail.")
	}

	//crypt
	_, err := enc.NewCryptWithEnv()
	if err != nil {
		panic(err)
	}

	//conf
	if *ConfFile == "" {
		*ConfFile = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.toml"
	}
	conf.New(*ConfFile, true)

	//json
	if *JSONFile == "" {
		//default
		*JSONFile = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/testdata/json/teachers.json"
	}

	//bench
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		BenchFlg = true
	}
}

// Logf is t.Logf()
//	tu.Logf(t, "%+v", mRet)
func Logf(t *testing.T, format string, args ...interface{}) {
	if *LogFlg == 1 {
		t.Logf(format, args...)

	}
}

// Log is t.Log()
//	tu.Log(t, mRet)
func Log(t *testing.T, args ...interface{}) {
	if *LogFlg == 1 {
		t.Log(args...)
	}
}

// SkipLog is to skip test with func name
func SkipLog(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(2)))
}

// SkipBLog is to skip test with func name
func SkipBLog(b *testing.B) {
	//b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(2)))
}
