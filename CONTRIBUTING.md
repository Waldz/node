# Contributing guide


## Development environment

* **Step 1.** Get Golang
```bash
brew install go
brew install glide

export GOPATH=~/workspace/go
git clone git@github.com:MysteriumNetwork/node.git $GOPATH/src/github.com/mysterium/node
cd $GOPATH/src/github.com/mysterium/node
```

* **Step 2.** Compile code
```bash
glide install
go build github.com/mysterium/node
```

* **Step 3.** Prepare configuration

Enter `MYSTERIUM_API_URL` address of running [api](https://github.com/MysteriumNetwork/api) instance<br/>
Enter `NATS_SERVER_IP` address of running [communication broker](https://github.com/nats-io/gnatsd) instance

```bash
cp .env_example .env
vim .env
```

## Running

```bash
# Start communication broker
docker-compose up broker

# Start node
bin/server_build
bin/server_run

# Client connects to node
bin/client_build
bin/client_run
```

## Running client in interactive cli

```bash
# Start client with --cli
bin/client_run_cli

# Show commands
» help
[INFO] Mysterium CLI tequilapi commands:
  connect
  identities
  ├── new
  ├── list
  status
  proposals
  ip
  disconnect
  help
  quit
  stop
  unlock

# Create a customer identity
» identities new

# Unlock a customer identity
» unlock <identity>

# Show provider identities
» proposals

# Connect to a server
» connect <consumer-identity> <provider-identity>
```

## Running multiple nodes and clients at once

```bash
To run small network of nodes and clients in docker, you can use:
AGREED_TERMS_AND_CONDITIONS=<date of terms> docker-compose up
```

Be sure to replace `<date of terms>` with date of latest terms.
Date of terms will be shown when running without `AGREED_TERMS_AND_CONDITIONS` variable:

```bash
docker-compose up
```

## Dependency management

* Install project's frozen packages
```bash
glide install
glide install --force
```

* Add new package to project
```bash
glide get --quick github.com/ccding/go-stun
```

* Update package in project
```bash
vim glide.yaml
glide update
```


## Debian packaging

* **Step 1.** Get FPM tool
See http://fpm.readthedocs.io/en/latest/installing.html

```bash
brew install gnu-tar
gem install --no-ri --no-rdoc fpm
```

* **Step 2.** Get Debber tool
See https://github.com/debber/debber-v0.3

```bash
go get github.com/debber/debber-v0.3/cmd/...
```

* **Step 3.** Build .deb package
```bash
bin/server_package_debian 0.0.6 amd64
bin/client_package_debian 0.0.6 amd64
```

## Creating pull request

To contribute a code, fist you must create a pull request (PR). If your changes will be accepted
this PR will be merged into main branch.

Before creating PR be sure to: 

* **Step 1.** Make sure that no linter errors remain in **your** code

```bash
bin/lint_git
```

* **Step 2.** Ensure that all unit tests passes, no vet errors remain and code formatted.

```bash
bin/test_commit
```

After you forked a project, modified sources and run tests, you can create a pull request using this procedure:
 
 https://help.github.com/articles/creating-a-pull-request-from-a-fork/
