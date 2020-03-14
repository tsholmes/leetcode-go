SELECT
  c.Name as Name
FROM
  Candidate c
JOIN (
  SELECT
    CandidateId,
    COUNT(*) as vcount
  FROM
    Vote
  GROUP BY 1
) v
ON
  c.id = v.CandidateId
ORDER BY
  v.vcount DESC
LIMIT 1
