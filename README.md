GitDinger
=========

GitDinger monitors your local copies of git repositories and notifies you of
changes by playing a selected sound file and saying out loud who did commit what.

The ```autofetch``` configuration causes gitdinger to fetch the repositories and
thus monitor even remote branches easily. Just specify for example ```origin/master```
and set ```autofetch``` to ```true```.

## Requirements

* ```git``` installed in your $PATH
* Mac OS
	* binaries ```say``` and ```afplay```
* Linux
	* ```aplay``` (http://alsa.opensrc.org/Aplay), ```espeak``` (http://espeak.sourceforge.net/)

## Configuration

```
{
	"repos": [
		{
			"path": "/home/test/testgitrepo",
			"branch": "origin/master",
			"autofetch": true,
			"soundfile": "lm_coin.wav",
			"voice": "Zarvox"
		}
	],
	"period": 30
}
```
 
* ```repos``` - A list of all repos that should be monitored
* ```path``` - Path to the git repository
* ```branch``` - Which branch to track (```master```, ```origin/master```, ...)
* ```autofetch``` - Run ```fetch --all``` every time a check is to be performed
* ```soundfile``` - Path to a soundfile to play when changes are registered
* ```voice``` - A voice to use (get all usable by running ```say -v?``` [Only supported on OSX]
* ```period``` - How often to check (seconds)

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
