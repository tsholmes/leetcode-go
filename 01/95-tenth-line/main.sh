cat file.txt | awk 'BEGIN {i = 0} {i += 1; if (i == 10) print $0}'
