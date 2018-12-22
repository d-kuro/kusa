package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/d-kuro/kusa/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

const layout = "2006-01-02"

func init() {
	now := time.Now()
	kusaCmd.PersistentFlags().StringVarP(
		&date, "date", "d", now.Format(layout), "date [format: yyyy-mm-dd]")
	kusaCmd.PersistentFlags().StringVarP(
		&repoDir, "repo", "r", "",
		"local directory path for clone GitHub repository (required)")
	kusaCmd.PersistentFlags().StringVarP(
		&commitMsg, "commit", "c", ":herb:", "commit message")
	kusaCmd.PersistentFlags().StringVarP(
		&name, "name", "n", "kusa", "commit author name")
	kusaCmd.PersistentFlags().StringVarP(
		&mail, "mail", "m", "kusa@example.com", "commit author mail address")
	rootCmd.AddCommand(kusaCmd)
}

var (
	repoDir   string
	date      string
	commitMsg string
	name      string
	mail      string
	kusaCmd   = &cobra.Command{
		Use:   "create",
		Short: "Create GitHub contribution",
		Long:  "Create GitHub contribution on date specified by date option",
		Run: func(cmd *cobra.Command, args []string) {
			if err := createKusa(); err != nil {
				log.Error("create GitHub contribution error", zap.Error(err))
			}
		},
	}
)

func createKusa() error {
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		log.Error("open git repository error", zap.String("dir_path", repoDir), zap.Error(err))
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Error("failed to get work tree", zap.Error(err))
		return err
	}

	time, err := time.Parse(layout, date)
	if err != nil {
		log.Error("time parse error", zap.String("date", date), zap.Error(err))
		return err
	}

	log.Info("execute commit",
		zap.String("mame", name), zap.String("e-mail", mail),
		zap.String("date", date), zap.String("commit_message", commitMsg))
	commit, err := wt.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: mail,
			When:  time,
		},
	})
	if err != nil {
		log.Error("commit error", zap.Error(err))
		return err
	}
	log.Info("complete commit", zap.String("commit_hash", commit.String()))

	log.Info("input credential")
	auth, err := inputCredentials()
	if err != nil {
		log.Error("failed read credentials", zap.Error(err))
	}

	log.Info("execute push", zap.String("repository", repoDir))
	if err := repo.Push(&git.PushOptions{
		Auth:     auth,
		Progress: os.Stdout,
	}); err != nil {
		log.Error("push error", zap.Error(err))
		return err
	}
	log.Info("complete push")

	return nil
}

func inputCredentials() (transport.AuthMethod, error) {
	fmt.Print("user name: ")
	user, err := terminal.ReadPassword(syscall.Stdin)
	// new line
	fmt.Println()
	if err != nil {
		log.Error("failed read user name", zap.Error(err))
		return nil, err
	}

	fmt.Print("password: ")
	pass, err := terminal.ReadPassword(syscall.Stdin)
	// new line
	fmt.Println()
	if err != nil {
		log.Error("failed read password", zap.Error(err))
		return nil, err
	}

	return &http.BasicAuth{
		Username: string(user),
		Password: string(pass),
	}, nil
}
