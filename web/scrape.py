import praw, configparser
from os import path

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

# Searching SQLite3 Tutorials (Have search terms from form be here)
search_posts = reddit.subreddit('all').search("Learning SQLite3")
index = 0
limit = 10
for post in search_posts:
    index += 1 
    print(post.title)
    print(post.selftext)
    for i in range(50):
        print("-", end="")
    print()
    if index > limit:
        break
