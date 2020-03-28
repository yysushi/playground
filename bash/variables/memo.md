# Notes

## Environement Variables

- environmental variables can be checked with `printenv`

```shell-session
$ docker run -v $(pwd)/variables:/script --rm bash:4.4 bash /script/environment-variable.sh; echo $?
0
```

- environmental variables can be inherited to child process, but shell variables can't.

```shell-session
$ docker run -v $(pwd)/variables:/script --rm bash:4.4 bash /script/shell-variable.sh
FUGA1
FUGA2

FUGA2
```

log

```shell-session
bash-4.4# bash script/shell-variable.sh
FUGA1
FUGA2

FUGA2
bash-4.4# ls /tmp
tmp.gHpIHo
bash-4.4# cat /tmp/tmp.gHpIHo
echo $HOGE1
echo $HOGE2
```
