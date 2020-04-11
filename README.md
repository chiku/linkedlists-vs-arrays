[![Build Status](https://travis-ci.org/chiku/linkedlists-vs-arrays.svg?branch=golang)](https://travis-ci.org/chiku/linkedlists-vs-arrays)

Purpose
=======

Prove that insertions into arrays are faster than insertions into linked-lists.

Prerequisites
-------------
* Golang
* Ruby
* Bundler gem
* make
* A C++ compiler
* Boost test library

How to run the benchmarks
-------------------------

* Clone the repository
```
mkdir -p gitbub.com/chiku
cd gitbub.com/chiku
git clone https://github.com/chiku/linkedlists-vs-arrays.git -b golang
```
* Nagivate to the source code
```
cd linkedlists-vs-arrays
```
* Build the app
```
./build.sh
```
* Run the app
```
./run.sh
```

Result
------

Precomputed results are available at http://htmlpreview.github.io/?https://github.com/chiku/linkedlists-vs-arrays/blob/golang/public/index.html

License
-------
This software is released under the MIT license. Please refer to LICENSE for more details.
