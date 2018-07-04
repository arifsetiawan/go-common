package open

import (
	"os/exec"
	"runtime"
)

// OpenBrowser is
func OpenBrowser(url string) error {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start()
}

/*
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*

func openBrowser() {
	openBrowser := env.Getenv("OPEN_BROWSER", "false")
	if openBrowser == "true" {
		appURL := env.Getenv("BASE_URL", "http://localhost:7000")
		nfLog.Info("Opening default browser to " + appURL)
		err := open.OpenBrowser(appURL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

*/
