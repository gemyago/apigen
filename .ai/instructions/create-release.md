# Instruction to prepare a new release

Release process involves updating version and creating a release PR by following steps below:

1. Get latest APP_VERSION from .versions file

2. Understand changes since last release:
```bash
git log --oneline --no-merges $(gh release list --limit 1 --json tagName --jq '.[0].tagName')..main
```

3. Based on the changes, determine the new version (major, minor, patch) according to semantic versioning. Show user the following and ask for confirmation (yes or ok):
  - Current version: X.Y.Z
  - Suggested new version: A.B.C
  - List of changes: 
    - Change 1
    - Change 2
    - ...

4. Once user confirms with "yes" or "ok" or provides alternative version: create release branch: `release/A.B.C`

5. Update version:
  - in `.versions` file and set APP_VERSION to new version
  - in `generators/go-apigen-server/pom.xml` file and set `<version>` to new version

6. Commit with message "Release A.B.C" 

7. Create PR
  ```bash
  gh pr create --title "Release A.B.C" --body "List of changes properly formatted" --base <base branch> --head <current branch>
  ```