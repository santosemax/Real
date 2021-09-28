#!/usr/local/bin/python3
import tweepy, configparser, sqlite3
from os import path

Config = configparser.ConfigParser()
Config.read("./botConfigs/twitterBot.ini")
consumer_key = Config['twitter']['apiKey']
consumer_secret = Config['twitter']['apiSecretKey']
# OAuth 2 Authentication (May change this to 1a, not sure)
auth = tweepy.OAuthHandler(consumer_key, consumer_secret)
api = tweepy.API(auth, wait_on_rate_limit=True)


# DB init
conn = sqlite3.connect('results.db')
c = conn.cursor()

# Searching twitter using query (from db)
search = ""
for (query,) in c.execute("SELECT query FROM queryQ"):
    search = query
postLimit = 50 # How many results should the API search for?
tweets = tweepy.Cursor(api.search_tweets, q=search, lang='en', result_type='recent', include_entities=True).items(10)
index = 0
limit = 10 # How many rows should you add?
results = {}
for tweet in tweets:
    results[tweet.user.name] = tweet.text
    # Construct URL
    url = f"https://twitter.com/{tweet.user.screen_name}/status/{tweet.id}"
    if tweet.user.followers_count >= 30 and tweet.user.protected == False:
        if tweet.text == "":
            c.execute("INSERT INTO twitterQ VALUES (?, ?, ?, ?, ?, ?, ?)", (
                tweet.user.screen_name,
                tweet.user.name,
                tweet.created_at,
                "MEDIACONTENTONLY",
                tweet.retweet_count,
                tweet.favorite_count,
                url
            ))
        else:
             c.execute("INSERT INTO twitterQ VALUES (?, ?, ?, ?, ?, ?, ?)", (
                tweet.user.screen_name,
                tweet.user.name,
                tweet.created_at,
                tweet.text,
                tweet.retweet_count,
                tweet.favorite_count,
                url
            ))
        index += 1
    if index > limit:
        break


# Printing only keys (FOR DEBUG)
print("\n\n")
keys = list(results.keys())
counter = 0
for item in keys:
    print(item)
    counter += 1
print(f"The number of items = {counter} \n\n")

conn.commit()
conn.close()
