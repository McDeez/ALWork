#!/bin/sh
# Simple Integration level testing to ensure good queries go through
# and bad ones do not. 
#
# Usage: ./integrationTests.sh

host="localhost"
port="8080"

declare -a listOfGoodURLs=(
  "curl -I -s $host:$port"
  "curl -I -s $host:$port/"
  "curl -I -s $host:$port/name/nathan/"
  "curl -I -s $host:$port/name/nathan/test/this/long/path/too"
  "curl -I -s $host:$port/name/nathan/test/this/long/path/too/"
  "curl -I -s $host:$port/?name=nathan"
  "curl -I -s $host:$port/?name="
  "curl -I -s $host:$port/?name=name"
  "curl -I -s $host:$port/test/?name=name"
  "curl -I -s $host:$port/?name=nathan&name=nathan&name=Nathan"
)

declare -a listOfBadURLs=(
  "curl -I -s $host:$port/?name=name%27"                                            #name'
  "curl -I -s $host:$port/?name=name%27%20OR%20%27a%27%3D%27a"                      #name' OR 'a'='a
  "curl -I -s $host:$port/?name=name%27%20OR%20%27b%27%3D%27b"                      #name' OR 'b'='b
  "curl -I -s $host:$port/?name=name%27%20OR%20%27bob%27%3D%27bob"                  #name' OR 'bob'='bob
  "curl -I -s $host:$port/?name=name%27)%3B%20DELETE%20FROM%20items%3B%20--"        #name'); DELETE FROM items; --
  "curl -I -s $host:$port/?name=name%27)%3B%20DELETE%20FROM%20users%3B%20--"        #name'); DELETE FROM users; --
  "curl -I -s $host:$port/?name=name%27)%3B%20DELETE%20FROM%20people%3B%20--"       #name'); DELETE FROM people; --
)

declare -i passed=0
declare -i failed=0
declare -i total=0

testGoodURLS(){
  for url in "${listOfGoodURLs[@]}"
  do
    total=$total+1
    echo
    echo "Testing $url"
    ret=`$url`
    if ! echo "$ret" | grep -q "HTTP/1.1 200" ; then
      failed=$failed+1
      echo "\tFAILED: $ret" 
    else
      passed=$passed+1
      echo "\tPASSED"
    fi
  done
  echo "Done Valid tesing"
  echo
}


testBadURLS(){
  for url in "${listOfBadURLs[@]}"
  do
    total=$total+1
    echo
    echo "Testing $url"
    ret=`$url`
    if echo "$ret" | grep -q "HTTP/1.1 200" ; then
      failed=$failed+1
      echo "\tFAILED: $ret" 
    else
      passed=$passed+1
      echo "\tPASSED"
    fi
  done
  echo "Done Invalid testing"
}

printSummary () {
  echo
  echo "Summary:"
  echo 
  echo "Total: $total"
  echo "Passed: $passed"
  echo "Failed: $failed"
}


##### MAIN #####
testGoodURLS
testBadURLS
printSummary
