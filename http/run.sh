#!/bin/sh

dowork()
{
	stopsv $1 $2
	startsv $1 $2
}

startsv()
{
	server=$1
	config=$2
	echo "starting $server"
	$PWD/$server -d -c $config
	sleep 1
	ps x|grep $server |sed -e '/grep/d'
}
stopsv()
{
	server=$1
	ps x |grep $PWD/$server| sed -e '/grep/d' | gawk '{print "panic."$1}' | xargs rm -v
	ps x |grep $PWD/$server| sed -e '/grep/d' | gawk '{print $1}' | xargs kill -9
	echo "stop $server"
}

echo "--------------------------------------------------"
echo "--------------------START-------------------------"
echo "--------------------------------------------------"

dowork httpunilighttest

echo "--------------------------------------------------"
echo "----------------------DONE------------------------"
echo "--------------------------------------------------"


