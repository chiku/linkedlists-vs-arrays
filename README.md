[![Build Status](https://travis-ci.org/chiku/linkedlists-vs-arrays.svg?branch=cpp)](https://travis-ci.org/chiku/linkedlists-vs-arrays)
[![Build Status](https://travis-ci.org/chiku/linkedlists-vs-arrays.svg?branch=golang)](https://travis-ci.org/chiku/linkedlists-vs-arrays)

Purpose
=======

Prove that insertions into arrays are faster than insertions into linked-lists.

Prerequisites
-------------
* Ruby
* Bundler gem
* make
* A C++ compiler
* Boost test library
* lcov

How to run the benchmarks
-------------------------

### C++

* Clone the repository
```
git clone https://github.com/chiku/linkedlists-vs-arrays.git -b cpp
```
* Nagivate to the source code
```
cd linkedlists-vs-arrays
```
* Fetch the prerequisites
```
./prereqs.sh # Ubuntu or Debian are supported
# Please adjust this script if you are using another operating-system
```
* Build the app
```
./build.sh
```
* Run the app
```
./run.sh
```

### golang

* Clone the repository
```
cd $GOROOT
mkdir -p gitbub.com/chiku
cd gitbub.com/chiku
git clone https://github.com/chiku/linkedlists-vs-arrays.git -b golang
```
* Nagivate to the source code
```
cd linkedlists-vs-arrays
```
* Fetch the prerequisites
```
./prereqs.sh
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

Precomputed results are available at http://htmlpreview.github.io/?https://github.com/chiku/linkedlists-vs-arrays/blob/cpp/public/index.html
and at http://htmlpreview.github.io/?https://github.com/chiku/linkedlists-vs-arrays/blob/golang/public/index.html


License
-------
This software is released under the MIT license. Please refer to LICENSE for more details.
