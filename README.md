# Filament [![Go Report Card](https://goreportcard.com/badge/github.com/panicpanicpanic/filament)](https://goreportcard.com/report/github.com/panicpanicpanic/filament)

<p align="center">
  <img src="https://media.giphy.com/media/3o6Ztm25ikO467NGOk/giphy.gif">
</p>

Filament is an unofficial Go wrapper for [LIFX's HTTP API](https://api.developer.lifx.com/docs).

## Build Status
| branch | status |
| --- | --- |
| `develop` | [![CircleCI](https://circleci.com/gh/panicpanicpanic/filament/tree/develop.svg?style=svg)](https://circleci.com/gh/panicpanicpanic/filament) |
| `master` | [![CircleCI](https://circleci.com/gh/panicpanicpanic/filament/tree/master.svg?style=svg)](https://circleci.com/gh/panicpanicpanic/filament) |


## Installation

You can install Filament with either ```go get```
```sh
$ go get github.com/panicpanicpanic/filament
```

 or ```dep```

```sh
$ dep ensure -add github.com/panicpanicpanic/filament
```

## How To Use

### Authentication
In order to access your LIFX devices, you must first have a valid access token. You can visit the [LIFX Cloud page](https://cloud.lifx.com/settings) to set one up:
![](https://files.readme.io/Uw1PRNPoQ7nFUYDuT6oA_GetLIFXToken.gif)

## Contributing
Want to contribute? Just fork the repo, make a feature/bug PR and send it over! Just be sure you write some tests, too!

License
----

MIT


**I'm not affiliated with LIFX, but they are fucking awesome, so you should buy their 💡!**
