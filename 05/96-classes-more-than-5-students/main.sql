SELECT
  class
FROM
  courses
GROUP BY
  1
HAVING
  COUNT(distinct student) >= 5
