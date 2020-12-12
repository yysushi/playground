# Tutorial PHP

<https://www.php.net/manual/ja/tutorial.php>

## Findings

- [php tag](https://www.php.net/manual/ja/language.basic-syntax.phptags.php)
  - php tag in html is `<?php ... ?>`.
  - php echo tag in html is `<?= 'print me' ?>`. this is same with `<?php echo 'print me' ?>`.
  - `<?php echo 'print me';` is also ok when the file is composed of only php code.

- php-defined variables: `$_SERVER['HTTP_USER_AGENT']` is like `Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36`

- control structure, function

- php mode allows us to skip or not html block

## Log

### first page

- how to run

```
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ php firstpage.php
<html>
 <head>
  <title>PHP Test</title>
 </head>
 <body>
 <p>Hello World</p>
 <p>Hello World, again</p>
 </body>
</html>
```

or run web server

```shell-session
php -S localhost:8080
```

and access `localhost:8080/firstpage.php` via browser

```shell-session
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ curl http://localhost:8080/firstpage.php
<html>
 <head>
  <title>PHP Test</title>
 </head>
 <body>
 <p>Hello World</p>
 <p>Hello World, again</p>
 </body>
</html>
```

- local code (no render purpose)

```shell-session
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ cat onlyphp.php
<?php

echo 'only php';
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ php onlyphp.php
only php% 
```

### Practice

- predefined variables

```shell-session
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ curl localhost:8080/predefined-variable.php
curl/7.64.1% 
```

- control structure and function

```shell-session
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ curl localhost:8080/if.php
あなたはInternet Explorerを使用していません<br />% 
```

- mixed of html and php

```shell-session
koketani: ~/g/g/k/p/p/01_tutorial (php ?)$ curl localhost:8080/mixed.php
<h3>strposがfalseを返しました</h3>
<center><b>あなたはInternet Explorerを使用していません</b></center>
```
