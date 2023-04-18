package git

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"
)

// To Check if the given repository is Public(No Authentication needed), send a HTTP GET request to the URL
// If response code is 200, the repository is Public.
func IsGitRepoPublic(gitURL string) bool {
	resp, err := http.Get(gitURL)

	if err != nil {
		return false
	}
	// if the status code is 200, our get request is successful.
	// It only happens when the repository is public.
	if resp.StatusCode == 200 {
		return true
	}

	return false
}

// Check if the GITHUB_TOKEN is present
func GetGitHubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func Clone(path string, fs billy.Filesystem, branch string, auth transport.AuthMethod) (*git.Repository, error) {
	return git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           path,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Progress:      os.Stdout,
		SingleBranch:  true,
		Depth:         1,
		Auth:          auth,
	})
}

func ListFiles(fs billy.Filesystem, path string, predicate func(fs.FileInfo) bool) ([]string, error) {
	path = filepath.Clean(path)
	if _, err := fs.Stat(path); err != nil {
		return nil, err
	}
	files, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var results []string
	for _, file := range files {
		name := filepath.Join(path, file.Name())
		if file.IsDir() {
			children, err := ListFiles(fs, name, predicate)
			if err != nil {
				return nil, err
			}
			results = append(results, children...)
		} else if predicate(file) {
			results = append(results, name)
		}
	}
	return results, nil
}

func ListYamls(f billy.Filesystem, path string) ([]string, error) {
	return ListFiles(f, path, IsYaml)
}
