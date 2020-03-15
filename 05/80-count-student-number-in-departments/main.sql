SELECT
  d.dept_name AS dept_name,
  COALESCE(s.scount, 0) AS student_number
FROM
  department d
LEFT JOIN (
  SELECT
    dept_id,
    COUNT(*) as scount
  FROM
    student
  GROUP BY 1
) s
ON
  d.dept_id = s.dept_id
ORDER BY
  2 DESC,
  1 ASC
