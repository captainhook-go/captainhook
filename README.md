# Warning #
Nothing here works, this is still in very early development stages and not really working so far.
This README is kind of a declaration of intent and shows what captainhook could or should be one day.

# CaptainHook

<img src="https://captainhookphp.github.io/captainhook/gfx/ch.png" alt="CaptainHook logo" align="right" width="200"/>

*CaptainHook* is an easy to use and very flexible git hook library for developers.
It enables you to configure your git hook actions in a simple json file and share them within your team.

You can use *CaptainHook* to validate or prepare your commit messages, ensure code quality
or run unit tests before you commit or push changes to git. You can automatically clear
local caches or install the correct dependencies after pulling the latest changes.

You can run any cli command, use loads of built in functionality.
For more information have a look at the [documentation](https://captainhook-go.github.io/captainhook/ "CaptainHook Documentation").

## Installation

Use apt to install the captain
```bash
apt-get install captainhook
```

Or use *Brew* to install *CaptainHook*.
```bash
brew install captainhook
```

## Setup
After installing CaptainHook you can use the *captainhook* executable to create a configuration.
```bash
captainhook configure
```
Now there should be a *captainhook.json* configuration file.

Now you have to activate the hooks by installing them to
your local .git repository. To do so just run the following *captainhook* command.
```bash
captainhook install
```

## Configuration

Here's an example *captainhook.json* configuration file.
```json
{
  "hooks": {
    "commit-msg": {
      "actions": [
        {
          "action": "CaptainHook::Message.Beams"
        }
      ]
    },
    "pre-commit": {
      "actions": [
        {
          "action": "go test"
        }
       ]
    },
    "pre-push": {
      "actions": []
    }
  }
}
```

## Contributing

So you'd like to contribute to the `CaptainHook` library? Excellent! Thank you very much.
I can absolutely use your help.

Have a look at the [contribution guidelines](CONTRIBUTING.md).