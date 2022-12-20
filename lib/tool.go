package lib

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetProgramPath() (string, error) {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex), err
	}

	exReal, err := filepath.EvalSymlinks(ex)
	return filepath.Dir(exReal), err
}

func ExecGetSysInfoStdout(pathBin string, args ...string) ([]byte, error) {
	emptyEnv := make([]string, 0)
	return ExecGetCmdStdoutWithEnv(emptyEnv, pathBin, args...)
}

func ExecGetCmdStdoutWithEnv(appendEnv []string, pathBin string, args ...string) ([]byte, error) {
	env := os.Environ()
	cmd := exec.Command(pathBin, args...)
	// argArr := []string{"-c"}
	// argArr = append(argArr, pathBin)
	// argArr = append(argArr, args...)
	// cmd := exec.Command("bash", argArr...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Pdeathsig: syscall.SIGINT, //如果主进程退出，则将 SIGINT 发送给子进程
	// }
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.Env = append(env, appendEnv...)
	cmd.Stdin = os.Stdin
	return cmd.Output()
}
