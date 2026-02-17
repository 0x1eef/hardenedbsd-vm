## About

This repository provides a GitHub action for running builds and tests
on a [hardenedBSD](https://hardenedbsd.org) virtual machine. It is
inspired by the
[vmactions](https://github.com/vmactions)
project that provides a similar service for the mainstream BSD operating
systems (FreeBSD, OpenBSD, NetBSD, etc). Their work inspired me and it was
adapted for hardenedBSD.

## Usage

#### Workflow

The following is an example GitHub workflow that uses this action to run
tests on a hardenedBSD virtual machine. It checks out the code, boots the
VM, installs the Go programming language, and then runs `make test` on the
virtual machine:

```yaml
name: My workflow
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
  workflow_dispatch:

jobs:
  test:
    name: Build
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Run test
      uses: 0x1eef/hardenedbsd-vm@v1
      with:
        release: '16-CURRENT'
        run: |
          mdo -u root pkg-static install -y go
          make test
```

#### Inputs

All GitHub actions accept inputs via the "with" directive. This
action provides a couple of input variables that can be used this
way. In the future, more variables may be supported. Certain variables,
like the CPU architecture and filesystem type are always amd64 and ufs
respectively but might be configurable in the future.

* release<br>
The hardenedBSD release to use. <br>
This can be `16-CURRENT`. Eventually, we would like to support `15-STABLE` as well.
* run<br>
The command to run on the hardenedBSD virtual machine. <br>
This can be any valid shell command(s).
* arch<br>
The CPU architecture.<br>
This is always x86_64 for now.
* filesystem<br>
The filesystem type.<br>
This is always ufs for now.
* mem<br>
The amount of memory to allocate for the VM.
This defaults to 6144MB.

#### Environment

This action is setup to boot ubuntu first, and then a hardenedBSD
virtual machine is booted from a modified virtual machine image that
is optimized for GitHub actions.

The virtual machine is configured to run your commands as the `runner`
user although root privileges can be obtained with the
[mdo(1)](https://man.freebsd.org/cgi/man.cgi?mdo)
utility (eg `mdo -u root <command>`).

The environment is configured to use pkg-static instead of pkg
for installing packages because the former is less error prone,
especially on hardenedBSD where the virtual machine could be
more recent than the package repository. For example:

```sh
mdo -u root pkg-static install -y <package>
```

Commands are written to a shell script with the name `hardenedbsd-vm.sh`,
and if a repository has a file with the same name the file will be
overwritten. Please choose a different name to avoid conflicts. Eventually
we would like to find a more robust solution.

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/hardenedbsd-vm)
* [git.hardenedBSD.org/@0x1eef](https://git.hardenedBSD.org/0x1eef/hardenedbsd-vm)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
