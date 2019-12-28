SELECT
  d.Name as Department,
  e.Name as Employee,
  e.Salary as Salary
FROM
  Employee e
JOIN (
  SELECT
    de.Id,
    de.Name,
    MAX(ee.Salary) as MaxSalary
  FROM
    Employee ee
  JOIN
    Department de
  ON
    ee.DepartmentId = de.Id
  GROUP BY
    de.Id, de.Name
) d
ON
  e.DepartmentId = d.Id
WHERE
  e.Salary = d.MaxSalary;
