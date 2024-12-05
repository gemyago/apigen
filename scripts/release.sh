#!/usr/bin/env bash

set -e

# This script resolves the release tag for the current branch.
# If the branch is in form release/ or hotfix/, the release tag is the APP_VERSION from .versions file.
# In other case the release will be rc-<short-hash>.

usage() {
  echo "Usage: $0 <command>"
  echo "Commands:"
  echo "  resolve-release-tag - Resolves the release tag for the current branch."
  exit 1
}

command=$1

if [ -z "$command" ]; then
  usage
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

case $command in
  resolve-release-tag)
    # If GITHUB_HEAD_REF is set, we take the branch name from it.
    # otherwise we take the branch name from the current branch. 
    branch_name=${GITHUB_HEAD_REF:-$(git rev-parse --abbrev-ref HEAD)}

    if [[ $branch_name =~ ^release/.*$ || $branch_name =~ ^hotfix/.*$ ]]; then
      release_tag=$(grep 'APP_VERSION' ${SCRIPT_DIR}/../.versions | cut -d ' ' -f 2)
    else
      short_hash=$(git rev-parse --short HEAD)
      release_tag="rc-$short_hash"
    fi

    echo $release_tag
    ;;
  *)
    echo "Unknown command: $command"
    usage
    ;;
esac