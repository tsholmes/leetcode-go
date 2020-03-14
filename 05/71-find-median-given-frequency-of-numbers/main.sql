SELECT
  AVG(Number) AS median
FROM (
  SELECT
    Number,
    Frequency,
    @cumsum := @cumsum + Frequency as cumsum
  FROM
    Numbers
  JOIN (
    SELECT @cumsum := 0
  ) vars
  ORDER BY
    1 ASC
) n1
JOIN (
  SELECT
    SUM(Frequency) AS TotalFreq
  FROM
    Numbers
) n2
WHERE
  (FLOOR((TotalFreq+1)/2) BETWEEN cumsum - Frequency + 1 AND cumsum) OR
  (CEIL((TotalFreq+1)/2) BETWEEN cumsum - Frequency + 1 AND cumsum)
