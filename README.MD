# Wizard home
Useful tool that provides an easy use, easy install method to manage yours config files offering an easy way to setup an zero state OS

## Install

```bash
git clone https://github.com/PullolilEzequiel/wizard-home && cd wizard-home && go install ./bin/wizard_home.go && cd .. && rm -r wizard-home
```
---
To use `go install` to install the project you may need to have the golang environment configured, checkout this link [setting up go env](https://gist.github.com/fakihariefnoto/5b5fa888cecf5cd7910d3e19a9febbde)

You can also do the **"normal"** way and build the project and move it to the /usr/bin folder.

1. **Clone the repo**
``` 
git clone https://github.com/PullolilEzequiel/wizard-home && cd wizard-home
``` 
2. **Build and move to bin apps**
```
go build && sudo mv /usr/bin/wizard_home
```

## Configure
wizard home_works in a particular way, that is, to configure the project you have to create another configuration file :)

But the power of managing all your configuration files by simply creating one configuration file is worth it, right?

So lets start, first init the config 
```bash
wizard_home init
```
Once the configuration has been initialized, we are going to take a short tour of the fields that exist


```bash
nvim $HOME/.config/wizard_home/config.json
```

The initial configuration mainly has 2 fields.
   - repository_url: the url of your remote repository where you will save your config
   - configs_to_persist: The configuration files you are looking to persist

```json 
{
    "repository_url": <set the remote url of your github/gitlab repo>,
    "configs_to_persist" : [
        "example/path/*",
        "otherExample/path/**",
        "exampleFile/.zshrc"
    ]
}
```



## Use
With `wizard_home` you can: save the  config, reverse all the config or install the config of other users using their repository

### save configs
Every time you want save the changed configs you have to use wizard_home save.
```bash
wizard_home save
```

### reverse
If you want to reverse your old configs use the following command

```bash
wizard_home reverse
```
> [!WARNING]
> This command uses the configurations that are already in your remote repository

### install config
You can install all the wizard home config of others users parsing the git url, this command needs the config.json file for works.

> [!CAUTION]
> This command not have any integriti check.

```bash
wizard_home setup <remote url>
```

### add file
The command wizard_home add <file or directory path> allows you to add a file or directory to your setup without modifying the configuration file (config.json) manually.
```bash
wizard_home add <file or directory path>
```



## Upcoming features

1. Add 'date' argument to reverse command
   * The ability to revert settings to a previous temporary state
   * Add a README.md to the remote repository of the config wizard to be able to view possible reverse dates
