---
title: "feedsync"
version: "0.1.0"
summary: "A CLI tool for pushing RSS/Atom feed items to your Pocket list"
link: "https://github.com/mattmeyers/feedsync"
date: 2020-10-22T22:42:07-04:00
draft: false
weight: 400
---

{{< extlink href="https://getpocket.com" title="Pocket" >}} is a news/link aggregation app that allows users to save any link for later by using their browser extension or mobile app. However, Pocket is unable to automatically listen for new items in a RSS/Atom feed. `feedsync` solves this problem by checking configured feeds for new items, and then pushing these items to Pocket via the Pocket REST API. By automating the execution of this command with a tool like `cron`, Pocket can become an automatically updated personal repository of blog posts and articles.
