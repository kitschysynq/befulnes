befulnes
========

[![GoDoc](http://godoc.org/github.com/kitschysynq/befulnes?status.svg)](http://godoc.org/github.com/kitschysynq/befulnes)
[![Build Status](https://travis-ci.org/kitschysynq/befulnes.svg?branch=master)](https://travis-ci.org/kitschysynq/befulnes)

Stop wasting time naming things. Names are hard, and generally not nearly as important as people make them out to be. Even when they turn out to be important, they're usually pretty easy to change.

Usage
-----

### Spit out a random name

```
> befulnes
quasce
>
```

This is very useful for incporporating into other scripts, or just pasting into chat when you need a name.


### Initialize a git repository
```
> befulnes --project
> ls
hockstod
>
```

This will generate a name, then create and initialize a git repository so you can start working right away.

Caveats
-------

Names are not actually generated but are chosen randomly from http://soybomb.com/tricks/words/ . You will need internet access to fetch the list. This is not my site, and I have no control or influence there, so I take no responsibility if it stops working.
