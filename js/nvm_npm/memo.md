# about nvm and npm

## nvm
manage version of nodejs

- commands

```
nvm install --lts
nvm use --lts

node -v
```

- difference b/w node and nodejs
<https://stackoverflow.com/questions/20057790/what-are-the-differences-between-node-js-and-node?answertab=votes#tab-top>
<https://stackoverflow.com/questions/48513484/what-is-the-difference-between-node-and-nodejs?answertab=votes#tab-top>

## npm

### first things to know

- package: a file or directory described by a package.json; group of one or more modules
  - to publish npm registry, package.json is must to be included
  - name and version are included
- module: any file or directory in the `node_modules` directory that can be located by the Node.js `require()` function
- scope: package private/public type. unscoped packages are always public. some scoped packages are public, but, others are private. by defualt, scoped packages are private (= need pass command line to make them public).
  - scoped package: @user-name/package-name
    - it is sometimes public
  - unscoped package: package-name

### install/download

- download/install packages locally

  ```bash
  \$ npm install package-name
  \$ # or
  \$ npm install @scope/package-name
  \$ # check
  \$ ls node_modules
  ```

- installed package version depends on `package.json`
- install with dist-tags

  ```bash
  \$ npm install example-package@beta
  ```

- download/install packages locally

  ```bash
  \$ npm instal -g package-name
  ```

- troubleshooting in download: <https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally>

- how to package project
  - initialize and describe

    ```bash
    \$ npm init
    ```

  - install library

    ```bash
    \$ # install and note in dependencies in package.json
    \$ npm instal package-name --save-prod
    \$ # install and note in devDependencies in package.json
    \$ npm instal package-name --save-dev
    \$ # install without note in devDependencies in package.json
    \$ npm instal package-name --no-save
    ```

## playground

- npm

```bash
\$ docker build -t hoge .
```

## note for myself

- what's entrypoint
  - <https://stackoverflow.com/questions/32800066/what-is-entry-point-in-npm-init?answertab=votes#tab-top>
- what's test command
