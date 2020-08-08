SELECT
  c1.seat_id
FROM
  cinema c1
LEFT JOIN
  cinema c2
ON
  c2.seat_id = c1.seat_id - 1
LEFT JOIN
  cinema c3
ON
  c3.seat_id = c1.seat_id + 1
WHERE
  c1.free = 1 AND
  (c2.free = 1 OR c3.free = 1)
