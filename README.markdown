# github-required-context-manager [![Build Status](https://travis-ci.org/aereal/github-required-context-manager.svg?branch=master)](https://travis-ci.org/aereal/github-required-context-manager)

## Usage

### List required contexts

```
env GITHUB_API_TOKEN=... github-required-context-manager -owner aereal -repo github-required-context-manager -branch master -list
```

### Add a required context

```
env GITHUB_API_TOKEN=... github-required-context-manager -owner aereal -repo github-required-context-manager -branch master -add -context merge-blocker
```

### Delete the required context

```
env GITHUB_API_TOKEN=... github-required-context-manager -owner aereal -repo github-required-context-manager -branch master -delete -context merge-blocker
```
