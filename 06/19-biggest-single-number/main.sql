SELECT
  max(num) as num
FROM (
  SELECT
    num
  FROM
    my_numbers
  GROUP BY 1
  HAVING COUNT(*) = 1
) ns
