<!-- omit in toc -->
# Contributing

Thank you so much for your interest in contributing!. All types of contributions are encouraged and valued. Below are some guidelines for how to contribute.

## Getting Started

Contributions are made to this project via issues and Pull Requests (PRs), below is some general guidelines

- Submit all changes directly to the `main` branch, there's no seperate `develop` or `release` branch
- Please open an issue before embarking on any work 
- For any questions feel free to contact me (see [how to contact](#how-to-contact))
- Changes might occur to this document, check for changes regularly


## Project Setup

This guide will help you get started to contribute to pz-battlesnake, it's recommended that you read the project setup without doing any of the commands the steps tells you to do, as this will help you understand how to setup your project locally. 

Once you've read the project setup, you can setting up your project by following and doing what the steps tell you to do. 

<!-- omit in toc -->
### Prerequisites

Before starting to contribute, you need to have the following prerequisites:
- Go 1.18 or higher
- Python 3.10 or higher
- Make
- [poetry](https://python-poetry.org/docs/) - used to manage dependencies

<!-- omit in toc -->
### Step 1: Clone the repository & get latest code

[Fork](http://help.github.com/fork-a-repo/) the repo, clone your fork and setup remotes. if you are not sure if you should use HTTPS or SSH, pick HTTPS or setup your account to use SSH if you haven't already (see [guide](https://docs.github.com/en/authentication/connecting-to-github-with-ssh))
> **Remember to replace \<your-username> with your username **

<details>
  <summary>HTTPS</summary>  
  
  ```bash
    # Clone your fork of the repo into the current directory
    git clone https://github.com/<your-username>/pz-battlesnake
    # Navigate to the newly cloned directory
    cd pz-battlesnake
    # Assign the original repo to a remote called "upstream"
    git remote add upstream https://github.com/DaBultz/pz-battlesnake
  ```
</details>

<details>
  <summary>SSH</summary>
  
  ```bash
    # Clone your fork of the repo into the current directory
    git clone git@github.com:<your-username>/pz-battlesnake.git
    # Navigate to the newly cloned directory
    cd pz-battlesnake
    # Assign the original repo to a remote called "upstream"
    git remote add upstream git@github.com:DaBultz/pz-battlesnake.git
  ```
</details>


if you have previously cloned the repo, pull the latest changes from the upstream repo

```bash
git checkout main
git pull upstream main
```

<!-- omit in toc -->
### Step 2: Install dependencies

Start my making an virtualenv using poetry by executing the following command:
```
poetry shell
```

Once you have created your virtualenv, you can install dependencies using the following:
```
poetry install
```

All dependencies are installed into the virtualenv

<!-- omit in toc -->
### Step 3: Profit

The project should now be installed locally and be ready to use.

## How to build docs 

The builds support a live reloading of the docs, run the following command:
```
make serve-docs
```

## How to contact

if there's any questions, you can always get in contact with me through these channels:
- Send me a DM on Discord (username: `Bultz`)
  - This will be the fastest way to contact me, I will respond to you as soon as I can
  - This option is possible if you are on the Battlesnake discord server, you will be able to easily find me on there.
- Open an issue on GitHub
