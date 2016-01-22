# Simple Revision Locking Tool [![wercker status](https://app.wercker.com/status/2205aafe44b9890ae9483913d95ed689/s "wercker status")](https://app.wercker.com/project/bykey/2205aafe44b9890ae9483913d95ed689)

Simple tool to snapshot and restore **state of all existing repositories** on the given path. It detects repositories of all famous VCS's(git/hg/bzr) and saves metadata about the repositories states to yaml the file. It will be enough to restore the state. That's it =)

### Why?
Golang has an [unusual](http://golang.org/doc/faq#get_version) approach to manage package versions. 

This is a common situation that storage and distribution of external artifacts for the project are very inconvenient. That's why there are [many](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) projects, designed to solve this problem. So, this is my solution.

**What is the difference?**

- there is no golang centric logic(`$GOPATH`), you only save and restore  state, thats all
- it deadly simple, only [three main commands](#usage)
- to survive you need only know where is your ~~towel~~ file with dependencies
- it is [binary distributed](https://github.com/olebedev/srlt/releases/), anyone can install and use it
- it is stable and [well tested](https://app.wercker.com/project/bykey/16ab277aafad959b674654a1fbc3ce9e)
- there are no agreements to follow, feel free to organize you project dependencies

### Installation

If you have Golang at your system, you may install `srlt` by `go get` tool:   
```
$ go get github.com/olebedev/srlt
```

Or you may download already compiled binary:

```
$ # osx example
$ curl -L https://github.com/olebedev/srlt/releases/download/v1.0.0/srlt-v1.0.0-64-osx.tar.gz | tar xvz
```

And install it to you `$PATH`, if you prefer.
All compiled binaries you can find [here](https://github.com/olebedev/srlt/releases/).

### Usage
It isn't neccesary to have Golang in your system and not neccesary to know, what the Golang is. This isn't a large package manager, this is just thin tool to do one simple thing.    
It have just three commands: `snapshot`, `restore` and `exec`.   
First of all, type it:

```
$ srlt
NAME:
   srlt - save and restore repositories at given path

USAGE:
   srlt [global options] command [command options] [arguments...]
   
VERSION:
   1.0.0
   
COMMANDS:
   snapshot, s	Save your current state into the file
   restore, r	Restore state from the file
   exec, e	Execute give shell programm to each dependency
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --path, -p "."		path to scan and restore, will be saved at first time
   --file, -f "srlt.yaml"	filename for read and write
   --verbose			verbose mode
   --help, -h			show help
   --version, -v		print the version
```

As you can see, there is no way to get confused.

To take a snapshot just type:

```
$ srlt snapshot
```

This commad save metadata about state of your current work directory to the `./srlt.yaml` file. It's easy to change this behavior using `path` flag:

```
$ # for current directory
$ srlt shapshot
$ # or for golang projects
$ srlt --path=`$GOPATH/src` shapshot 
```
Srlt will save base path into the file and read them next time. No need to specify path any time.

To restore type:

```
$ srlt restore
```

This will restore the state of repositories exactly as it was before. You will see operation log in stdout. If you don't have yet repositories at file system, they will be cloned as usual. It is possible to change path like in the previous example.

As additional functioanality there is `exec` command that allow us to execute _bash_ one-liner with dependency context and templating. For example, run `go install` for each repo:

```
$ srlt exec go install {{.Name}}/...
```

Or remove VCS's metadata:

```
$ srlt exec rm -rf {{.Name}}/.{{.Type}}
```

Available: `.Name` `.Type`, `.Remote`, `.Commit`.  
> Note the command will executed at base path(saved at the snapshot step).


### Welcome to contribute

Please, feel free to send pull request if you want to improve or fix some bugs. If you have some reason  to be added as collaborator, send me an [email](mailto:oolebedev@gmail.com?subject=srlt).

### TODO
Svn support.

### License
MIT
