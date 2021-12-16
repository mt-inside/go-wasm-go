package runner

import (
	"path/filepath"
	"runtime"
	"testing"
)

const (
	goPluginSo                = "./provider/plugin/ok/plugin.so"
	goPluginSoThird           = "./provider/plugin/third/plugin.so"
	goPluginSoThirdDiffModVer = "./provider/plugin/third_diff_mod_ver/plugin.so"
	goPluginSo_1_15           = "./provider/plugin/ok/plugin-1.15.so"
	goPluginSo_1_16           = "./provider/plugin/ok/plugin-1.16.so"
	goPluginSo_1_17_1         = "./provider/plugin/ok/plugin-1.17.1.so"

	wasmTinygo = "./provider/wasm-tinygo/wasm.wasm"
	wasmGo     = "./provider/wasm-go/wasm.wasm"
)

var fbTests = []fbTestItem{
	{name: "5", in: 5, want: 5},
	{name: "10", in: 10, want: 55},
	{name: "20", in: 20, want: 6765},
	{name: "30", in: 30, want: 832040},
	// {name: "40", in: 40, want: 102334155},
}

type fbTestItem struct {
	name string
	in   int32
	want int32
}

// selfDir get current test file dir path
func selfDir(t testing.TB) string {
	t.Helper()

	// nolint: dogsled
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}
