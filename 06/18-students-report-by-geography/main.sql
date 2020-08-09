SELECT
  s1.name AS America,
  s2.name AS Asia,
  s3.name AS Europe
FROM (
  SELECT
    @idx := @idx + 1 AS idx,
    s_am.name
  FROM (
    SELECT
      name
    FROM
      student
    WHERE
      continent = 'America'
    ORDER BY 1 ASC
  ) s_am
  JOIN ( SELECT @idx := 0 ) v_am
) s1
LEFT JOIN (
  SELECT
    @idx1 := @idx1 + 1 AS idx,
    s_as.name
  FROM (
    SELECT
      name
    FROM
      student
    WHERE
      continent = 'Asia'
    ORDER BY 1 ASC
  ) s_as
  JOIN ( SELECT @idx1 := 0 ) v_as
) s2
ON s1.idx = s2.idx
LEFT JOIN (
  SELECT
    @idx2 := @idx2 + 1 AS idx,
    s_eu.name
  FROM (
    SELECT
      name
    FROM
      student
    WHERE
      continent = 'Europe'
    ORDER BY 1 ASC
  ) s_eu
  JOIN ( SELECT @idx2 := 0 ) v_eu
) s3
ON s1.idx = s3.idx
