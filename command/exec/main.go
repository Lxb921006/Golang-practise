package main

import (
	"log"
	"os/exec"
)

func main() {
	data, err := exec.Command("sh", "/web/wwwroot/shell/opt/cron_script/new_run_cron.sh", "101", "/usr/local/php/bin/php /web/wwwroot/shell/truco_cron/index_cli.php Shell/Gameonlinedata/index").Output()
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(string(data))
}
