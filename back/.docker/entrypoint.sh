#!/bin/bash

#if no env. file, create one based on .env.example file:
if [ ! -f ".env" ]; then
  cp .env.example .env
fi

npm install

npm run start:dev