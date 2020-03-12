SELECT
  player_id,
  event_date,
  games_played_so_far
FROM (
  SELECT
    player_id,
    event_date,
    @cumsum := IF(player_id = @lastid, @cumsum + games_played, games_played) as games_played_so_far,
    @lastid := player_id
  FROM (
    SELECT
      player_id,
      event_date,
      SUM(games_played) as games_played
    FROM
      Activity
    GROUP BY
      1, 2
    ORDER BY
      1, 2 ASC
  ) a1
  JOIN (
    SELECT
      @lastid := -1,
      @cumsum := 0
  ) vars
) a2;
