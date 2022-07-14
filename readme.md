# `BS64`

Encode & Decode base64

```bash
bs64 [-e (--encode) | -d (--decode)] [-f (--file) | -s (--string)] <data> [<output>]
```

## 🌿 Install Guide

- Clone repository

```bash
git clone https://github.com/aThebigbot/bs64
```

- Run install

```bash
[sudo] sh install
```
## 🐛 Example & Guide 

##### 🔒 Encode example

Encode string
```bash
bs64 --encode --string Hello
# SGVsbG8=
```


Encode file (Hello.txt, `Hello`)

```bash
bs64 --encode --file Hello.txt
# SGVsbG8=
```

Set output file (The output file may or may not exist) 

```bash
bs64 --encode [--string | --file] [string | file] [output file]
```
