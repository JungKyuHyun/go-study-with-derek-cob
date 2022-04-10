package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"
const usage = `Usage:
%s [command]
Commands:
    Greet
    Version
`

const greetUsage = `Usage:
%s greet name [flag]
Positional Arguments:
    name
        the name to greet
Flags:
`

type MenuConf struct {
	Goodbye bool
}

// 메뉴를 세팅
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	// Usage : flags 파싱하는 동안 에러가 나는지 체크하는 함수.
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		// PrintDefaults:  prints default values of all defined command-line flags in the set.
		menu.PrintDefaults()
	}
	return menu
}

// subMemu를 가져옴
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	// BoolVar : defines a bool flag with specified name, default value, and usage string.
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")

	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}

// m 값에 따라 출력
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Println("Goodbye " + name + "!")
	} else {
		fmt.Println("Hello " + name + "!")
	}
}

func (m *MenuConf) Version() {
	fmt.Println("Version: " + version)
}
