<!-- omit in toc -->
# Contributing

<!-- omit in toc -->
## Content <a name="toc"></a>

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Project Setup](#project-setup)
- [How to contribute documentation](#how-to-contribute-documentation)
- [How to contribute code](#how-to-contribute-code)
- [How to contact](#how-to-contact)
- [Style Guide](#style-guide)
  - [Git](#git)
    - [Commit Messages](#commit-messages)
    - [Pull Requests](#pull-requests)
    - [Branch Naming](#branch-naming)

## Introduction

Thank you so much for your interest in contributing!. All types of contributions are encouraged and valued. See the [table of contents](#toc) for different ways to help and details about how this project handles them!📝


## Getting Started

Contributions are made to this project via issues and Pull Requests (PRs), below is some general guidelines

- This repository follows the covential commits standard, (read the spec [here](https://www.conventionalcommits.org/en/v1.0.0/#specification) if you are unfamiliar), this applies to the following:
  - Commits, see [commit messages](#commit-messages) under style guide
  - Pull Request (Only the title), see [pull requests](#pull-requests) under style guide
  - Branch names, see [branch names](#branch-naming) under style guide
- Submit all changes directly to the `main` branch, there's no seperate `develop` or `release` branch
  - We will do our best to keep the `main` in a good and stable state
- All changes are welcome, please consult one of the following for what you wish to contribute:
  - [How to contribute documentation](#how-to-contribute-documentation)
  - [How to contribute code](#how-to-contribute-code)
  - Anything else, please contact me (see [how to contact](#how-to-contact))
- Changes might occur to this document, check for changes regularly

## Project Setup

## How to contribute documentation

Documentation is a key part of this project, contribution of any size is welcome and encouraged, below is a list of some examples of what we would like to see:
- Tutorials (e.g. how to use the environment, how to implement DQN/A2C/PPO, NNUE and other algorithms)
- Battlesnakes made using this environment (they can also be submitted to [awesome-battlesnake](https://github.com/xtagon/awesome-battlesnake) if they comply with the contribution guidelines of awesome-battlesnake)
- Chores (Fixing grammar, Improving Wording etc.)

The list is not exhaustive, but it provides some ideas of what we would like to see.

To contribute documentation:

- [Set up the project](#project-setup).
- Edit or add any relevant documentation.
- Make sure your changes are formatted correctly and consistently with the rest of the documentation.
- Re-read what you wrote, and run a spellchecker on it to make sure you didn't miss anything.
- Write clear, concise commit message(s) (see [commit messages](#commit-messages) style found under style guide).
- Go to https://github.com/DaBultz/pz-battlesnake/pulls and open a new pull request with your changes.
  - Make sure your pull request follows the [pull requests](#pull-requests) style (found under style guide).

Once you've submitted the PR:

- The PR will be reviewed by the maintainers,
- They may request changes, if that's the case apply the changes and ask for another review (don't open a new PR)
- If your PR gets declined, That's fine, we still appreciate your contribution and thank you for your contribution.
- If your PR gets approved, the PR will be merged into the `main` branch and go live shortly


## How to contribute code

Code is a key part of this project, contribution of any size is welcome and encouraged

To contribute code:

* [Set up the project](#project-setup).
* Make any necessary changes to the source code.
* Include any [additional documentation](#contribute-documentation) the changes might need.
- Write clear, concise commit message(s) (see [commit messages](#commit-messages) style found under style guide).
- Dependency updates, additions, or removals must be in individual commits, and most follow this format: `<prefix>(deps): PKG@VERSION`, where `<prefix>` is a one of the available prefixes (see [commit messages](#commit-messages) style found under style guide).
- Go to https://github.com/DaBultz/pz-battlesnake/pulls and open a new pull request with your changes.
- Make sure your pull request follows the [pull requests](#pull-requests) style (found under style guide).

Once you've submitted the PR:

- The PR will be reviewed by the maintainers,
- They may request changes, if that's the case apply the changes and ask for another review (don't open a new PR)
- If your PR gets declined, That's fine, we still appreciate your contribution and thank you for your contribution.
- If your PR gets approved, the PR will be merged into the `main` branch and go be published shortly (only for `bug` and `feat`) otherwise it will be hold off until the next release.


## How to contact

if there's any questions, you can always get in contact with me through these channels:
- Send me a DM on Discord (username: `Bultz`)
  - This will be the fastest way to contact me, I will respond to you as soon as I can
  - This option is possible if you are on the Battlesnake discord server, you will be able to easily find me on there, as i am active
- Open an issue on GitHub

## Style Guide

### Conventional Commit

Conventional Commits set the following structure for the commit messages:
```
<type>(scope): <subject>

<body>

<footer(s)>
```

**Type: (required)**

Only the following types are to be used:
- `build` - Changes to the build tools or dependencies.
- `chore` - maintaining the project (no change of functionality)
- `ci` - Changes to the CI system
- `docs` - Changes to the documentation
- `fix` - Bug fixes (not build/ci related)
- `perf`- Changes which improve performance
- `refactor` - A change that neither fixes a bug or adds a feature
- `revert` - Revert a previous commit
- `style` - Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `test` - Changes to the tests


**Scope: (required):**

Only the following scopes are to be used:
- `deps` - Changes to dependencies
- `env` - Changes to the environment (all files in the `pz_battlestake/env` directory)
- `lib` - Changes to the Go library (all .go files)
- `misc` - Changes that do not fit into any of the other categories

**Subject: (required):**

- English
- Imperative, present tense (i.e. »change« instead of »changed« or »changes«)
- Short description (~50 characters is the soft limit)
- first letter lowercase
- Skip the full stop (.)

**Body (optional):**

- English
- Imperative, present tense (i.e. »change« instead of »changed« or »changes«)
- Breaking change must be written in the body, using this format:
  - `BREAKING CHANGE: <description>`, where `<description>` is a description of the breaking change

**Footer: (optional)**

- Reference to the issue or PR that this commit is associated with.

For more information please check out:
- [Semantic branch names and commit messages in Git](https://dev-tips.com/git/semantic-branch-names-and-commit-messages-in-git)
- [Conventional Commits](https://conventionalcommits.org/en/v1.0.0/#specification)


### Commit Messages

When doing any commit, please make sure they follow the [Conventional Commit](#conventional-commit) style style found under style guide..

### Pull Requests

Our pull requests follow the following structure:
- Title must follow this pattern: `<type>(scope): <subject>`, where `<type>` is one of the types above, `<scope>` is one of the scopes above and `<subject>` is the subject of the pull request.
- The description, does not enforce any style. it's up to you what's relevant.

### Branch Naming

We do not enforce any branch naming conventions on forks, but we do recommend the following structure to make:
- `<type>/<subject>`, where type is one of the types found under [commit messages](#commit-messages)

