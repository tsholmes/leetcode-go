SELECT
  e1.pay_month,
  e1.department_id,
  IF(e1.avg > e2.avg, 'higher', IF(e1.avg < e2.avg, 'lower', 'same')) AS comparison
FROM (
  SELECT
    DATE_FORMAT(s.pay_date, '%Y-%m') AS pay_month,
    e.department_id AS department_id,
    AVG(s.amount) AS avg
  FROM
    salary s
  JOIN
    employee e
  ON
    s.employee_id = e.employee_id
  GROUP BY 1, 2
) e1
JOIN (
  SELECT
    DATE_FORMAT(pay_date, '%Y-%m') AS pay_month,
    AVG(amount) AS avg
  FROM
    salary
  GROUP BY 1
) e2
ON
  e1.pay_month = e2.pay_month
