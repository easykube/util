package util

import (
	"bufio"
	"io"
	"os/exec"
	"strings"

	"github.com/axgle/mahonia"
)

//控制台字符集
func ConsoleCharset() string {
	return "gbk"
}

//执行命令输出回调
type CmdOutCallback func(line string)

//执行命令行
func ExecCmdCallback(callback CmdOutCallback, name string, args ...string) error {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = strings.Replace(line, "\r", "", -1)
		line = strings.Replace(line, "\n", "", -1)
		cs := ConsoleCharset()
		if cs != "" {
			line = mahonia.NewDecoder(cs).ConvertString(line)
		}
		if callback != nil {
			callback(line)
		}
	}

	cmd.Wait()
	return nil
}

//执行命令，将输出作为字符串返回
func ExecCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	lines := string(out)
	lines = strings.Replace(lines, "\r", "", -1)
	cs := ConsoleCharset()
	if cs != "" {
		lines = mahonia.NewDecoder(cs).ConvertString(lines)
	}
	return lines, err
	/**
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	cmd.Start()
	buf := bytes.Buffer{}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = strings.Replace(line, "\r", "", -1)
		line = mahonia.NewDecoder("gbk").ConvertString(line)
		println(line)
		buf.WriteString(line)
	}

	cmd.Wait()

	return buf.String(), err
	**/
}

//执行命令
func ExecCmdSimple(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

//执行命令行,参数使用空格分开
func ExecCmdLine(cmdLine string) (string, error) {
	name := ""
	param := SplitByTrim(cmdLine, " ")
	if len(param) > 0 {
		name = param[0]
	}
	if len(param) > 1 {
		args := param[1:]
		return ExecCmd(name, args...)
	} else {
		return ExecCmd(name)
	}
}

//按指定的分界符拆分并去掉空
func SplitByTrim(src string, comma string) []string {
	words := strings.Split(src, comma)
	ret := []string{}
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w != "" {
			ret = append(ret, w)
		}

	}
	return ret

}
