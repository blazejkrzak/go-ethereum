# Contributing

Thank you for considering to help out with the source code! We welcome 
contributions from anyone on the internet, and are grateful for even the 
smallest of fixes!

If you'd like to contribute to our fork please contact us at ed@silesiacoin.org
In further process we will be having open repo, so anyone could join.

## Coding guidelines
Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go 
[formatting](https://golang.org/doc/effective_go.html#formatting) guidelines 
(i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go 
[commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `release/1.9` branch.
 * Commit messages should be in format: `Feature: ${issue_number} ${issue_description}`

## Synchronizing with go-ethereum official
Master branch should be updated if in need with ethereum-go implementation. 
We work on release branch and if release is completed we can merge commits that are needed to `devel` branch.
Devel branch should be cross-compatible with master of main go-ethereum implementation

## Can I have feature X

Before you submit a feature request, please check and make sure that it isn't 
possible through some other means. The JavaScript-enabled console is a powerful 
feature in the right hands. Please check our 
[Wiki page](https://github.com/ethereum/go-ethereum/wiki) for more info
and help.

## Configuration, dependencies, and tests

Please see the [Developers' Guide](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies
and testing procedures.
