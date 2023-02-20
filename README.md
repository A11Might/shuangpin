<h1 align="center">Welcome to shuangpin 👋</h1>
<p>
  <a href="https://github.com/A11Might/shuangpin" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="https://github.com/A11Might/shuangpin/blob/master/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/huqihh" target="_blank">
    <img alt="Twitter: huqihh" src="https://img.shields.io/twitter/follow/huqihh.svg?style=social" />
  </a>
</p>

> Practice shuangpin in your terminal.

### ✨ Demo

![typer](assets/shuangpin@0.05.gif?raw=true)

## Install

you can [install Go](https://golang.org/dl/) and build from source (requires Go 1.16+):

```sh
go install github.com/A11Might/shuangpin@latest
```

## Usage

By default, the natural code scheme is used, and pinyin and key prompts are displayed for random Chinese character practice.

```sh
shuangpin
```

It is recommended for beginners to choose the shuangpin scheme they want to learn, and then use the `all sequence` mode to practice.

```
shuangpin -s zrm -m sequence
```

To get a full overview of all available options, run `shuangpin --help`.

```sh
NAME:
   shuangpin - Practice shuangpin in your terminal

USAGE:
   shuangpin [global options] command [command options] [arguments...]

COMMANDS:
   support, s  View the supported shuangpin schemes and practice mode
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --scheme value, -s value  choose shuangpin scheme (default: "zrm")
   --mode value, -m value    choose practice mode (default: "random")
   --pinyin, -p              disable pinyin prompt (default: false)
   --keyboard, -k            disable key prompt (default: false)
   --help, -h                show help
```

## Author

👤 **Kohath Hu**

* Website: https://a11might.github.io/
* Twitter: [@huqihh](https://twitter.com/huqihh)
* Github: [@A11Might](https://github.com/A11Might)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/A11Might/shuangpin/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2023 [Kohath Hu](https://github.com/A11Might).<br />
This project is [MIT](https://github.com/A11Might/shuangpin/blob/master/LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_