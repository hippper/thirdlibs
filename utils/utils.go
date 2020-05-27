package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"

	. "tl-gitlab/base/logger"
)

func ASSERT(exp bool, info ...string) {
	if !exp {
		infostr := ""
		if len(info) > 0 {
			infostr = info[0]
		}
		Log.Errorf("ASSERT FAILED!\ninfo=[%v]\nstack = [%v]\n", infostr, string(debug.Stack()))
		panic("ASSERT FAILED")
	}
}

func CatchPanic() {
	if err := recover(); err != nil {
		Log.Errorf("panic !!! err = %v ", err)
	}
}

func CatchException() {
	if err := recover(); err != nil {
		fullPath, _ := exec.LookPath(os.Args[0])
		fname := filepath.Base(fullPath)

		datestr := NowDateStr()
		outstr := fmt.Sprintf("\n======\n[%v] err=%v, stack=%v\n======\n", time.Now(), err, string(debug.Stack()))
		filename := "./log/panic_" + fname + datestr + ".log"
		f, err2 := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		ASSERT(err2 == nil)
		defer f.Close()
		f.WriteString(outstr)

		Log.Errorf("err = %v ", err)
	}
}

func UseMaxCpu() {
	// multiple cpus using
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func FormatJsonStr(instr string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(instr), "", "  ")

	return "\n" + out.String() + "\n"
}

func FormatStruct(inst interface{}) string {
	instr := SerializeToJson(inst)
	return FormatJsonStr(instr)
}

func SerializeToJson(st interface{}) string {
	ba, _ := json.Marshal(st)
	jsonstr := string(ba)

	return jsonstr
}
