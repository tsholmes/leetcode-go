SELECT
  IF(
    (accepted_reqs / total_reqs) IS NOT NULL,
    ROUND(accepted_reqs / total_reqs, 2),
    0.00
  ) AS accept_rate
FROM (
  SELECT
    COUNT(DISTINCT requester_id, accepter_id) AS accepted_reqs
  FROM
    request_accepted
) ar
JOIN (
  SELECT
    COUNT(DISTINCT sender_id, send_to_id) AS total_reqs
  FROM
    friend_request
) tr
