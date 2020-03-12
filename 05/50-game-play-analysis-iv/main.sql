SELECT
  ROUND(SUM(a2.player_id IS NOT NULL) / COUNT(*), 2) AS fraction
FROM (
  SELECT
    player_id,
    MIN(event_date) as first_login
  FROM
    Activity
  GROUP BY
    1
) a1
LEFT JOIN
  Activity a2
ON
  a1.player_id = a2.player_id AND
  a2.event_date = DATE_ADD(a1.first_login, interval 1 day)
