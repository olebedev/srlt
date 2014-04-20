# Simple Revision Locking Tool [![wercker status](https://app.wercker.com/status/16ab277aafad959b674654a1fbc3ce9e/s/ "wercker status")](https://app.wercker.com/project/bykey/16ab277aafad959b674654a1fbc3ce9e)

Simple tool to snapshot and restore **state of all exists repositories** at the given path. It detect repos of all famous VCS's(git/hg/bzr) and save to json file metadata about the repos states. So, it will be enough to restore the state. That's it =)

### What for?
Golang has an unusual approach to managing dependencies in the project. approach is to have at all depending on the local machine, in the state in which they are installed. And in other words - do not have such a general approach. 

But often there is a situation that store and distribute all external artifacts for the project is very inconvenient. Therefore, there is [a lot of](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) projects, designed to solve this problem. But, as often happens, some people have their own vision implement any utility. So, this is my own.

### Installation

If you have Golang at your system, you may install `srlt` by `go` tool:   
```bash
$ go get github.com/olebedev/srlt
```

Or you may download already compiled binary file:

```bash
$ curl -L https://github.com/olebedev/srlt/releases/download/0.1.0/srlt-0.1.0-64-osx > srlt
```

And install it to you `$PATH`, if you prefer.
All compiled binary you can find [here](https://github.com/olebedev/srlt/releases/).

### Usage
It isn't neccesary to have Golang in your system and not neccesary to know, what is Golang. This isn't a large package manager, this is just thin tool to do one simple thing.    
It have two main commands: `snapshot` and `restore`.   
First of all, type it:

```bash
$ srlt

Usage of srlt:
  srlt [options] snapshot  : save your current state
  srlt [options] restore   : restore state from file

Options:
  --basepath="$GOPATH/src" : path to scan and restore it
  --file="srlt.json"       : filename for read and write it
```

As you can see, there is no way to get confused.

To take a snapshot just type it:

```bash
$ srlt shapshot
```

This commad save metadata about state of your `$GOPATH/src` to the `./srlt.json` file. `srlt` was originally done as a dependencies manager for Golang. Therefore `basepath` option has defaults value - `$GOPATH/src`. It's easy to change this behavior using `basepath` flag:

```bash
$ srlt --basepath=`pwd` shapshot
```
In the next version I plan to cut this default behavior, and do it for the current working directory.

To restore type it:

```bash
$ srlt restore
```

You will see operation log in stdout.  
Enjoy!

### Welcome to contribute

Please, feel free to send pull request if you want to improve or fix some bugs. If you have some reason  to be added as collaborator, send me an [email](mailto:oolebedev@gmail.com?subject=srlt).














