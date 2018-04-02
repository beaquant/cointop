# cointop

> Coin tracking for hackers

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/cointop/master/LICENSE.md) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/cointop?)](https://goreportcard.com/report/github.com/miguelmota/cointop) [![GoDoc](https://godoc.org/github.com/miguelmota/cointop?status.svg)](https://godoc.org/github.com/miguelmota/cointop)

<img src="./assets/screenshot-001.gif" width="880" />

[`cointop`](https://github.com/miguelmota/cointop) is a fast command-line interface application for viewing cryptocurrency stats and information in your terminal. The interface is inspired by [`htop`](https://en.wikipedia.org/wiki/Htop).

## Install

Make sure to have [go](https://golang.org/) (1.9+) installed, then do:

```bash
go get -u github.com/miguelmota/cointop
```

<!--
#### Alternatively (without go)

```
sudo curl -s "https://raw.githubusercontent.com/miguelmota/cointop/master/install.sh?$(date +%s)" | bash
```
-->

## Usage

```bash
$ cointop
```

### Table commands

List of shortcuts:

|Key|Action|
|----|------|
|`<up>`|navigate up|
|`<down>`|navigate down|
|`<right>`|next page|
|`<left>`|previous page|
|`<ctrl-u>`|page up|
|`<ctrl-d>`|page down|
|`<ctrl-n>`|alias to next page|
|`<ctrl-p>`|alias to previous page|
|`<enter>`|visit highlighted coin on CoinMarketCap|
|`<space>`|alias to `<enter>`
|`j`|alias to `<down>`|
|`k`|alias to `<up>`|
|`r`|sort by *[r]ank*|
|`n`|sort by *[n]ame*|
|`s`|sort by *[s]ymbol*|
|`p`|sort by *[p]rice*|
|`m`|sort by *[m]arket cap*|
|`v`|sort by *24 hour [v]olume*|
|`1`|sort by *[1] hour change*|
|`2`|sort by *[2]4 hour change*|
|`7`|sort by *[7] day change*|
|`t`|sort by *[t]otal supply*|
|`a`|sort by *[a]vailable supply*|
|`l`|sort by *[l]ast updated*|
|`q`|[q]uit|
|`<esc>`|alias to quit|
|`<ctrl-c>`|alias to quit|

<!--
|`h`|toggle [h]elp|
|`?`|alias to help|
-->

## FAQ

- Q: Where is the data from?

  - A: The data is from [Coin Market Cap](https://coinmarketcap.com/).

- Q: What coins does this support?

  - A: This supports any coin listed on [Coin Market Cap](https://coinmarketcap.com/).

- Q: How often is the data polled?

  - A: Data gets polled once every minute by default.

- Q: I installed cointop without errors but the command is not found.

  - A: Make sure your `GOPATH` and `PATH` is set correctly.
    ```bash
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```

- Q: What is the size of the binary?

  - A: The executable is only ~1.9MB in size.

## Authors

- [Miguel Mota](https://github.com/miguelmota)

## License

Released under the MIT license.
