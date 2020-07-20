# gitlab-tools

Tooling to help out with our usage of Gitlab

## Table of contents

- [Usage](#usage)
- [CI Usage](#ci)
- [Contributing](#contributing)

## Usage

Gitlab-tools offers the following commands:

### is-latest-commit

is-latest-commit is used to compare a given commit against a given branch to verify that the commit is truly the latest compared to the Github API. This allows the user to fail pipelines that may be redundant if needed.

> Usage: `gitlab-tools is-latest-commit`

#### Configuration

| Flag         | Env Variable       | Required | Description                                                                      | From version |
| ------------ | ------------------ | -------- | -------------------------------------------------------------------------------- | ------------ |
| github-token | GITHUB_TOKEN       | true     | Github token with API permissions                                                | v0.1.0       |
| repository   | GITHUB_REPOSITORY  | true     | Github repository name in format owner/repo                                      | v0.1.0       |
| reference    | CI_COMMIT_REF_NAME | true     | Branch name to check against                                                     | v0.1.0       |
| commit       | CI_COMMIT_SHA      | true     | Full commit SHA which will be used to compare against the latest on given branch | v0.1.0       |

## CI

To use in Gitlab you can use the following template:

```yaml
check-branch:
    stage: check-branch # a stage which preceeds an expensive pipeline
    image: alpine:latest # any image can be used
    variables:
        GIT_STRATEGY: none # Git is not needed
        GITHUB_REPOSITORY: productboardlabs/gitlab-tools # repo name in owner/repo format
        GITHUB_TOKEN: sometoken
    cache: {}
    rules:
        # run only on branches. Can be set up for master as well
        - if: '$CI_COMMIT_BRANCH != "master"'
        when: on_success
        - when: never
    before_script:
        - wget https://raw.githubusercontent.com/productboardlabs/gitlab-tools/master/godownloader-gitlab-tools.sh -O - | sh
    script:
        - ./bin/gitlab-tools is-latest-commit
```

## Contributing

This repo uses https://magefile.org/ to handle running command. Please run `go install github.com/magefile/mage` in order to use mage.

To test: `mage test`
