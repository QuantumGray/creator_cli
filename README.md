# fluttercreator

### toolkit to make create apps from templates & create templates from apps

![](https://img.shields.io/github/stars/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/forks/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/tag/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/release/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/issues/QuantumGray/fluttercreator_cli.svg)

##### features :

- create Flutter apps from templates that are published on [CREATOR REGISTRY](http://google.com "fluttercreator_registry")
- cache templates locally
- create templates from existing flutter apps (soon)
- automatically set up firebase for your project (soon)

##### installation :
- download fluttercreator via [this link](http://google.com "this link") or clone the repository

```bash
git clone https://github.com/QuantumGray/fluttercreator_cli
```


- set PATH variable for to the SDK folder - [how to set a path variable for your specific machine](http://google.com "how to set a path variable for your specific machine")


##### usage :

how to use fluttercreator

###### create Flutter apps
command **create** plus an **8**-digit hex identifier from the template registry downloads the app and creates a flutter app in the current directory - it prompts you to enter a name for your app

```bash
fluttercreator create 1T3F5S7E
```

initiates the Flutter app with firebase preconfigured

- `-f BUNDLE_ID`

stores the template locally
- `-c`

###### create app templates :
command **createtemplate** plus a name for your template creates a template from your Flutter app in current directory

```bash
fluttercreator createtemplate myTemplate
```

optionaly declare a path to the Flutter app

- `-p "PATH/TO/FLUTTER/APP"`

load up the template to [CREATOR REGISTRY](http://google.com "CREATOR REGISTRY") with your API KEY that you can find in the account panel

- `-up API_KEY`
