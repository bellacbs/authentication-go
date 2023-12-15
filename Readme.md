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
* jwt

<h2 id="funciona">:heavy_check_mark: What works</h2>

* Model, view, controller, repository;</br>
* Login User;</br>
* Create User;</br>
* Get User By ID;</br>
* Get User By Email;</br>
* Delete User;</br>
* Middleware for authentication;</br>
* Docker for mongoDB;</br>
* Docker image for dev;</br>
* jwt authentication;</br>


 
<h2 id="pendente">:construction: In development</h2>

- [x] tests;
- [x] Docker image for deploy;

<h2 id="requirements">:leftwards_arrow_with_hook: Prerequisites</h2>

Before you start, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Go](https://go.dev/doc/install), [Docker](https://docs.docker.com/engine/install/). 
Additionally, it's good to have a code editor to work with, such as [VSCode](https://code.visualstudio.com/)


<h4>:checkered_flag: Running the project </h4>

```bash
# Clone this repository

# To start
$ docker compose up


# The server will start on port 8080 - <http://localhost:8080>
# To see mongoDB database using mongo-express, can use por 8081, username and password = admin - <http://localhost:8081>

```


<a href="#top">Go back to the top</a>
