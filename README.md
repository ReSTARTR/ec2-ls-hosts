ec2-ls-hosts: an alternative tool for list ec2 instances
====

`ls-hosts` is a simple cli-tool for describing ec2 instances.
This tool will simplify the describing instances.
You can integrate this tool with unix tools (eg: awk, ssh, peco, and so on.)

dependencies
----

- [aws-sdk-go](https://github.com/aws/aws-sdk-go)
- [go-ini](https://github.com/go-ini/ini)
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter)

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

Integration with zsh and peco
----

- With this integration, you can ssh login with interactive host selector

Dependencies

- zsh
- [peco](https://github.com/peco/peco)

```~/.zshrc
function peco-ec2-ls-hosts () {
  BUFFER=$(
    /path/to/ls-hosts -fields tag:Name,private-ip | \
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
