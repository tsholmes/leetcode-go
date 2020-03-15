SELECT
  e1.Id,
  e1.Month,
  (
    SELECT
      SUM(Salary)
    FROM
      Employee
    WHERE
      Id = e1.Id AND
      (Month BETWEEN e1.Month-2 AND e1.Month)
  ) AS Salary
FROM
  Employee e1
JOIN (
  SELECT
    Id,
    MAX(Month) AS MaxMonth
  FROM
    Employee
  GROUP BY
    1
) e2
ON
  e1.Id = e2.Id
WHERE
  e1.Month < e2.MaxMonth
ORDER BY
  1 ASC, 2 DESC
