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
Gets the user info thats connected to reddit.
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
Gets the karma of the user.

```json
{
  "comment_karma": "<int>",
  "link_karma": "<int>",
}
```

### GET: localhost:8080/reddit/api/me/friends/
Gets all the friends of the user.

```json
{
	"date": "<float32>",
	"name": "<string>",
	"id": "<string>",
}
```

### GET: localhost:8080/reddit/api/me/prefs/
Get the preferences of the user.

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
POST a submission //FILL IN

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
Get the karma of an arbitrary user

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/{cap}/frontpage/{sortby}/
Get //FILL IN

**{cap} - \<int\>  that specifies how many posts to be received
**{sortby} - \<"string"\>: new,best,top,rising,hot,controversial

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/subreddit/{subreddit}/{sortby}/{cap}/
Get //FILL IN

**{subreddit} - \<"string"\> - e.g "r/soccer"
**{cap} - \<int\>  that specifies how many posts to be received
**{sortby} - \<"string"\>: new,best,top,rising,hot,controversial

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/comments/{submission}/{cap}/
Get a specific amount of submissions

**{submission} - \<"string"\> "cat"
**{cap} - \<int\>  that specifies how many posts to be received

```json
{
  fill in
}
```

### GET: localhost:8080/reddit/api/{username}/user/
Get information of an specific user

**{username} - \<"string"\> 

```json
{
  fill in
}
```

## Admin API

### GET: localhost:8080/reddit/api/admin/users/{username}/{pwd}/
Returns every user in the database

**{username} - \<"string"\> admin username
**{pwd} - \<"string"\> Pre- specified admin token 

```json
{
  fill in
}
```


### GET: localhost:8080/reddit/api/admin/user/{id}/{username}/{pwd}/
Returns a specific user from the database

**{id} - \<"string"\> 
**{username} - \<"string"\> 
**{pwd} - \<"string"\> Pre- specified token 

```json
{
  fill in
}
```

### DELETE: localhost:8080/reddit/api/admin/delete/{id}/{username}/{pwd}/
Deletes a specific user in the database

**{username} - \<"string"\> 
**{pwd} - \<"string"\> Pre- specified token 

```json
{
  fill in
}
```

### DELETE: localhost:8080/reddit/api/admin/wipe/{username}/{pwd}/
Deletes every user in the database

**{username} - \<"string"\> 
**{pwd} - \<"string"\> Pre- specified token 

```
  Output: Either <"successful"> or <"failed">
```


## Webhook

### POST: localhost:8080/reddit/api/webhook/new/
Creates a new webhook

```json
{
  fill in
}
```
