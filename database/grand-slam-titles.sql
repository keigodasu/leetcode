select
t.player_id,
player_name,
count(*) as grand_slams_count
from
(
select
    wimbledon as player_id
from
    Championships 
union all
select
   fr_open  as player_id
from
    Championships 
union all
select
   us_open  as player_id
from
    Championships 
union all
select
   au_open  as player_id
from
    Championships 
) t
join
players
using(player_id)
group by t.player_id
