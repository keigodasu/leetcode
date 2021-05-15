with scores as (
select
match_id,
first_player as player_id,
first_score as score
from
matches
union all
select
match_id,
second_player as player_id,
second_score as score
from
matches
)


select
group_id,
player_id
from
(
select
group_id,
player_id,
row_number() over(partition by group_id order by score desc, player_id asc) as rnk
from
(
select
player_id,
group_id,
sum(score) as score
from
scores
join
players
using(player_id)
group by player_id
) t
) tt
where tt.rnk = 1
