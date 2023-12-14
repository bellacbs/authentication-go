<h1 align="center" id="top">Authentication - Golang</h1>

<p align="center">
  <a href="#sobre">About</a> &#xa0; | &#xa0; 
  <a href="#funciona">What works</a> &#xa0; | &#xa0;
  <a href="#pendente">In development</a> &#xa0; | &#xa0;
  <a href="#requirements">Requirements</a> &#xa0; | &#xa0;
</p>

<h2 id="sobre">:notebook: About </h2>

<p align="center">:rocket: Project to practice MVC architecture using Golang and MongoDB</p>

<h2 id="tecnologias"> 🛠 Technologies and programming languages </h2>

The following libraries and languages were used in the project's construction:

* Go
* Gin-gonic
* Makefile
* Traslator
* Validator
* Docker Compose whith Mongodb
* zap package from uber

<h2 id="funciona">:heavy_check_mark: What works</h2>

* Model, view, controller, repository;</br>
* createUser;</br>

 
<h2 id="pendente">:construction: In development</h2>

- [x] jwt token;
- [x] Get  user;
- [x] Update User;
- [x] Delete User;

<h2 id="requirements">:leftwards_arrow_with_hook: Prerequisites</h2>

Before you start, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Go](https://go.dev/doc/install). 
Additionally, it's good to have a code editor to work with, such as [VSCode](https://code.visualstudio.com/)


<h4>:checkered_flag: Running the project </h4>

```bash
# Clone this repository

#set variables
PORT
LOG_OUTPUT
LOG_LEVEL
COST
MONGODB_URL
MONGODB_AUTH_DB
MONGODB_USER_COLLECTION

# To start
$ go run *.go


# The server will start on port that you chose - ex: <http://localhost:8080>

```


<a href="#top">Go back to the top</a>