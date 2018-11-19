# reddit_api
# Assignment 3 (Group) in IMT2681 Cloud Technologies

## Descriptions

## API
### GET: localhost8080:/reddit/
Redirects to localhost8080:/reddit/api/

### GET: localhost8080:/reddit/api/
Uptime of the service
```json
{
  "uptime": "<uptime>",
  "info": "Reddit api",
  "version": "v1" 
}
```
  
### GET: localhost:8080/reddit/api/me/
* What: Gets the user info thats connected to reddit.

* Response:

```json
{
    "id": "27lh22f3",
    "name": "EnvironmentalDonkey1",
    "created": 1541945447,
    "Karma": {
        "comment_karma": 0,
        "link_karma": 0
    },
    "url": ""
}
```

### GET: localhost:8080/reddit/api/me/karma/
* What: Gets the karma of the user.

* Response:

```json
{
  "comment_karma": "<int>",
  "link_karma": "<int>",
}
```

### GET: localhost:8080/reddit/api/me/friends/
* What: Gets all the friends of the user.

* Response:

```json
{
	"date": "<float32>",
	"name": "<string>",
	"id": "<string>",
}
```

### GET: localhost:8080/reddit/api/me/prefs/
* What: Get the preferences of the user.

* Response:

```json
{
 
	"research": "<bool>",
	"show_trending": "<bool>",
	"private_feeds": "<bool>",
	"ignore_suggested_sort": "<bool>",
	"over_18": "<bool>",
	"email_messages": "<bool>",
	"force_https": "<bool>",
	"lang": "<string>",
	"hide_from_robots": "<bool>",
	"public_votes": "<bool>",
	"hide_ads": "<bool>",
	"beta": "<bool>",
}
```


### POST: localhost:8080/reddit/api/submission/
* What: POST a submission //FILL IN

* Body:
```json
{
  "title": "<string>",
	"author": "<string>",
	"subreddit": "<string>",
	"name": "<string>",
	"numComments": "<string>",
	"score": "<string>",
}
```

### GET: localhost:8080/reddit/api/{username}/karma/
* What: Get the karma of an arbitrary user

* Response:

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/{cap}/frontpage/{sortby}/
* What: Get //FILL IN

 {cap} - <int>  that specifies how many posts to be received
 {sortby} - <"string">: new, best, top, rising, hot, controversial

* Response:

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/subreddit/{subreddit}/{sortby}/{cap}/
* What: Get //FILL IN

 {subreddit} - <"string"> - e.g "r/soccer"
 {cap} - <int>  that specifies how many posts to be received
 {sortby} - <"string">: new,best,top,rising,hot,controversial

* Response:
```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/comments/{submission}/{cap}/
* What: Get a specific amount of submissions

* {submission} - <"string"> "cat"
* {cap} - <int>  that specifies how many posts to be received

* Response:

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/{username}/user/
* What: Get information of an specific user

 {username} - <"string">

* Response:

```json
{
  fill in
}
```

## Admin API

### GET: localhost:8080/reddit/api/admin/users/{username}/{pwd}/
* What: Returns every user in the database

 {username} - <"string"> admin username
 {pwd} - <"string"> Pre- specified admin token

* Response:

```json
{
  fill in
}
```


### GET: localhost:8080/reddit/api/admin/user/{id}/{username}/{pwd}/
* What: Returns a specific user from the database

 {id} - <"string">
 {username} - <"string">
 {pwd} - <"string"> Pre- specified token

* Response:

```json
{
  fill in
}
```

### DELETE: localhost:8080/reddit/api/admin/delete/{id}/{username}/{pwd}/
* What: Deletes a specific user in the database

 {username} - <"string">
 {pwd} - <"string"> Pre- specified token

* Response:

```
 fill in  
```

### DELETE: localhost:8080/reddit/api/admin/wipe/{username}/{pwd}/
* What: Deletes every user in the database

 {username} - <"string">
 {pwd} - <"string"> Pre- specified token

* Response:

```
  Output: Either <"successful"> or <"failed">
```


## Webhook

### POST: localhost:8080/reddit/api/webhook/new/
* What: Creates a new webhook

* Body:

```json
{
  fill in
}
```
