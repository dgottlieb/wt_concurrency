#!/bin/bash

clang++ -L/Users/dgottlieb/xgen/wiredtiger/.libs/ -ggdb -std=c++14 -fPIC ./wt_sequence.cpp -l wiredtiger && ./a.out
