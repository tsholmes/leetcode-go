SELECT
  e1.Name
FROM
  Employee e1
JOIN (
  SELECT
    ManagerId,
    COUNT(*) as ecount
  FROM
    Employee
  WHERE
    ManagerId IS NOT NULL
  GROUP BY
    1
) e2
ON
  e1.Id = e2.ManagerId
WHERE
  e2.ecount >= 5
