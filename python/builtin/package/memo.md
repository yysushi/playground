# python package

## all commands

```bash
python setup.py --help-commands
...
  build             build everything needed to install
  build_py          "build" pure Python modules (copy to build directory)
  build_ext         build C/C++ extensions (compile/link to build directory)
  build_clib        build C/C++ libraries used by Python extensions
  build_scripts     "build" scripts (copy and fixup #! line)
  clean             clean up temporary files from 'build' command
  install           install everything from build directory
  install_lib       install all Python modules (extensions and pure Python)
  install_headers   install C/C++ header files
  install_scripts   install scripts (Python or otherwise)
  install_data      install data files
  sdist             create a source distribution (tarball, zip file, etc.)
  register          register the distribution with the Python package index
  bdist             create a built (binary) distribution
  bdist_dumb        create a "dumb" built distribution
  bdist_rpm         create an RPM distribution
  bdist_wininst     create an executable installer for MS Windows
  check             perform some checks on the package
  upload            upload binary package to PyPI
```

## source distribution

```bash
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ python setup.py sdist -v
running sdist
running egg_info
creating mypackage.egg-info
writing requirements to mypackage.egg-info/requires.txt
writing mypackage.egg-info/PKG-INFO
writing top-level names to mypackage.egg-info/top_level.txt
writing dependency_links to mypackage.egg-info/dependency_links.txt
writing manifest file 'mypackage.egg-info/SOURCES.txt'
reading manifest file 'mypackage.egg-info/SOURCES.txt'
writing manifest file 'mypackage.egg-info/SOURCES.txt'
running check
creating mypackage-1.0.2
creating mypackage-1.0.2/mypackage.egg-info
copying files to mypackage-1.0.2...
copying README.md -> mypackage-1.0.2
copying setup.py -> mypackage-1.0.2
copying mypackage.egg-info/PKG-INFO -> mypackage-1.0.2/mypackage.egg-info
copying mypackage.egg-info/SOURCES.txt -> mypackage-1.0.2/mypackage.egg-info
copying mypackage.egg-info/dependency_links.txt -> mypackage-1.0.2/mypackage.egg-info
copying mypackage.egg-info/requires.txt -> mypackage-1.0.2/mypackage.egg-info
copying mypackage.egg-info/top_level.txt -> mypackage-1.0.2/mypackage.egg-info
Writing mypackage-1.0.2/setup.cfg
creating dist
Creating tar archive
removing 'mypackage-1.0.2' (and everything under it)
```

- diffs after the command

```bash
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ ls -alh mypackage.egg-info/
total 28K
drwxr-xr-x 2 koketani koketani 4.0K May  7 18:03 .
drwxr-xr-x 5 koketani koketani 4.0K May  7 18:04 ..
-rw-r--r-- 1 koketani koketani  312 May  7 18:03 PKG-INFO
-rw-r--r-- 1 koketani koketani  182 May  7 18:03 SOURCES.txt
-rw-r--r-- 1 koketani koketani    1 May  7 18:03 dependency_links.txt
-rw-r--r-- 1 koketani koketani   24 May  7 18:03 requires.txt
-rw-r--r-- 1 koketani koketani    1 May  7 18:03 top_level.txt
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ tar -ztvf dist/mypackage-1.0.2.tar.gz
drwxr-xr-x koketani/koketani 0 2019-05-07 18:03 mypackage-1.0.2/
-rw-r--r-- koketani/koketani 7 2019-05-07 17:50 mypackage-1.0.2/README.md
drwxr-xr-x koketani/koketani 0 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/
-rw-r--r-- koketani/koketani 1 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/top_level.txt
-rw-r--r-- koketani/koketani 182 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/SOURCES.txt
-rw-r--r-- koketani/koketani   1 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/dependency_links.txt
-rw-r--r-- koketani/koketani 312 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/PKG-INFO
-rw-r--r-- koketani/koketani  24 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/requires.txt
-rw-r--r-- koketani/koketani 312 2019-05-07 18:03 mypackage-1.0.2/PKG-INFO
-rw-r--r-- koketani/koketani  38 2019-05-07 18:03 mypackage-1.0.2/setup.cfg
-rw-r--r-- koketani/koketani 475 2019-05-07 16:27 mypackage-1.0.2/setup.py
```

## binary distribution

- candidates
  - python setup.py bdist
  - python setup.py bdist_egg
  - python setup.py bdist_wheel

```bash
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ date; python setup.py bdist -v
Tue May  7 19:07:43 JST 2019
running bdist
running bdist_dumb
running build
installing to build/bdist.linux-x86_64/dumb
running install
running install_egg_info
running egg_info
writing requirements to mypackage.egg-info/requires.txt
writing mypackage.egg-info/PKG-INFO
writing top-level names to mypackage.egg-info/top_level.txt
writing dependency_links to mypackage.egg-info/dependency_links.txt
reading manifest file 'mypackage.egg-info/SOURCES.txt'
writing manifest file 'mypackage.egg-info/SOURCES.txt'
Copying mypackage.egg-info to build/bdist.linux-x86_64/dumb/usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info
running install_scripts
Creating tar archive
removing 'build/bdist.linux-x86_64/dumb' (and everything under it)
```

- diffs after the command

```bash
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ tar -ztvf dist/mypackage-1.0.2.tar.gz
drwxr-xr-x koketani/koketani 0 2019-05-07 18:03 mypackage-1.0.2/
-rw-r--r-- koketani/koketani 7 2019-05-07 17:50 mypackage-1.0.2/README.md
drwxr-xr-x koketani/koketani 0 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/
-rw-r--r-- koketani/koketani 1 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/top_level.txt
-rw-r--r-- koketani/koketani 182 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/SOURCES.txt
-rw-r--r-- koketani/koketani   1 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/dependency_links.txt
-rw-r--r-- koketani/koketani 312 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/PKG-INFO
-rw-r--r-- koketani/koketani  24 2019-05-07 18:03 mypackage-1.0.2/mypackage.egg-info/requires.txt
-rw-r--r-- koketani/koketani 312 2019-05-07 18:03 mypackage-1.0.2/PKG-INFO
-rw-r--r-- koketani/koketani  38 2019-05-07 18:03 mypackage-1.0.2/setup.cfg
-rw-r--r-- koketani/koketani 475 2019-05-07 16:27 mypackage-1.0.2/setup.py
koketani:~/Developments/git/github.com/koketani/playground/python/builtin/package (master *%>)$ tar -ztvf dist/mypackage-1.0.2.linux-x86_64.tar.gz
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/local/
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/local/lib/
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/local/lib/python2.7/
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/
drwxr-xr-x koketani/koketani 0 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/
-rw-r--r-- koketani/koketani 1 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/top_level.txt
-rw-r--r-- koketani/koketani 182 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/SOURCES.txt
-rw-r--r-- koketani/koketani   1 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/dependency_links.txt
-rw-r--r-- koketani/koketani 312 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/PKG-INFO
-rw-r--r-- koketani/koketani  24 2019-05-07 19:07 ./usr/local/lib/python2.7/dist-packages/mypackage-1.0.2.egg-info/requires.txt
```

### wheel

```
```

### egg
