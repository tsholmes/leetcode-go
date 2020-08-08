SELECT DISTINCT
  s.id,
  s.visit_date,
  s.people
FROM
  stadium s
LEFT JOIN (
  SELECT
    s1.id as id
  FROM
    stadium s1
  INNER JOIN
    stadium s2
  ON
    s2.id = s1.id - 1
  INNER JOIN
    stadium s3
  ON
    s3.id = s1.id + 1
  WHERE
    s1.people >= 100 AND
    s2.people >= 100 AND
    s3.people >= 100
) ss
ON
  s.id = ss.id OR
  s.id + 1 = ss.id OR
  s.id - 1 = ss.id
WHERE
  s.people >= 100 AND
  ss.id IS NOT NULL
