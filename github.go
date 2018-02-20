package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func newGithubClient(ctx context.Context, baseURL string, insecureSkipVerify bool) (*github.Client, error) {
	ghToken, err := getGithubToken()
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: ghToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	if tct, ok := tc.Transport.(*oauth2.Transport); ok && insecureSkipVerify {
		tct.Base = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	if baseURL == defaultBaseURL {
		client := github.NewClient(tc)
		return client, nil
	}

	client, err := github.NewEnterpriseClient(baseURL, baseURL, tc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getGithubToken() (string, error) {
	token := os.Getenv("GITHUB_API_TOKEN")
	if token == "" {
		return "", fmt.Errorf("GITHUB_API_TOKEN must be provided")
	}
	return token, nil
}
