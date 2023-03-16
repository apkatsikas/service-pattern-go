service-pattern-go
-------
## mocking
* mockery v 2.22.1 https://github.com/vektra/mockery/releases/download/v2.22.1/mockery_2.22.1_Linux_arm64.tar.gz
* make exexcutable
* cd interfaces
* /usr/local/bin/mockery or mockery --name=IPlayerService
* you need to do this after updating services!

## go tooling
* i dont have GOROOT set
* i deleted /usr/bin/go - go 1.15
* to trace stuff between double quotes "PS4='+$BASH_SOURCE> ' BASH_XTRACEFD=7 bash -xl 7>&2"
* check  /etc/bash.bashrc
* made a symlink like "ln -s /root/Go/bin/go1.20.1 /usr/bin/go"
* be careful of /root/go vs /root/Go/

## TODOS
* update all deps once repo is ready
* start with app serving CSV and /random endpoint
* cleanup README
* through composition, entity should have create/update dates
* start new repo with this as boilerplate
* make my model!
* log rotation
* figure out staticcheck or equivalent
* negative input tests
* 5 minute timer after insert - sync to GCS bucket
* move from google drive to GCS bucket
* backup logic - every night prob better than after every insert wait 5 mins...
* suppress sqlite warnings
* * https://www.sqlite.org/c3ref/backup_finish.html#sqlite3backupinit - online backup API
* get this working in digital ocean
* try 3 month kubernetes from google
* create e2e tests
* kubernetes/docker
* GCS
* GCS auth
* VACUUM main into '/mnt/c/Users/apkat/Desktop/vacuum.sqlite';
* ATTACH "/mnt/c/Users/apkat/Desktop/vacuum.sqlite" as main2;

## RUnning
* docker build -t test .
* docker run -p 8080:8080 test
