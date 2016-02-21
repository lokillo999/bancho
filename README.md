# Bancho [![Build Status](https://travis-ci.org/bnch/bancho.svg?branch=master)](https://travis-ci.org/bnch/bancho)

This repository is a clone of the bancho server. Bancho is the software that
powers [osu!](https://osu.ppy.sh). Currently players of osu! are getting a bit
sick of the whole ecosystem, so there are many attempts to move out of it, and
that means that plenty of private servers are popping up, and being developed.

This server is written with the help of:

1. Coffee
2. justm3/HoLLy-HaCKeR's [custom-bancho](https://github.com/HoLLy-HaCKeR/custom-bancho).
3. Last but not least, the private server I and a friend of mine had been
   writing since about August 2015. The server is currently closed source due to
   some privacy issues and bad code in general. That friend of mine in about
   January 2016 started developing the version 1.5 of the said private server,
   which included a bancho server. It's even worse in code than justm3's
   custom-bancho, so don't expect to understand much if you haven't dug deep
   into the structure of the osu! packets. Anyway, if you want to take it for
   reference for making your bancho server, [there you go](http://hastebin.com/opadinohej.php).

## Setting up

We're just gonna assume you have [Go](https://golang.org) installed on your
computer, and have your goenv up and running. To make sure of that, run `go env`
in your terminal/command line and check that it's good. We're also assuming that
you're developing on a UNIX environment, thus the command will be in shell.
Despite that, commands should be fairly easy to port over to Windows.

We're also gonna assume you have a webserver up and running on your computer,
such as nginx or apache. Development without those will come in the near future.

And, as a last thing, we're gonna assume you have a mysql server set up,
listening to port 3306 on localhost, and having the user root with no password.
This will change in the very near future.

First step: clone the repo.

```sh
cd $GOPATH
# If you haven't forked
mkdir -p src/github.com/bnch/bancho
cd src/github.com/bnch/bancho
git clone git@github.com:bnch/bancho.git .
# If you have forked
mkdir -p src/github.com/GITHUBUSERNAME/bancho
cd src/github.com/GITHUBUSERNAME/bancho
git clone git@github.com:GITHUBUSERNAME/bancho.git .
```

Second step: build executable.

```sh
go build
```

Third step: set up a reverse proxy on either nginx or apache. Just google
"nginx reverse proxy" and you should easily get a sample configuration. It should
proxy requests to c1.ppy.sh, c.ppy.sh, osu.ppy.sh and a.ppy.sh to 127.0.0.1:3000.

Fourth step: change your hosts file to work on bancho. (note that on windows
you'll need to write a line for each host, as you can't point multiple hostnames
to an IP with a single line on windows' hosts file. blame microsoft.)

```
127.0.0.1 osu.ppy.sh a.ppy.sh c.ppy.sh c1.ppy.sh
```

Fifth step: create your user!

```
./bancho mkuser 'username' 'password'
```

Sixth step: start server

```
./bancho
```

Final step: open up osu! (should be fallback unless you have managed to set up
HTTPS), and if it works, high five yourself.
