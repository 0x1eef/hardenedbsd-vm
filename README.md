## About

This repository provides a GitHub action for running builds and tests
on a [HardenedBSD](https://hardenedbsd.org) virtual machine. It is
inspired by the
[vmactions](https://github.com/vmactions)
project that provides a similar service for the mainstream BSD operating
systems (FreeBSD, OpenBSD, NetBSD, etc). Their work inspired me and it was
adapted for HardenedBSD.

## Example

```yaml
name: Example
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
  workflow_dispatch:

jobs:
  vm-test:
    name: Build
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Run test
      uses: hardenedbsd/vm@v1
      with:
        release: '15-STABLE'
        run: |
          uname -a
```

## Sources

* [GitHub](https://github.com/0x1eef/hardenedbsd-vm)
* [git.HardenedBSD.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/hardenedbsd-vm)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
