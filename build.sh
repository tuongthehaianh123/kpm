#!/bin/bash



export DEFAULTPATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
export PATH=$DEFAULTPATH:/usr/local/go/bin:/opt/go/bin:/root/go/bin



cd e2ap && \
gcc -c -fPIC -Iheaders/ lib/*.c wrapper.c && \
gcc *.o -shared -o libe2apwrapper.so && \
cp libe2apwrapper.so /usr/local/lib/ && \
mkdir /usr/local/include/e2ap && \
cp wrapper.h headers/*.h /usr/local/include/e2ap && \
ldconfig



cd ..



cd e2sm && \
gcc -c -fPIC -Iheaders/ lib/*.c wrapper.c -lm && \
gcc *.o -shared -o libe2smwrapper.so -lm && \
cp libe2smwrapper.so /usr/local/lib/ && \
mkdir /usr/local/include/e2sm && \
cp wrapper.h headers/*.h /usr/local/include/e2sm && \
ldconfig



cd ..
go build ./kpimon.go