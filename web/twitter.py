#!/usr/local/bin/python3
import tweepy, configparser, sqlite3
from os import path

Config = configparser.ConfigParser()
Config.read("./botConfigs/twitterBot.ini")
consumer_key = Config['twitter']['apiKey']
consumer_secret = Config['twitter']['apiSecretKey']

# OAuth 2 Authentication (May change this to 1a, not sure)
auth = tweepy.OAuthHandler(consumer_key, consumer_secret)
api = tweepy.API(auth)

for tweet in tweepy.Cursor(api.search_tweets, q='Ronaldo', lang='en', result_type='recent').items(10):
    print(tweet.text)


# DB init
conn = sqlite3.connect('results.db')
c = conn.cursor()

# Searching twitter using query (from db)
