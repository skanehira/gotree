# gotree
This is tree command written in go

# Install
```
go get -u github.com/skanehira/gotree
```

# Usage
```
// default display current dir
$ gotree

// specified dir
$ gotree ./dir

// limit depth
// default depth is 99
$ gotree -L 2 .

// print ansi color mode
$ gotree -C

// exclude file or dir
$ gotree -EX node_modules
```
