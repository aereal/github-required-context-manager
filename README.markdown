# github-required-context-manager

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
