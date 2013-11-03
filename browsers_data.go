package bapp

var DefaultBrowser *Browser
var Chrome *Browser
var FirefoxTab *Browser

func init() {
	xdgOpen := &BrowserCmd{
		Cmd: "xdg-open",
	}
	darwinOpen := &BrowserCmd{
		Cmd: "open",
	}
	windowsOpen := &BrowserCmd{
		Cmd:  "cmd",
		Args: []string{"/c", "start"},
	}
	DefaultBrowser = &Browser{
		OS: map[string][]*BrowserCmd{
			"linux":   {xdgOpen},
			"freebsd": {xdgOpen},
			"netbsd":  {xdgOpen},
			"openbsd": {xdgOpen},
			"windows": {windowsOpen},
			"darwin":  {darwinOpen},
		},
	}

	googleChrome := &BrowserCmd{
		Cmd: "google-chrome",
	}
	Chrome = &Browser{
		OS: map[string][]*BrowserCmd{
			"linux":   {googleChrome},
			"freebsd": {googleChrome},
			"netbsd":  {googleChrome},
			"openbsd": {googleChrome},
			"windows": {googleChrome},
			"darwin":  {googleChrome},
		},
	}

	firefoxTab := &BrowserCmd{
		Cmd:  "firefox",
		Args: []string{"-new-tab"},
	}
	FirefoxTab = &Browser{
		OS: map[string][]*BrowserCmd{
			"linux":   {firefoxTab},
			"freebsd": {firefoxTab},
			"netbsd":  {firefoxTab},
			"openbsd": {firefoxTab},
			"windows": {firefoxTab},
			"darwin":  {firefoxTab},
		},
	}
}
