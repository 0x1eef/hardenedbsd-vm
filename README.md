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

## Inputs

#### Options

All GitHub actions accept inputs via the "with" directive. This
action provides the following input variables. Some are
hardcoded for now but others can be customized:

* run<br>
The command to run on the hardenedBSD virtual machine. <br>
This can be any valid shell command(s).
* mem<br>
The amount of memory to allocate for the VM. <br>
This defaults to 6144MB but can be customized.
* cpu<br>
The number of CPU cores to allocate for the VM. <br>
This defaults to all available cores but can be customized.
* release<br>
The hardenedBSD release to use. <br>
This is always `16-CURRENT` for now.
* arch<br>
The CPU architecture.<br>
This is always x86_64 for now.
* filesystem<br>
The filesystem type.<br>
This is always ufs for now.

## Environment

#### VM

At the time of writing, the virtual machine has 4 vCPU cores
and 6GB of RAM by default. This can be decreased but not
increased unless GitHub increase the resources available
on the host machine. The VM image is 80GB in size and that
leaves roughly 65GB of free space for the user to use.

#### Permissions

The virtual machine is configured to run your commands as the `runner`
user although root privileges can be obtained with the
[mdo(1)](https://man.freebsd.org/cgi/man.cgi?mdo)
utility (eg `mdo -u root <command>`).

#### pkg-static

It is recommended to use pkg-static instead of pkg
for installing packages because the former is less error prone
in the virtual machine environment &ndash; where the base system
can be more recent than the package repository or vice versa and
that can sometimes cause errors related to dynamic libraries.

This is a quirk of how the virtual machine images and pkg repositories
are built manually by two different people, and I try to minimize
it happening as much as possible with a long-term solution in the
works:

```sh
mdo -u root pkg-static install -y <package>
```

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/hardenedbsd-vm)
* [git.hardenedBSD.org/@0x1eef](https://git.hardenedBSD.org/0x1eef/hardenedbsd-vm)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
