package main

import (
	"flag"
	"fmt"
)

var (
	keepAlive        int
	connectTimeout   int
	ackTimeout       int
	timeoutRetries   int
	authenticator    string
	sessionsProvider string
	topicsProvider   string
	cpuprofile       string
)

const (
	DefaultKeepAlive        = 300
	DefaultConnectTimeout   = 2
	DefaultAckTimeout       = 20
	DefaultTimeoutRetries   = 3
	DefaultSessionsProvider = "mem"
	DefaultAuthenticator    = "mockSuccess"
	DefaultTopicsProvider   = "mem"
)

func init() {
	flag.IntVar(&keepAlive, "keepalive", DefaultKeepAlive, "Keepalive (sec)")
	flag.IntVar(&connectTimeout, "connecttimeout", DefaultConnectTimeout, "Connect Timeout (sec)")
	flag.IntVar(&ackTimeout, "acktimeout", DefaultAckTimeout, "Ack Timeout (sec)")
	flag.IntVar(&timeoutRetries, "retries", DefaultTimeoutRetries, "Timeout Retries")
	flag.StringVar(&authenticator, "auth", DefaultAuthenticator, "Authenticator Type")
	flag.StringVar(&sessionsProvider, "sessions", DefaultSessionsProvider, "Session Provider Type")
	flag.StringVar(&topicsProvider, "topics", DefaultTopicsProvider, "Topics Provider Type")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "CPU Profile Filename")
	flag.Parse()
}

/**
 * flag为golang用于提取命令行参数的一个库
 */
func main() {

	fmt.Println(keepAlive)
}

func FlagXXX() {

	data_path := flag.String("D", "/home/manu/sample/", "DB data path")
	log_file := flag.String("l", "/home/manu/sample.log", "log file")
	nowait_flag := flag.Bool("W", false, "do not wait until operation completes")

	flag.Parse()

	var cmd string = flag.Arg(0)

	fmt.Printf("action   : %s\n", cmd)
	fmt.Printf("data path: %s\n", *data_path)
	fmt.Printf("log file : %s\n", *log_file)
	fmt.Printf("nowait     : %v\n", *nowait_flag)

	fmt.Printf("-------------------------------------------------------\n")

	fmt.Printf("there are %d non-flag input param\n", flag.NArg())
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}
}
