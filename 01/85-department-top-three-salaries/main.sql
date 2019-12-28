SELECT
  Department,
  Employee,
  Salary
FROM (
  SELECT
    de.Department,
    de.Employee,
    de.Salary,
    @pos := IF(de.Department = @lastD,
      IF(de.Salary < @lastS, @pos+1, @pos),
      1
    ) as Position,
    @lastS := de.Salary,
    @lastD := de.Department
  FROM (
    SELECT
      d.Name as Department,
      e.Name as Employee,
      e.Salary as Salary
    FROM
      Employee e
    JOIN
      Department d
    ON
      e.DepartmentId = d.Id
    ORDER BY
      e.DepartmentId ASC,
      e.Salary DESC
  ) de
  CROSS JOIN (
    SELECT @pos := 1, @lastS := -1, @lastD := ""
  ) vars
) es
WHERE
  Position <= 3;
