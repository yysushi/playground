# Note for pipenv

## alternatives 

- `pipenv graph`: pip list

```shell-session
koketani:pipenv (master %=)$ pipenv graph
Courtesy Notice: Pipenv found itself running within a virtual environment, so it will automatically use that environment, instead of creating its own for any project. You can set PIPENV_IGNORE_VIRTUALENVS=1 to force pipenv to ignore that environment and create its own instead. You can set PIPENV_VERBOSITY=-1 to suppress this warning.
Warning: No virtualenv has been created for this project yet! Consider running `pipenv install` first to automatically generate one for you or see`pipenv install --help` for further instructions.

koketani:pipenv (master %=)$ source venv/bin/activate

koketani:pipenv (master %=)$ pipenv graph
Warning: No virtualenv has been created for this project yet! Consider running `pipenv install` first to automatically generate one for you or see`pipenv install --help` for further instructions.

koketani:pipenv (master %=)$ ls
memo.md venv
```

- `pipenv install`: mkvirtualenv hoge

```shell-session
koketani:pipenv (master %=)$ pipenv install
Creating a virtualenv for this project…
Pipfile: /Users/koketani/Developments/go/src/github.com/koketani/playground/python/pipenv/Pipfile
Using /Users/koketani/Developments/go/src/github.com/koketani/playground/python/pipenv/venv/bin/python (3.7.5) to create virtualenv…
⠦ Creating virtual environment...Already using interpreter /Users/koketani/Developments/go/src/github.com/koketani/playground/python/pipenv/venv/bin/python
Using base prefix '/usr/local/bin/../Cellar/python/3.7.5/bin/../Frameworks/Python.framework/Versions/3.7'
New python executable in /Users/koketani/.local/share/virtualenvs/pipenv-NockF3QV/bin/python
Installing setuptools, pip, wheel...
done.

✔ Successfully created virtual environment!
Virtualenv location: /Users/koketani/.local/share/virtualenvs/pipenv-NockF3QV
Creating a Pipfile for this project…
Pipfile.lock not found, creating…
Locking [dev-packages] dependencies…
Locking [packages] dependencies…
Updated Pipfile.lock (a65489)!
Installing dependencies from Pipfile.lock (a65489)…
  🐍   ▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉ 0/0 — 00:00:00
To activate this project's virtualenv, run pipenv shell.
Alternatively, run a command inside the virtualenv with pipenv run.

koketani:pipenv (master %=)$ cat Pipfile
[[source]]
name = "pypi"
url = "https://pypi.org/simple"
verify_ssl = true

[dev-packages]

[packages]

[requires]
python_version = "3.7"
koketani:pipenv (master %=)$ cat Pipfile.lock
{
    "_meta": {
        "hash": {
            "sha256": "7e7ef69da7248742e869378f8421880cf8f0017f96d94d086813baa518a65489"
        },
        "pipfile-spec": 6,
        "requires": {
            "python_version": "3.7"
        },
        "sources": [
            {
                "name": "pypi",
                "url": "https://pypi.org/simple",
                "verify_ssl": true
            }
        ]
    },
    "default": {},
    "develop": {}
}

koketani:pipenv (master %=)$ pipenv graph
koketani:pipenv (master %=)$
```

- `pipenv shell`: workon hoge

```shell-session
koketani:pipenv (master %=)$ pipenv shell
Launching subshell in virtual environment…
 . /Users/koketani/.local/share/virtualenvs/pipenv-NockF3QV/bin/activate
 koketani:pipenv (master %=)$  . /Users/koketani/.local/share/virtualenvs/pipenv-NockF3QV/bin/activate
 koketani:pipenv (master %=)$ pip list
 Package    Version
 ---------- -------
 pip        19.3.1
 setuptools 42.0.2
 wheel      0.33.6

koketani:pipenv (master %=)$ which python
/Users/koketani/.local/share/virtualenvs/pipenv-NockF3QV/bin/python

koketani:pipenv (master %=)$ deactivate
```

- `pipenv install requests`: pip install requests

```shell-session
koketani:pipenv (master %=)$ pipenv install requests
Installing requests…
Adding requests to Pipfile's [packages]…
✔ Installation Succeeded
Pipfile.lock (444a6d) out of date, updating to (a65489)…
Locking [dev-packages] dependencies…
Locking [packages] dependencies…
✔ Success!
Updated Pipfile.lock (444a6d)!
Installing dependencies from Pipfile.lock (444a6d)…
  🐍   ▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉ 5/5 — 00:00:00

koketani:pipenv (master %=)$ pipenv graph
requests==2.22.0
  - certifi [required: >=2017.4.17, installed: 2019.11.28]
  - chardet [required: >=3.0.2,<3.1.0, installed: 3.0.4]
  - idna [required: >=2.5,<2.9, installed: 2.8]
  - urllib3 [required: >=1.21.1,<1.26,!=1.25.1,!=1.25.0, installed: 1.25.7]

koketani:pipenv (master %=)$ cat Pipfile
[[source]]
name = "pypi"
url = "https://pypi.org/simple"
verify_ssl = true

[dev-packages]

[packages]
requests = "*"

[requires]
python_version = "3.7"
koketani:pipenv (master %=)$ cat Pipfile.lock
{
    "_meta": {
        "hash": {
            "sha256": "bb57e0d7853b45999e47c163c46b95bc2fde31c527d8d7b5b5539dc979444a6d"
        },
        "pipfile-spec": 6,
        "requires": {
            "python_version": "3.7"
        },
        "sources": [
            {
                "name": "pypi",
                "url": "https://pypi.org/simple",
                "verify_ssl": true
            }
        ]
    },
    "default": {
        "certifi": {
            "hashes": [
                "sha256:017c25db2a153ce562900032d5bc68e9f191e44e9a0f762f373977de9df1fbb3",
                "sha256:25b64c7da4cd7479594d035c08c2d809eb4aab3a26e5a990ea98cc450c320f1f"
            ],
            "version": "==2019.11.28"
        },
        "chardet": {
            "hashes": [
                "sha256:84ab92ed1c4d4f16916e05906b6b75a6c0fb5db821cc65e70cbd64a3e2a5eaae",
                "sha256:fc323ffcaeaed0e0a02bf4d117757b98aed530d9ed4531e3e15460124c106691"
            ],
            "version": "==3.0.4"
        },
        "idna": {
            "hashes": [
                "sha256:c357b3f628cf53ae2c4c05627ecc484553142ca23264e593d327bcde5e9c3407",
                "sha256:ea8b7f6188e6fa117537c3df7da9fc686d485087abf6ac197f9c46432f7e4a3c"
            ],
            "version": "==2.8"
        },
        "requests": {
            "hashes": [
                "sha256:11e007a8a2aa0323f5a921e9e6a2d7e4e67d9877e85773fba9ba6419025cbeb4",
                "sha256:9cf5292fcd0f598c671cfc1e0d7d1a7f13bb8085e9a590f48c010551dc6c4b31"
            ],
            "index": "pypi",
            "version": "==2.22.0"
        },
        "urllib3": {
            "hashes": [
                "sha256:a8a318824cc77d1fd4b2bec2ded92646630d7fe8619497b142c84a9e6f5a7293",
                "sha256:f3c5fd51747d450d4dcf6f923c81f78f811aab8205fda64b0aba34a4e48b0745"
            ],
            "version": "==1.25.7"
        }
    },
    "develop": {}
}
```

- dev
```
koketani:pipenv (master +%=)$ pipenv install --dev
Installing dependencies from Pipfile.lock (444a6d)…
  🐍   ▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉ 5/5 — 00:00:00

koketani:pipenv (master *+%=)$ pipenv install --dev -r ~/.dotfiles/global-requirements.txt
Requirements file provided! Importing into Pipfile…
Pipfile.lock (7bb5f3) out of date, updating to (444a6d)…
Locking [dev-packages] dependencies…
✔ Success!
Locking [packages] dependencies…
✔ Success!
Updated Pipfile.lock (7bb5f3)!
Installing dependencies from Pipfile.lock (7bb5f3)…
  🐍   ▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉ 42/42 — 00:00:04

koketani:pipenv (master *%>)$ pip list
Package          Version
---------------- ----------
certifi          2019.11.28
pip              19.2.3
pipenv           2018.11.26
setuptools       41.2.0
virtualenv       16.7.9
virtualenv-clone 0.5.3
koketani:pipenv (master *%>)$ pipenv run pip list
Package                Version
---------------------- ----------
attrs                  19.3.0
bcrypt                 3.1.7
cached-property        1.5.1
certifi                2019.11.28
cffi                   1.13.2
chardet                3.0.4
Click                  7.0
cryptography           2.8
docker                 4.1.0
docker-compose         1.25.0
dockerpty              0.4.1
docopt                 0.6.2
entrypoints            0.3
flake8                 3.7.9
future                 0.18.2
idna                   2.8
importlib-metadata     1.3.0
jedi                   0.15.2
jsonschema             3.2.0
mccabe                 0.6.1
more-itertools         8.0.2
paramiko               2.7.1
parso                  0.5.2
pathspec               0.7.0
pip                    19.3.1
pipenv                 2018.11.26
pluggy                 0.13.1
proselint              0.10.2
pycodestyle            2.5.0
pycparser              2.19
pyflakes               2.1.1
PyNaCl                 1.3.0
pyrsistent             0.15.6
python-jsonrpc-server  0.3.2
python-language-server 0.31.4
PyYAML                 3.13
requests               2.22.0
setuptools             42.0.2
six                    1.13.0
texttable              1.6.2
ujson                  1.35
urllib3                1.25.7
virtualenv             16.7.9
virtualenv-clone       0.5.3
websocket-client       0.57.0
wheel                  0.33.6
yamllint               1.20.0
zipp                   0.6.0
```

## Command notes 

```shell-session
koketani:pipenv (master %=)$ pip help

Usage:
  pip <command> [options]

Commands:
  install                     Install packages.
  download                    Download packages.
  uninstall                   Uninstall packages.
  freeze                      Output installed packages in requirements format.
  list                        List installed packages.
  show                        Show information about installed packages.
  check                       Verify installed packages have compatible dependencies.
  config                      Manage local and global configuration.
  search                      Search PyPI for packages.
  wheel                       Build wheels from your requirements.
  hash                        Compute hashes of package archives.
  completion                  A helper command used for command completion.
  debug                       Show information useful for debugging.
  help                        Show help for commands.

General Options:
  -h, --help                  Show help.
  --isolated                  Run pip in an isolated mode, ignoring environment variables and user configuration.
  -v, --verbose               Give more output. Option is additive, and can be used up to 3 times.
  -V, --version               Show version and exit.
  -q, --quiet                 Give less output. Option is additive, and can be used up to 3 times (corresponding to WARNING, ERROR, and CRITICAL logging levels).
  --log <path>                Path to a verbose appending log.
  --proxy <proxy>             Specify a proxy in the form [user:passwd@]proxy.server:port.
  --retries <retries>         Maximum number of retries each connection should attempt (default 5 times).
  --timeout <sec>             Set the socket timeout (default 15 seconds).
  --exists-action <action>    Default action when a path already exists: (s)witch, (i)gnore, (w)ipe, (b)ackup, (a)bort.
  --trusted-host <hostname>   Mark this host as trusted, even though it does not have valid or any HTTPS.
  --cert <path>               Path to alternate CA bundle.
  --client-cert <path>        Path to SSL client certificate, a single file containing the private key and the certificate in PEM format.
  --cache-dir <dir>           Store the cache data in <dir>.
  --no-cache-dir              Disable the cache.
  --disable-pip-version-check
                              Don't periodically check PyPI to determine whether a new version of pip is available for download. Implied with --no-index.
  --no-color                  Suppress colored output
```

```shell-session
koketani:pipenv (master %=)$ pipenv --help
Usage: pipenv [OPTIONS] COMMAND [ARGS]...

Options:
  --where             Output project home information.
  --venv              Output virtualenv information.
  --py                Output Python interpreter information.
  --envs              Output Environment Variable options.
  --rm                Remove the virtualenv.
  --bare              Minimal output.
  --completion        Output completion (to be eval'd).
  --man               Display manpage.
  --support           Output diagnostic information for use in GitHub issues.
  --site-packages     Enable site-packages for the virtualenv.  [env var:
                      PIPENV_SITE_PACKAGES]
  --python TEXT       Specify which version of Python virtualenv should use.
  --three / --two     Use Python 3/2 when creating virtualenv.
  --clear             Clears caches (pipenv, pip, and pip-tools).  [env var:
                      PIPENV_CLEAR]
  -v, --verbose       Verbose mode.
  --pypi-mirror TEXT  Specify a PyPI mirror.
  --version           Show the version and exit.
  -h, --help          Show this message and exit.


Usage Examples:
   Create a new project using Python 3.7, specifically:
   $ pipenv --python 3.7

   Remove project virtualenv (inferred from current directory):
   $ pipenv --rm

   Install all dependencies for a project (including dev):
   $ pipenv install --dev

   Create a lockfile containing pre-releases:
   $ pipenv lock --pre

   Show a graph of your installed dependencies:
   $ pipenv graph

   Check your installed dependencies for security vulnerabilities:
   $ pipenv check

   Install a local setup.py into your virtual environment/Pipfile:
   $ pipenv install -e .

   Use a lower-level pip command:
   $ pipenv run pip freeze

Commands:
  check      Checks for security vulnerabilities and against PEP 508 markers
             provided in Pipfile.
  clean      Uninstalls all packages not specified in Pipfile.lock.
  graph      Displays currently-installed dependency graph information.
  install    Installs provided packages and adds them to Pipfile, or (if no
             packages are given), installs all packages from Pipfile.
  lock       Generates Pipfile.lock.
  open       View a given module in your editor.
  run        Spawns a command installed into the virtualenv.
  shell      Spawns a shell within the virtualenv.
  sync       Installs all packages specified in Pipfile.lock.
  uninstall  Un-installs a provided package and removes it from Pipfile.
  update     Runs lock, then sync.


koketani:pipenv (master +%>)$ pipenv run vim main.py
```
