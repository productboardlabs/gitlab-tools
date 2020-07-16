# gitlab-tools

Tooling to help out with our usage of Gitlab

## Table of contents

* [Usage](#usage)
* [Contributing](#contributing)

## Usage

Gitlab-tools offers the following commands:

### is-latest-commit

is-latest-commit is used to compare a given commit against a given branch to verify that the commit is truly the latest compared to the Github API. This allows the user to fail pipelines that may be redundant if needed.

> Usage: `gitlab-tools is-latest-commit`

#### Configuration

| Flag | Env Variable | Required | Description | From version |
| --|--|--|--|--|
| github-token | GITHUB_TOKEN | true | Github token with API permissions | v0.1.0 |
| repository | GITHUB_REPOSITORY | true | Github repository name in format owner/repo | v0.1.0 |
| reference | CI_COMMIT_REF_NAME | true | Branch name to check against |  v0.1.0 |
| commit | CI_COMMIT_SHA | true | Full commit SHA which will be used to compare against the latest on given branch | v0.1.0 |

## Contributing

This repo uses https://magefile.org/ to handle running command. Please run `go install github.com/magefile/mage` in order to use mage.

To test: `mage test`
