SELECT
  w1.Id
FROM
  Weather w1
JOIN
  Weather w2
ON
  w1.RecordDate = DATE_ADD(w2.RecordDate, interval 1 day)
WHERE
  w1.Temperature > w2.Temperature;
