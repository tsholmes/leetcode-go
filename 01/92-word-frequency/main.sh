cat words.txt | xargs -n 1 echo | sort | uniq -c | sort -rn | awk '{print $2 " " $1}'
