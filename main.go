package main

import (
	"os"
	"os/exec"
	"golang.org/x/sys/unix"
)

// コマンドライン引数の処理
// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		execute(os.Args[2], os.Args[3:]...)
	default:
		panic("コマンドライン引数が正しくありません。")
	}
}

// 隔離されたプロセスの中で cmd を引数 arg と共に実行
func execute(cmd string, args ...string) {
	// ルートディレクトリとカレントディレクトリを ./rootfs に設定
	must(unix.Chroot("./rootfs"))
	must(unix.Chdir("/"))

	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	must(command.Run())
}

// 雑にエラー処理
func must(err error) {
	if err != nil {
		panic(err)
	}
}
