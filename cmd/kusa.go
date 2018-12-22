package cmd

import (
	"time"

	"github.com/d-kuro/kusa/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const layout = "2006-01-02"

func init() {
	now := time.Now()
	kusaCmd.PersistentFlags().StringVarP(
		&date, "date", "d", now.Format(layout), "date [format: yyyy-mm-dd]")
	kusaCmd.PersistentFlags().StringVarP(
		&repoDir, "repo", "r", "", "local directory path for clone GitHub repository")
	kusaCmd.PersistentFlags().StringVarP(
		&commitMsg, "commit", "c", "草", "commit message")
	kusaCmd.PersistentFlags().StringVarP(
		&name, "name", "n", ":herb:", "commit author name")
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
		Long:  "kusa",
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
		zap.String("mame", name), zap.String("e-mail", mail), zap.String("date", date))

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
	log.Info("execute push", zap.String("repository", repoDir))

	if err := repo.Push(&git.PushOptions{}); err != nil {
		log.Error("push error", zap.Error(err))
		return err
	}

	log.Info("complete push")

	return nil
}
