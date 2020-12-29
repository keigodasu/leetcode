select
name,
sum(if(u.id = r.user_id, distance, 0)) travelled_distance
from
users u, rides r
group by u.id
order by travelled_distance desc, name asc;

