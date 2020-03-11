SELECT
  a1.player_id,
  a1.device_id
FROM
  Activity a1
JOIN (
  SELECT
    player_id,
    MIN(event_date) as first_login
  FROM
    Activity
  GROUP BY
    1
) a2
ON
  a1.player_id = a2.player_id and a1.event_date = a2.first_login
ORDER BY
  1 ASC;
