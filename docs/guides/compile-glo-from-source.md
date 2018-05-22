# Compiling `Glo` From Source

So you want to compile `Glo` on your own instead of using prebuilt? Maybe you are developing against it, or maybe you just don't trust the prebuilts :eyes:. This guide will walk you through setting that up.

## Install Go

First thing you will need is to install `Go` for your operating system.

Official Go docs are here: https://golang.org/doc/install

**Via Snap on Linux**

Snap Website: https://snapcraft.io/

```bash
sudo snap install --classic go
```

**Via Homebrew on MacOS**

Homebrew Website: https://brew.sh/

```bash
brew install go
```
**Via Chocolatey on Windows**

Chocolatey Website: https://chocolatey.org/

```bash
choco install golang
```

## Clone `Glo` Github Repo

```bash
git clone https://github.com/haloplatform/glo.git

cd glo
```

## Compile It

**Linux/MacOS**

```bash
make all
```

**Windows**

_coming soon_:tm: (Windows is very complicated to compile atm, detailed docs are coming for it.)

## Use It

Refer to the [guides section](../README.md#guides-100) of the docs to learn how to use `Glo`
