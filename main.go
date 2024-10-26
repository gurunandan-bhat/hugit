package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	wktree, err := OpenIfClean("/Users/nandan/repos/hugot")
	if err != nil {
		log.Fatal("error opening an existing repo: ", err)
	}

	fmt.Printf("%#v", wktree)
}

func OpenIfClean(path string) (*git.Worktree, error) {

	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	wktree, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	status, err := wktree.Status()
	if err != nil {
		return nil, err
	}

	isClean := status.IsClean()
	if !isClean {
		err := fmt.Errorf("unclean work tree: %s", status.String())
		return nil, err
	}

	branchIter, err := r.Branches()
	if err != nil {
		return nil, err
	}

	err = branchIter.ForEach(func(ref *plumbing.Reference) error {
		println("Ref: ", ref.Hash().String())
		return nil
	})

	return wktree, nil
}
