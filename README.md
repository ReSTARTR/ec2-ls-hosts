aws-ls-hosts - list aws instances
====

usage
----

```bash
$ ls-hosts -filters=Env:production,Role:bar -columns=instance-id,private-ip,tag:Role,tag:aws:autoscaling:groupName
```

options
----

### `-filters`

tag filter

```
-filters=key1:value1,key2:value2...
```


### `-columns`

support columns

- instance-id
- private-ip
- public-ip
- tag:*

```
-columns=c1,c2,c3,...
```

build
----

```bash
$ go get github.com/aws/aws-sdk-go/aws
$ go build -o ls-hosts ./main.go
```

with zsh and peco
----

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
