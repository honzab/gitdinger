GitDinger
=========

GitDinger monitors your local copies of git repositories and notifies you of
changes by playing a selected sound file an saying out loud who did commit what.

## Requirements

* ```git``` installed in your $PATH
* Mac OS, specifically preinstalled binaries ```say``` and ```afplay```


## Installation / Running

* Clone this repository
* Edit the sample configuration file
* Run ```go run gitdinger.go```
* Profit!

Optionally you can specify a different config file, by usigng the ```-config``` flag.
Use ```-h``` to get a usage description.

## Sample output

```
2014/05/27 19:21:28 Setting up listener for repo '/home/test/testgitrepo':master to check every 1 seconds.
2014/05/27 19:21:29 Checking testgitrepo:master
2014/05/27 19:21:29  Setting initial state on testgitrepo to 18d9b82
2014/05/27 19:21:30 Checking testgitrepo:master
2014/05/27 19:21:31 Checking testgitrepo:master
2014/05/27 19:21:32 Checking testgitrepo:master
2014/05/27 19:21:32  Dinging 3 times.
2014/05/27 19:21:36  Saying: "Jan Brucek's commit Commit 26524"
2014/05/27 19:21:40  Saying: "Jan Brucek's commit Commit 21404"
2014/05/27 19:21:45  Saying: "Jan Brucek's commit Commit 28014"
2014/05/27 19:21:45 Checking testgitrepo:master
2014/05/27 19:21:45 Checking testgitrepo:master
2014/05/27 19:21:46 Checking testgitrepo:master
2014/05/27 19:21:47 Checking testgitrepo:master
```
