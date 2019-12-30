SELECT
  Request_at as Day,
  ROUND(
    SUM(Status IN ('cancelled_by_driver', 'cancelled_by_client')) / COUNT(*),
    2
  ) as `Cancellation Rate`
FROM
  Trips t
JOIN
  Users u1
ON
  t.Client_Id = u1.Users_Id
JOIN
  Users u2
ON
  t.Driver_Id = u2.Users_Id
WHERE
  u1.Banned = 'No' AND
  u2.Banned = 'No' AND
  t.Request_at >= '2013-10-01' AND
  t.Request_at <= '2013-10-03'
GROUP BY
  Request_at
