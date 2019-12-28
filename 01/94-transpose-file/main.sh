COUNT=$(head -n 1 file.txt | wc -w)

for ((i = 1; i <= ${COUNT}; i++))
do
  cat file.txt | awk "{ print \$${i}}" | xargs echo
done
