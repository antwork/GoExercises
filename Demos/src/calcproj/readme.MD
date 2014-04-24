GOPATH settings

1.open file: ~/.bashrc
2.add export to bashrc
export GOPATH= <Your project directory> /calcproj:$GOPATH
example: export GOPATH=/Users/apple/Documents/GoExercises/Demos/src/calcproj:$GOPATH
3.save
4.run source command 		$ ~/.bashrc
5.open calcproj  	  		$ cd <Your project directory>/calcproj
6.new directory  		    $ mkdir bin
7.open bin directory	    $ cd bin
8.build         			$ go build calc
9.if everything go well,you will get a runable file,run 	$ ./calc
USAGE:calc command [arguments] ...

The Commands are:
	add	Addition of two values.
	sqrt	Square root of a non-negetive value.


<#-----  中文版本 -----#>

GOPATH 路径配置:

1.打开文件  ~/.bashrc
2.添加export语句(因为GOPATH可以多个,所以后面添加了$GOPATH,并且以:隔开)
GOPATH= <Your project directory> /calcproj:$GOPATH
示例:export GOPATH=/Users/apple/Documents/GoExercises/Demos/src/calcproj:$GOPATH
3.保存
4.运行 source   $ ~/.bashrc
5.打开calcproj  $ cd <Your project directory>/calcproj
6.新建bin文件夹  $ mkdir bin
7.打开bin	   $ cd bin
8.编译          $ go build calc
9.如果没有出错,则在该目录下会有一个calc的可执行文件,运行 $ ./calc
USAGE:calc command [arguments] ...

The Commands are:
	add	Addition of two values.
	sqrt	Square root of a non-negetive value.





