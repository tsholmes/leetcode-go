SELECT
  id,
  count(*) AS num
FROM (
  SELECT
    requester_id AS id,
    accept_date
  FROM
    request_accepted
  UNION ALL
  SELECT
    accepter_id as id,
    accept_date
  FROM
    request_accepted
) rs
GROUP BY 1
ORDER BY 2 DESC
LIMIT 1
