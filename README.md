# Go & Mysql With Docker 

- Test db code with docker-compose.yml file
- Start :  docker-compose --force-recreate
- Stop : docker-compose down

- Let's connect to the MySQL server inside the Docker container :
 <br> docker-compose exec database mysql -u your-username -pyour-password  (By the way, this username and password are for example. I show it for understanding.) 

 - After the connect let's show the database : SHOW databases;
  <br>

  ![code](https://github.com/Furkanturan8/docker-golang/blob/main/fotos/terminal-code-2.png)
  <br>

- Let's get inside the db : use db;
  <br>

  ![code 2](https://github.com/Furkanturan8/docker-golang/blob/main/fotos/terminal-code-3.png)
  <br>

- You can see the book database but there is no book in inside. So we need to add a new book. As you can see in the main.go code, we also wrote a new book information. Now let's run the project.

  <br>
  
  ![code 3](https://github.com/Furkanturan8/docker-golang/blob/main/fotos/terminal-code-5.png)
  <br>

- Now we have a new book. Let's check the database : select * from book;
  <br>
  
  ![code 4](https://github.com/Furkanturan8/docker-golang/blob/main/fotos/terminal-code-4.png)


  ***Yeeep! Very nice! The process is successful. Here we created a container using docker and added our image (as a mysql database).  Then we tried the create function for testing purposes.  And the addition process is successful.***
