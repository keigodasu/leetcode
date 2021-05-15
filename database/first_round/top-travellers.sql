select
name,
sum(if(u.id = r.user_id, distance, 0)) travelled_distance
from
users u, rides r
group by u.id
order by travelled_distance desc, name asc;

select u.name, ifnull(sum(r.distance),0) travelled_distance
from users as u
left join rides as r
on u.id = r.user_id
group by u.name
order by travelled_distance desc, u.name;
