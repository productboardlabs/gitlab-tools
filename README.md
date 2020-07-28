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

### set-status-check

set-status-check allows a user arbitrarily set a status check for a given commit. This allows us to mirror the job status to Github from Gitlab as by default the Gitlab <> Github integration only displays `ci/gitlab/gitlab.com` in the status section in Github. This leads to improved DX.

> Usage: `gitlab-tools set-status-check --status pending`

| Flag | Env Variable | Required | Description | From version |
| ---- | ------------ | -------- | ------------ | ----------- |
| github-token | GITHUB_TOKEN | true | Github token with API permissions | v0.2.0 |
| repository | GITHUB_REPOSITORY | true | Github repository name in format owner/repo | v0.2.0 |
| status | STATUS | true | Status to be set, allowed are: **error**, **failure**, **pending** or **success** | v0.2.0 |
| exit-code | none | false | Sets status based on a given error code **overrides status flag** | v0.2.0 |
| commit | CI_COMMIT_SHA | true | Commit sha for which to set the status | v0.2.0 |
| status-name | CI_JOB_STAGE/CI_JOB_NAME | true | Name of status that will be shown in Github | v0.2.0 |
| description | STATUS_DESCRIPTION | false | Extra information to show in Github. e.g. number of failed checks | v0.2.0 |
| url | CI_JOB_URL | false | Link to the given job in Gitlab | v0.2.0 |

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
