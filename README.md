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

### Using your LIFX Access Token in Filament
Once you have an access token, you use Filament like this:
```
import(
    'github.com/panicpanicpanic/filament'
    'github.com/panicpanicpanic/filament/lifx'
)

    client := lifx.Client{
            AccessToken: "someRandomToken"
        }
    
    _, err := filament.SomeMethod(&client)
    if err != nil {
        fmt.Println("handle your error")
    }
```

# Available Methods
Below are the current methods available in Filament
### GetLights
> Gets lights belonging to the authenticated account.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/lights/:selector`

`GetLights` accepts a `lifx.Client` (which should contain your access token), and a `selector` string (the `uid` of the device you want to retrieve). If a `selector` is not provided, `filament` will return all of the devices your access token grants access to.

#### Example:
```
import(
    'github.com/panicpanicpanic/filament'
    'github.com/panicpanicpanic/filament/lifx'
)

var err error
var selector string

client := lifx.Client{
    AccessToken:"someRandomToken"
}

devices, err := filament.GetLights(&client, selector)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### GetScenes
### ValidateColor
### SetState
### SetStates
### ActivateScene
### Cycle
### PulseEffect
### BreatheEffect
### TogglePower
### StateDelta

## Contributing
Want to contribute? Just fork the repo, make a feature/bug PR and send it over! Just make sure your tests are passing :)

License
----

MIT


**I'm not affiliated with LIFX, but they are fucking awesome, so you should buy their 💡!**
