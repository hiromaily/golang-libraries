package utils_test

import (
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/utils"
	"os"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Utils]")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCheckInterface(t *testing.T) {
	val1 := 10
	lg.Debug(CheckInterface(val1))

	val2 := "aaaaa"
	lg.Debug(CheckInterface(val2))

	//if err != nil {
	//}
}

func TestCheckInterfaceByIf(t *testing.T) {
	val1 := 10
	lg.Debug(CheckInterface(val1))

	val2 := "aaaaa"
	lg.Debug(CheckInterface(val2))
}

func TestSearchString(t *testing.T) {
	data := []string{"aaaa", "bbbb", "cccc", "dddd", "eeee"}
	ret := SearchString(data, "cccc")
	if ret != 2 {
		t.Errorf("SearchString is wrong: %d", ret)
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkUtils(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
