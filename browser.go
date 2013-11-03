package bapp

// Browser defines the possible browser commands and args per Operating System
type Browser struct {
	OS map[string][]*BrowserCmd // map OS maps OS name to a list of browser options
}

// BrowserCmd defines the command and arguments to start a browser for a given OS
type BrowserCmd struct {
	Cmd  string   // command name
	Args []string // arguments to open the browser correctly
}

//++ `BrowserCmd.Exists() bool` (use os/exec.LookPath(..))
