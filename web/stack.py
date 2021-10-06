#!/usr/local/bin/python3
from stackapi import StackAPI
from datetime import datetime

query = "How to include python in golang"

SITE = StackAPI('stackoverflow')

# Parameters for search
#SITE.page_size = 10
#SITE.max_pages = 10

# Get search 
questions = SITE.fetch('similar', title='test')

print(questions)
