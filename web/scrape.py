#!/usr/local/bin/python3
import praw, configparser, sqlite3
from os import path


# Reddit config
basepath = path.dirname(__file__)
filepath = path.abspath(path.join(basepath, "..", "redditBot.ini"))

Config = configparser.ConfigParser()
Config.read(filepath)
client = Config['reddit']['client']
secret = Config['reddit']['secret']
agent = Config['reddit']['agent']

reddit = praw.Reddit(
    client_id=client,
    client_secret=secret,
    user_agent=agent,
)

# Reddit Class
class RedditResults:
    """Handles the results from Reddit API"""

    def __init__(self, title, body):
        self.title = title
        self.body = body

    @property
    def complete(self):
        return '{} -*-*-*- {}'.format(self.title, self.body)

    def __repr__(self):
        return "Result('{}', '{}')".format(self.first, self.last)



# DB init
conn = sqlite3.connect('results.db')
c = conn.cursor()
# Create Table
#c.execute("""CREATE TABLE results (
#            title text,
#            body text
#            )""")


# Searching SQLite3 Tutorials (Have search terms from form be here)
search_posts = reddit.subreddit('all').search("Learning SQLite3")
index = 0
limit = 20
results = {}
# Load Results into a dictionary
for post in search_posts:
    results[post.title] = post.selftext
    c.execute("INSERT INTO results VALUES (?, ?)", (post.title, post.selftext))
    index += 1 
    if index > limit:
        break

# Debug Line
#c.execute("DELETE FROM results")

c.execute("SELECT * FROM results")
#print(c.fetchmany(10))

#print(results)

# Printing only keys
keys = list(results.keys())
counter = 0
for item in keys:
    print(item)
    counter += 1
print(f"\nThe number of items = {counter}")



conn.commit()
conn.close()
