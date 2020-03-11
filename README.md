# simple_mongodb

You can run this app by 2 ways:
  1) Run MongoDB separately in Docker container and app by yourself
  2) Run MongoDB and app by docker-compose
 
## 1) Run MongoDB separately in Docker container and app by yourself
  * Run mongo using command `make mongo_docker`
  * First run app for mock generation 
    1) Set env `GENERATE_MOCK=true`
    2) `MOCK_MIN_GAMES_COUNT=` (minimum games count for each user)
    3) `MOCK_MAX_GAMES_COUNT=` (maximum games count for each user)
    4) `MOCK_USERS_COUNT=` (users count)
    5) `MOCK_GAMES_INSERT_BATCH_SIZE=` (size of users games will be inserted in one insertion (MOCK_GAMES_INSERT_BATCH_SIZE * average(MOCK_MIN_GAMES_COUNT,MOCK_MAX_GAMES_COUNT)))
  * Run app by command `go run main.go`
  * When mock will be generated you will see lines like:
      * `time="2020-03-08T21:58:11Z" level=info msg="Inserted users len: 4000"`
      * `time="2020-03-08T21:58:11Z" level=info msg="Mock generated. Please rerun the app with GENERATE_MOCK=false env"`
      * `time="2020-03-08T21:58:11Z" level=info msg="Connection to MongoDB closed."`
  * Then set env `GENERATE_MOCK=false` and run `go run main.go`
    
## 2) Run all in Docker containers
  * Set envs
    1) Set env `GENERATE_MOCK=true`
    2) `MOCK_MIN_GAMES_COUNT=` (minimum games count for each user)
    3) `MOCK_MAX_GAMES_COUNT=` (maximum games count for each user)
    4) `MOCK_USERS_COUNT=` (users count)
    5) `MOCK_GAMES_INSERT_BATCH_SIZE=` (size of users games will be inserted in one insertion (MOCK_GAMES_INSERT_BATCH_SIZE * average(MOCK_MIN_GAMES_COUNT,MOCK_MAX_GAMES_COUNT)))
   * Run command `make all_docker`
   * Wait when mock will be generated (you can check logs with command `docker container logs {container-name}`)
   * When mock will be generated you will see lines like:
      * `time="2020-03-08T21:58:11Z" level=info msg="Inserted users len: 4000"`
      * `time="2020-03-08T21:58:11Z" level=info msg="Mock generated. Please rerun the app with GENERATE_MOCK=false env"`
      * `time="2020-03-08T21:58:11Z" level=info msg="Connection to MongoDB closed."`
   *  Then set env `GENERATE_MOCK=false` and run `make all_docker`
   
  ## After application run you can get access to it's API by address `localhost:3034`
     1) Get list of users (with pagination) - `http://localhost:3034/users?page=1&count=1000`
     2) Get rating of users (with pagination) - `http://localhost:3034/games/rating?page=1&count=1000`
     3) Get games grouped by days and number of game - `http://localhost:3034/games/group`
     
If grouping by days or number has bad performance you can try run method working with aggregations
(uncomment lines 19 and 43 in file `server/api/rest/game.go` and comment lines 20 and 44)
