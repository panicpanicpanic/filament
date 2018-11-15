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

```sh
$ go get github.com/panicpanicpanic/filament
```

## How To Use

### Authentication
In order to access your LIFX devices, you must first have a valid access token. You can visit the [LIFX Cloud page](https://cloud.lifx.com/settings) to set one up:
![](https://files.readme.io/Uw1PRNPoQ7nFUYDuT6oA_GetLIFXToken.gif)

### Using your LIFX Access Token in Filament
Once you have an access token, you use Filament like this:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
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
### GetLights
> Gets lights belonging to the authenticated account.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/lights/:selector`

`GetLights` accepts a `lifx.Client` (which should contain your access token), and a `selector` string (the `uid` of the device you want to retrieve). If a `selector` is not provided, `filament` will return all of the devices your access token grants access to.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
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
> Lists all the scenes available in the users account. 

LIFX API Endpoint Reference: `https://api.lifx.com/v1/scenes`

`GetScenes` accepts a `lifx.Client` (which should contain your access token). Scenes listed here can be activated with the `ActivateScene` method.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

scenes, err := filament.GetScenes(&client)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### ValidateColor
> Lets you validate a user's color string and return the hue, saturation, brightness and kelvin values that the API will interpret as.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/color`

`ValidateColor` accepts a `lifx.Client` (which should contain your access token), and a `color` string. 

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

color := "red"

validateColor, err := filament.ValidateColor(&client, color)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### SetState
> Sets the state of the lights within the selector.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/lights/:selector/state`

`SetState` accepts a `lifx.Client` (which should contain your access token), a `selector` string and a `payload` inteface. All parameters (except for the selector) are optional. If you don't supply a parameter, `filament` will leave that value untouched.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "deviceUID"
payload := []byte(`
  {
    "power": "on"
  }
`)

results, err := filament.SetState(&client, selector, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### SetStates
> Sets different states on multiple selectors in a single request.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/lights/states`

`SetStates` accepts a `lifx.Client` (which should contain your access token) and a `payload` inteface. All parameters (except for the selector) are optional. If you don't supply a parameter, `filament` will leave that value untouched.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

payload := []byte(`
  {
    "states": [
      {
        "selector": "[selector 1]",
        "power": "on"
      },
      {
        "selector": "[selector N]",
        "brightness": 0.5
      }
    ],
    "defaults": {
      "duration": 5.0
    }
  }
`)

results, err := filament.SetStates(&client, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### ActivateScene
> Activates a scene from your LIFX account.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/scenes/scene_id::scene_uuid/activate`

`ActivateScene` accepts a `lifx.Client` (which should contain your access token), a `sceneUUID` string, and a `payload` inteface. All parameters (except for the selector) are optional. If you don't supply a parameter, `filament` will leave that value untouched. 

Please note some details about using the `fast` property when changing state: https://api.developer.lifx.com/docs/activate-scene#fast-mode

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

sceneUUID := "actualSceneUUID"

payload := []byte(`
  {
    "fast": true
  }
`)

results, err := filament.ActivateScene(&client, sceneUUID, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### Cycle
> Make the light(s) cycle to the next or previous state in a list of states.

LIFX API Endpoint Reference: `https://api.lifx.com/v1/lights/:selector/cycle`

`Cycle` accepts a `lifx.Client` (which should contain your access token), a `selector` string, and a `payload` inteface.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "lightUID"

payload := []byte(`
  {
    "states": [
      {
        "brightness": 1.0
      },
      {
        "brightness": 0.5
      },
      {
        "brightness": 0.1
      },
      {
        "power": "off"
      }
    ],
    "defaults": {
      "power": "on",
      "saturation": 0,
      "duration": 2.0
    }
  }
`)

results, err := filament.Cycle(&client, selector, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### PulseEffect
> Performs a pulse effect by quickly flashing between the given colors.

LIFX API Endpoint Reference:
`https://api.lifx.com/v1/lights/:selector/effects/pulse`

`PulseEffect` accepts a `lifx.Client` (which should contain your access token), a `selector` string, and a `payload` inteface.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "lightUID"

payload := []byte(`
  {
    "period": "2",
    "cycles": "5",
    "color": "green"
  }
`)

results, err := filament.PulseEffect(&client, selector, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### BreatheEffect
> Performs a breathe effect by slowly fading between the given colors.

LIFX API Endpoint Reference:
`https://api.lifx.com/v1/lights/:selector/effects/breathe`

`BreatheEffect` accepts a `lifx.Client` (which should contain your access token), a `selector` string, and a `payload` inteface.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "lightUID"

payload := []byte(`
  {
    "period": "2",
    "cycles": "5",
    "color": "green"
  }
`)

results, err := filament.BreatheEffect(&client, selector, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### TogglePower
> Turns off lights if any of them are on, or turns them on if they are all off.

LIFX API Endpoint Reference:
`https://api.lifx.com/v1/lights/:selector/toggle`

`TogglePower` accepts a `lifx.Client` (which should contain your access token), and a `selector` string.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "lightUID"

results, err := filament.TogglePower(&client, selector)
if err!= nil {
    fmt.Println("handle your error here")
}
```

### StateDelta
> Changes the state of the lights by the amount specified.

LIFX API Endpoint Reference:
`https://api.lifx.com/v1/lights/:selector/state/delta`

`StateDelta` accepts a `lifx.Client` (which should contain your access token), a `selector` string, and a `payload` inteface. Parameters other than `power` and `duration` change the state of the lights by the amount specified. `power` and `duration` act like they do for the `SetState` method.

#### Example:
```
import(
    "github.com/panicpanicpanic/filament"
    "github.com/panicpanicpanic/filament/lifx"
)

var err error

client := lifx.Client{
    AccessToken:"someRandomToken"
}

selector := "lightUID"

payload := []byte(`
  {
    "brightness": "0.1"
  }
`)

results, err := filament.StateDelta(&client, selector, payload)
if err!= nil {
    fmt.Println("handle your error here")
}
```

## Contributing
Want to contribute? Just fork the repo, make a feature/bug PR and send it over! Just make sure your tests are passing :)

License
----

MIT


**I'm not affiliated with LIFX, but they are fucking awesome, so you should buy their 💡!**
