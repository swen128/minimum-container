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
		initialize(os.Args[2:]...)
	case "child":
		execute(os.Args[2], os.Args[3:]...)
	default:
		panic("コマンドライン引数が正しくありません。")
	}
}

// Linux namespace を設定した子プロセスで、execute 関数を実行する
func initialize(args ...string) {
	// このプログラム自身に引数 child <cmd> <args> を渡す
	arg := append([]string{"child"}, args...)
	command := exec.Command("/proc/self/exe", arg...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWNS | unix.CLONE_NEWPID | unix.CLONE_NEWUTS,
	}

	must(command.Run())
}

// namespace 設定後の初期化処理と、ユーザー指定のコマンドを実行する
func execute(cmd string, args ...string) {
	// ルートディレクトリとカレントディレクトリを ./rootfs に設定
	must(unix.Chroot("./rootfs"))
	must(unix.Chdir("/"))

	must(unix.Mount("proc", "proc", "proc", 0, ""))
	must(unix.Sethostname([]byte("my-container")))

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
