```
diff -u orig.yml new.yml > diff.patch
patch -u -o new.yml orig.yml < diff.patch
```
