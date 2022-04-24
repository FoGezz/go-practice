package task1

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const RemoteNtpServer = "0.beevik-ntp.pool.ntp.org"

/*
Hello now()
Завести Go репозиторий на GitHub, написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.
*/

func PrintTime() {
	time, err := ntp.Time(RemoteNtpServer)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Hello, now is %s\n", time.Format("01-02-2006 15:04:05"))
}
