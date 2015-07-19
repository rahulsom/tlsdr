#!/bin/bash
set -e

gom install && cd tlsdr && gom build
cd ..

if [[ $TRAVIS_BRANCH == 'master' && $TRAVIS_REPO_SLUG == "rahulsom/tlsdr" \
          && $TRAVIS_PULL_REQUEST == 'false' ]]; then
    cat Gomfile| cut -d " " -f 2 | xargs -n 1 go get
    cd tlsdr
    gox -cgo -verbose
else
  echo "Not on master branch, so not publishing"
  echo "TRAVIS_BRANCH: $TRAVIS_BRANCH"
  echo "TRAVIS_REPO_SLUG: $TRAVIS_REPO_SLUG"
  echo "TRAVIS_PULL_REQUEST: $TRAVIS_PULL_REQUEST"
fi