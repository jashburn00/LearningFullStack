# Git and GitHub Notes


<!-- no toc -->
## Contents
- 

## Overview
- version control system etc etc
- working tree must be clean to checkout a different branch

## Commands I don't already use or might forget
- `git status`: lists working info
- `git diff filename`: shows differences in current file (staged vs unstaged)
- `git log`: logs commits in chrono order (including the SHA of each)
- `git stash` and `git stash pop` allow you to stash changes locally, in order to have a clean working tree
    - stored in a hidden directory
- to update a previous commit, stage changes and then use `git commit --amend`
    - this replaces the entire previous commit, so a new message will be requested (unless you use the flag `--no-edit`)
- git aliases
    - `git config --global alias.br "branch"` sets `br` as the alias for `branch`

## How to Fix Oopsies
- current commit is referred to as `HEAD`
    - view it with the command `git show HEAD`
    - this command shows current changes
- to restore a file in working directory to its last commit, use the command `git checkout HEAD filename`
- to unstage a file from the current commit, use `git reset HEAD filename`
- to revert current commit to a recent one, use `git reset commitSHA` where commitSHA is the first 7 chars of a recent commit 
