#!/usr/bin/env bash

set -e

# This script resolves the release tag for the current branch.
# If the branch is in form release/v1.0.0, the release tag is v1.0.0.
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

case $command in
  resolve-release-tag)
    # If GITHUB_HEAD_REF is set, we take the branch name from it.
    # otherwise we take the branch name from the current branch. 
    branch_name=${GITHUB_HEAD_REF:-$(git rev-parse --abbrev-ref HEAD)}

    if [[ $branch_name =~ ^release/v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
      release_tag=${branch_name#release/}
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