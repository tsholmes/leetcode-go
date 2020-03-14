SELECT
  e1.Id,
  e1.Company,
  e1.Salary
FROM (
  SELECT
    Id,
    Company,
    Salary,
    @idx := IF(Company = @lastCompany, @idx + 1, 1) as idx,
    @lastCompany := Company
  FROM (
    SELECT
      Id,
      Company,
      Salary
    FROM
      Employee
    ORDER BY
      2, 3 ASC
  ) ee
  JOIN (
    SELECT
      @idx := 0,
      @lastCompany := ""
  ) vars
) e1
JOIN (
  SELECT
    Company,
    COUNT(*) as ecount
  FROM
    Employee
  GROUP BY
    1
) e2
ON
  e1.Company = e2.Company
WHERE
  e1.idx = FLOOR((e2.ecount + 1) / 2) OR
  e1.idx = CEIL((e2.ecount + 1) / 2)
