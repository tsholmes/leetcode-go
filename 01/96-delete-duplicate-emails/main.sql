DELETE
  p1
FROM
  Person AS p1
WHERE
  p1.Id NOT IN (
    SELECT
      MIN(p2.Id)
    FROM
      Person p2
    GROUP BY
      p2.Email
  );
