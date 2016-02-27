aws-ls-hosts - list aws instances
====

dependencies
----

- [aws-sdk-go](https://github.com/aws/aws-sdk-go)
- [go-ini](https://github.com/go-ini/ini)

usage
----

```bash
$ ls-hosts
i-00000001 10.0.0.1 app01 running
i-00000002 10.0.0.2 app02 running
i-00000003 10.0.0.3 app03 running
```

options
----

### with options

`-filters`

tag filter

```
-filters=key1:value1,key2:value2...
```


`-columns`

support columns

- instance-id
- private-ip
- public-ip
- tag:*

```
-columns=c1,c2,c3,...
```

### with config file

/etc/ls-hosts.conf or ~/.ls-hosts

```
[options]
creds   = shared
region  = ap-northeast-1
filters = instance-state-name:running
tags    = Role:app,Env:production
fields = instance-id,tag:Name,public-ip,private-ip
```

build
----

```bash
$ make build (-B)
```

with zsh and peco
----

Dependencies

- zsh
- [peco](https://github.com/peco/peco)

```~/.zshrc
function peco-ec2-ls-hosts () {
  BUFFER=$(
    /path/to/ls-hosts -columns=tag:Role,public-ip | \
    peco --prompt "EC2 >" --query "$LBUFFER" | \
    awk '{printf ssh %s\n", $2}'
  )
  CURSOR=$#BUFFER
  zle accept-line
  zle clear-screen
}
zle -N peco-ec2-ls-hosts
bindkey '^oo' peco-ec2-ls-hosts
```
