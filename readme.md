# Simple Revision Locking Tool [![wercker status](https://app.wercker.com/status/16ab277aafad959b674654a1fbc3ce9e/s/ "wercker status")](https://app.wercker.com/project/bykey/16ab277aafad959b674654a1fbc3ce9e)

Simple tool to snapshot and restore **state of all existing repositories** on the given path. It detects repositories of all famous VCS's(git/hg/bzr) and saves to json file metadata about the repositories states. So, it will be enough to restore the state. That's it =)

### Why?
Golang has an [unusual](http://golang.org/doc/faq#get_version) approach to manage package versions. 

This is a common situation that storage and distribution of external artifacts for the project are very inconvenient. That's why there are [many](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) projects, designed to solve this problem. So, this is my solution.

**What is the difference?**

- there is no golang centric logic(`$GOPATH`, `$GOROOT`, etc.), you only save and restore  state, thats all
- it deadly simple, only [two main commands](#usage)
- to survive you need only know where is your ~~towel~~ file with dependencies
- it is [binary distributed](https://github.com/olebedev/srlt/releases/), anyone can install and use it
- it is stable and [well tested](https://app.wercker.com/project/bykey/16ab277aafad959b674654a1fbc3ce9e)
- there are no agreements to follow, feel free to organize you project dependencies

### Installation

If you have Golang at your system, you may install `srlt` by `go get` tool:   
```bash
$ go get github.com/olebedev/srlt
```

Or you may download already compiled binary file:

```bash
$ # osx example
$ curl -L https://github.com/olebedev/srlt/releases/download/v0.2.0/srlt-v0.2.0-64-osx.tar.gz | tar xvz
```

And install it to you `$PATH`, if you prefer.
All compiled binaries you can find [here](https://github.com/olebedev/srlt/releases/).

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
  --basepath="pwd"         : path to scan and restore 
  --file="srlt.json"       : filename for read and write 
```

As you can see, there is no way to get confused.

To take a snapshot just type it:

```bash
$ srlt shapshot
```

This commad save metadata about state of your current work directory to the `./srlt.json` file. It's easy to change this behavior using `basepath` flag:

```bash
$ # for current directory
$ srlt shapshot
$ # or for golang projects
$ srlt --basepath=`$GOPATH/src` shapshot 
```

To restore type it:

```bash
$ srlt restore
```

This will restore the state of repositories exactly as it was before. You will see operation log in stdout.  
If you don't have yet repositories at file system, they will be cloned as usual. It is possible to change path like in the previous example.

Enjoy!

### Welcome to contribute

Please, feel free to send pull request if you want to improve or fix some bugs. If you have some reason  to be added as collaborator, send me an [email](mailto:oolebedev@gmail.com?subject=srlt).

### TODO
Svn support.
