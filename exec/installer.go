package exec

import (
	"fmt"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
)

func Install(conf *config.Configuration, repo *git.Repository) {
	fmt.Println("Installing CaptainHook START")
	fmt.Println(conf.Path())
	fmt.Println(repo.Path())
	fmt.Println("Installing CaptainHook END")
}
