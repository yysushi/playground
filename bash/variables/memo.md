# Notes

```shell-session
$ docker run -v $(pwd)/variables:/script --rm bash:4.4 bash /script/environment-variable.sh; echo $?
0
```
