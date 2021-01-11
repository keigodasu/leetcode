(select
name results
from
movie_rating m
join
users u
using(user_id)
group by user_id
order by count(movie_id) desc, u.name asc
limit 1) 
union
(
select
title
from
movie_rating
join
movies
using(movie_id)
where left(created_at, 7) = '2020-02'
group by movie_id
order by  avg(rating) desc, title asc limit 1
);

