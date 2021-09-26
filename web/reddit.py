#!/usr/local/bin/python3
import praw, configparser, sqlite3
from os import path


# Reddit config
Config = configparser.ConfigParser()
Config.read("./botConfigs/redditBot.ini")
client = Config['reddit']['client']
secret = Config['reddit']['secret']
agent = Config['reddit']['agent']

reddit = praw.Reddit(
    client_id=client,
    client_secret=secret,
    user_agent=agent,
)

# DB init
conn = sqlite3.connect('results.db')
c = conn.cursor()

# Searching reddit using query (from db)
search = ""
for (query,) in c.execute("SELECT query FROM queryQ"):
    search = query
postLimit = 50 # How many results should the API search for?
search_posts = reddit.subreddit('all').search(search, limit=postLimit)
index = 0
limit = 10 # How many rows should you add?
results = {}
for post in search_posts:
    results[post.title] = post.selftext
    if post.over_18 != True:
        if post.selftext == "":
            c.execute("INSERT INTO results VALUES (?, ?, ?, ?, ?)", (
                post.title, 
                "noCONTENT", 
                post.url, 
                str(post.subreddit), 
                post.permalink)
            )
        else:
            c.execute("INSERT INTO results VALUES (?, ?, ?, ?, ?)", (
                post.title, 
                post.selftext, 
                post.url, 
                str(post.subreddit), 
                post.permalink)
            )
        index += 1 
    if index > limit:
        break

#c.execute("SELECT * FROM results")
#print(c.fetchmany(10))

#print(results)


# Printing only keys (FOR DEBUG)
keys = list(results.keys())
counter = 0
for item in keys:
    print(item)
    counter += 1
print(f"The number of items = {counter} \n\n")

conn.commit()
conn.close()
