# VC services - Build

## Prerequisites
- Go 1.18

## Prerequisites for Tests 
- Make
- Docker
- Docker-compose
- Configure Docker to use GitHub Packages: [Authenticate](https://help.github.com/en/packages/using-github-packages-with-your-projects-ecosystem/configuring-docker-for-use-with-github-packages#authenticating-to-github-packages) 
  using a [GitHub token](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line#creating-a-token) 

## Targets

```
# run checks and unit tests
make all

# run linter checks
make checks

# run unit tests
make unit-test

# run checks and unit tests
make all
    
# run BDD tests
make bdd-test
```

## BDD Test Prerequisites

To run BDD tests (`make bdd-test`) you need to modify your hosts file (`/etc/hosts` on \*NIX) to add the following lines, to allow a few of the bdd test containers to be connected to externally. 

    127.0.0.1 testnet.orb.local
    127.0.0.1 vault.kms.example.com
