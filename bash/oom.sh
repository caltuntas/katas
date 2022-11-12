#!/bin/bash

printf "%-50s %8s %8s\n" "Image" "PID" "OOM_SCORE"
docker ps --format "{{.ID}} {{.Image}}" |
while read -r line; do
  id=$(echo "$line" | cut -d ' ' -f1)
  img=$(echo "$line" | cut -d ' ' -f2)
  topResult=$(docker top "$id" -o pid | grep -v "PID")

  for r in $topResult;
  do
    score=$(cat /proc/"$r"/oom_score)
    printf "%-50s %8s %8s\n" "$img" "$r" "$score"
  done
done | sort -k3 -n -r


