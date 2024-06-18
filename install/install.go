package install

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func Installdbuggen(path string) {
	log.Printf("Trying to clone into %v", path)
	err := clone(path)

	if err != nil { // if clone fails
		log.Println("Clone failed, trying to git pull")
		if err := pull(path); err != nil { // try to pull
			log.Printf("Pull failed. Deleting %v and trying to clone", path)
			os.RemoveAll(path)                  // if pull fails then remove the directory
			if err := clone(path); err != nil { // and try to clone again
				log.Fatal(err) // womp womp.
			}
		}
	}
}

func clone(path string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/datasektionen/dbuggen",
		Progress: os.Stdout,
	})

	return err
}

func pull(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	if err := w.Pull(&git.PullOptions{
		Progress: os.Stdout,
	}); err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}
