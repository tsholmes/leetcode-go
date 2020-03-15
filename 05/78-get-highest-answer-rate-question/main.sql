SELECT
  question_id as survey_log
FROM
  survey_log
WHERE
  answer_id IS NOT NULL
GROUP BY 1
ORDER BY COUNT(*) DESC
LIMIT 1
