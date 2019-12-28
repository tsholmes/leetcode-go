SELECT DISTINCT
  num1 as Num
FROM (
  SELECT
    n1.Num as num1,
    n2.Num as num2,
    n3.Num as num3
  FROM
    Logs n1
  JOIN
    Logs n2
  ON
    n1.Id + 1 = n2.Id
  JOIN
    Logs n3
  ON
    n1.Id + 2 = n3.Id
) n
WHERE
  num1 = num2 AND num2 = num3;
