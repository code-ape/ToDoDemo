
To-Do App Demo
==============

This application has three parts:

1. A webserver built with Golang and the Gin framework.
* It has an API.
* Serves static files.
2. A UI built with Bootstrap.
3. A javascript client system built with AngularJS that powers the client and calls the API.

### Server

To build the server:
* Install Golang.
* Install gom `go get github.com/mattn/gom`.
* Go to the head file directory of the repo.
* Install dependencies `gom install`.
* Build server `gom build`.

To run server:
* `./ToDoDemo`


### Client


Keynotes:
* The app is **ONE** page, AngularJS views are used.
* Bootstrap is used for styling but to Javascript from it is used.
