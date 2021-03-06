#!/usr/local/bin/python3
import tweepy, configparser, sqlite3, re
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
    search = f"{query} -filter:retweets"
tweets = tweepy.Cursor(
    api.search_tweets, 
    q=search, 
    lang='en', 
    result_type='popular', 
    include_entities=True, 
    tweet_mode='extended').items(50)

index, limit = 0, 25 # How many rows should you add?
results = {}
for tweet in tweets:
    results[tweet.user.name] = tweet.full_text
    # Construct URL
    url = f"https://twitter.com/{tweet.user.screen_name}/status/{tweet.id}"
    # Construct Date of each tweet
    date = str(tweet.created_at)
    dateFmt = f"{date[5:7]}/{date[8:10]}"
    if tweet.user.followers_count >= 30 and tweet.user.protected == False:
        if tweet.full_text == "":
            c.execute("INSERT INTO twitterQ VALUES (?, ?, ?, ?, ?, ?, ?)", (
                tweet.user.screen_name,
                tweet.user.name,
                dateFmt,
                "MEDIACONTENTONLY",
                tweet.retweet_count,
                tweet.favorite_count,
                url
            ))
        else:
             c.execute("INSERT INTO twitterQ VALUES (?, ?, ?, ?, ?, ?, ?)", (
                tweet.user.screen_name,
                tweet.user.name,
                dateFmt,
                re.sub(r'https://t.co/\w{10}', '', tweet.full_text),
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
