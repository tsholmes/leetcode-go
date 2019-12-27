SELECT
  s.Score as Score,
  sr.Rank as Rank
FROM
  Scores s
JOIN (
  SELECT
    Score,
    @idx := @idx + 1 as Rank
  FROM (
    SELECT DISTINCT
      Score
    FROM
      Scores
    ORDER BY
      Score DESC
  ) sd
  CROSS JOIN (
    SELECT @idx := 0
  ) vars
) sr
ON
  s.Score = sr.Score
ORDER BY
  s.Score DESC;
