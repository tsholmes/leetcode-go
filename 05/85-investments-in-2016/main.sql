SELECT
  SUM(TIV_2016) AS TIV_2016
FROM
  insurance
WHERE
  (LAT, LON) IN (
    SELECT
      LAT,
      LON
    FROM
      insurance
    GROUP BY
      1, 2
    HAVING
      COUNT(*) = 1
  ) AND
  TIV_2015 IN (
    SELECT
      TIV_2015
    FROM
      insurance
    GROUP BY
      1
    HAVING
      COUNT(*) > 1
  )
