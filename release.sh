#!/usr/bin/env bash
git add .
git commit -m 'release http-client'
git push origin master
git tag v1.0.4
git push origin v1.0.4