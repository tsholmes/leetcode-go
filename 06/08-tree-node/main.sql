SELECT
  t.id AS id,
  IF(tp.id IS NULL, 'Root', IF(MIN(tc.id) IS NULL, 'Leaf', 'Inner')) AS Type
FROM
  tree t
LEFT JOIN
  tree tp
ON
  t.p_id = tp.id
LEFT JOIN
  tree tc
ON
  t.id = tc.p_id
GROUP BY t.id, tp.id
