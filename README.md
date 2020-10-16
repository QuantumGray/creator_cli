# creator

### toolkit to make create apps from templates & create templates from apps

![](https://img.shields.io/github/stars/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/forks/QuantumGray/creator_cli.svg) ![](https://img.shields.io/github/tag/QuantumGray/fluttercreator_cli.svg) ![](https://img.shields.io/github/release/QuantumGray/creator_cli.svg) ![](https://img.shields.io/github/issues/QuantumGray/creator_cli.svg)

##### features :

- create apps from templates that are published on [CREATOR REGISTRY](http://google.com "creator_registry")
- cache templates locally
- create templates from existing apps (soon)
- automatically set up firebase for your project (soon)

##### installation :
- download creator via [this link](http://google.com "this link") or clone the repository

```bash
git clone https://github.com/QuantumGray/creator_cli
```


- set PATH variable for to the SDK folder - [how to set a path variable for your specific machine](http://google.com "how to set a path variable for your specific machine")


##### usage :

how to use creator

###### create apps
command **create** plus an **8**-digit hex identifier from the template registry downloads the app and creates an app in the current directory - it prompts you to enter a name for your app

```bash
creator create 1T3F5S7E
```

initiates the Creator app with firebase preconfigured

- `-f BUNDLE_ID`

stores the template locally
- `-c`

###### create app templates :
command **createtemplate** plus a name for your template creates a template from your Creator app in current directory

```bash
creator createtemplate myTemplate
```

optionaly declare a path to the app

- `-p "PATH/TO/YOUR/APP"`

load up the template to [CREATOR REGISTRY](http://google.com "CREATOR REGISTRY") with your API KEY that you can find in the account panel

- `-up API_KEY`
