SELECT
  s.name
FROM
  salesperson s
LEFT JOIN
  company c
ON
  c.name = 'RED'
LEFT JOIN
  orders o
ON
  o.com_id = c.com_id AND
  o.sales_id = s.sales_id
WHERE
  o.order_id IS NULL
