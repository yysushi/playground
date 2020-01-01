# Notes for tox

- test log

```shell-session
koketani:tox (master *%>)$ tox
GLOB sdist-make: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/tox/setup.py
py27 inst-nodeps: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/tox/.tox/.tmp/package/1/hello-0.0.1.zip
py27 installed: DEPRECATION: Python 2.7 will reach the end of its life on January 1st, 2020. Please upgrade your Python as Python 2.7 won't be maintained after that date. A future version of pip will drop support for Python 2.7. More details about Python 2 support in pip, can be found at https://pip.pypa.io/en/latest/development/release-process/#python-2-support,discover==0.4.0,hello==0.0.1
py27 run-test-pre: PYTHONHASHSEED='1937168484'
py27 run-test: commands[0] | discover
.
----------------------------------------------------------------------
Ran 1 test in 0.000s

OK
py37 inst-nodeps: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/tox/.tox/.tmp/package/1/hello-0.0.1.zip
py37 installed: attrs==19.3.0,bcrypt==3.1.7,cached-property==1.5.1,certifi==2019.11.28,cffi==1.13.2,chardet==3.0.4,Click==7.0,cryptography==2.8,docker==4.1.0,docker-compose==1.25.0,dockerpty==0.4.1,docopt==0.6.2,entrypoints==0.3,filelock==3.0.12,flake8==3.7.9,future==0.18.2,greenlet==0.4.15,hello==0.0.1,idna==2.8,importlib-metadata==1.3.0,jedi==0.15.2,jsonschema==3.2.0,mccabe==0.6.1,more-itertools==8.0.2,msgpack==0.6.2,neovim==0.3.1,packaging==19.2,paramiko==2.7.1,parso==0.5.2,pathspec==0.6.0,pbr==5.4.4,pipenv==2018.11.26,pluggy==0.13.1,proselint==0.10.2,py==1.8.1,pycodestyle==2.5.0,pycparser==2.19,pyflakes==2.1.1,PyNaCl==1.3.0,pynvim==0.4.0,pyparsing==2.4.6,pyrsistent==0.15.6,python-jsonrpc-server==0.3.2,python-language-server==0.31.4,PyYAML==3.13,requests==2.22.0,six==1.13.0,stevedore==1.31.0,texttable==1.6.2,toml==0.10.0,tox==3.14.3,ujson==1.35,urllib3==1.25.7,virtualenv==16.7.9,virtualenv-clone==0.5.3,virtualenvwrapper==4.8.4,websocket-client==0.57.0,yamllint==1.20.0,zipp==0.6.0
py37 run-test-pre: PYTHONHASHSEED='1937168484'
py37 run-test: commands[0] | discover
.
----------------------------------------------------------------------
Ran 1 test in 0.000s

OK
______________________________________________ summary ______________________________________________
  py27: commands succeeded
    py37: commands succeeded
      congratulations :)
```
