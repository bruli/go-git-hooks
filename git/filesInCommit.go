package git

import "os/exec"

func GetCommittedFiles()([]byte, error)  {
	return exec.Command("bash", "-c", "git diff-index --cached --name-status HEAD | egrep '^(A|M)' | awk '{print $2;}'").Output()
}