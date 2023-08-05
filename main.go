package main

import (
	"flag"
	"fmt"
	lorem "github.com/drhodes/golorem"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/lmittmann/tint"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Options struct {
	Repo          string
	UserName      string
	AccessToken   string
	EmailAddress  string
	CommitsPerDay int
	WorkDaysOnly  bool
	StartDate     string
	EndDate       string
}

var (
	logger = slog.New(tint.NewHandler(os.Stderr, nil))
)

func main() {
	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelError,
			TimeFormat: time.Kitchen,
		}),
	))

	opts := new(Options)

	flag.IntVar(&opts.CommitsPerDay, "commitsPerDay", 10, "number of commits per day")
	flag.BoolVar(&opts.WorkDaysOnly, "workdaysOnly", false, "only workdays")
	flag.StringVar(&opts.StartDate, "startDate", "", "start date")
	flag.StringVar(&opts.EndDate, "endDate", "", "end date")
	flag.StringVar(&opts.Repo, "repo", "", "your private repo url (e.g. https://github.com/navicstein/fake-contributions-repo.git")

	flag.StringVar(&opts.UserName, "username", "navicstein", "your github username")
	flag.StringVar(&opts.AccessToken, "accessToken", "", "your github access token used for pushing & cloning private repos")
	flag.StringVar(&opts.EmailAddress, "emailAddress", "navicsteinrotciv@gmail.com", "your commit email address")

	flag.Parse()

	if err := runFakeContributions(opts); err != nil {
		panic(err)
	}

	logger.Info("If you're using this tool, please consider giving me a star on GitHub, I'll appreciate it :)")
}

func runFakeContributions(opts *Options) (err error) {
	var (
		tmpFolder = "tmp-contributions-repo"

		commitsPerDay = opts.CommitsPerDay
		workdaysOnly  = opts.WorkDaysOnly
		startDateStr  = opts.StartDate
		endDateStr    = opts.EndDate
	)

	commitsPerDayIntervals, err := parseCommitsPerDay(commitsPerDay)
	if err != nil {
		log.Panicf("Error parsing commitsPerDay: %s", err)
		return
	}

	startDate := parseDateOrDefault(startDateStr, time.Now().AddDate(-1, 0, 0))
	endDate := parseDateOrDefault(endDateStr, time.Now())

	commitDateList := createCommitDateList(commitsPerDayIntervals, workdaysOnly, startDate, endDate)

	// Remove git history folder in case if it already exists.
	if _, err := os.Stat(tmpFolder); !os.IsNotExist(err) {
		if err := os.RemoveAll(tmpFolder); err != nil {
			return err
		}
	}

	// Create git history folder.
	_ = os.Mkdir(tmpFolder, os.ModePerm)

	// Clone the repository
	logger.Info("Cloning repository into", "path", tmpFolder, "cloneUrl", opts.Repo)

	repo, err := git.PlainClone(tmpFolder, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: opts.UserName,
			Password: opts.AccessToken,
		},
		Progress: os.Stdout,
		URL:      opts.Repo,
	})
	if err != nil {
		return
	}

	// Create commits.
	for _, date := range commitDateList {
		dateFormatted := date.Format("January 2, 2006")
		msg := lorem.Sentence(4, 8)

		w, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("unable to create worktree: %+v", err)
		}

		_, err = w.Commit(msg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  opts.UserName,
				Email: opts.EmailAddress,
				When:  date,
			},
		})
		if err != nil {
			return fmt.Errorf("unable to commit to the repository: %+v", err)
		}

		logger.Debug("Generating your GitHub activity", "date", dateFormatted, "message", msg)
	}

	logger.Info("Total commits have been created:", "length", len(commitDateList))

	// Push commits to the repository.
	logger.Info("Pushing commits to the repository: ", "repo", opts.Repo)
	pushOpts := git.PushOptions{
		Auth: &http.BasicAuth{
			Username: opts.UserName,
			Password: opts.AccessToken,
		},
		Progress: os.Stdout,
	}

	if err := repo.Push(&pushOpts); err != nil {
		return fmt.Errorf("unable to push to the repository: %+v", err)
	}

	return nil
}
