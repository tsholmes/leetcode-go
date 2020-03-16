SELECT
  customer_number
FROM
  orders
GROUP BY
  1
ORDER BY
  COUNT(*) DESC
LIMIT
  1
