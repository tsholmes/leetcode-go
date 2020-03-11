SELECT
  player_id,
  min(event_date) as first_login
FROM
  Activity
GROUP BY
  1
ORDER BY
  1;
