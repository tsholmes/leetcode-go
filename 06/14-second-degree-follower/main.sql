SELECT DISTINCT
  f1.id as follower,
  COUNT(DISTINCT f2.follower) as num
FROM (
  SELECT DISTINCT
    follower as id
  FROM
    follow
) f1
JOIN
  follow f2
ON
  f1.id = f2.followee
GROUP BY 1
ORDER BY 1 ASC
