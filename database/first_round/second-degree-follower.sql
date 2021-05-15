select
followee as follower,
count(distinct follower) as num
from
follow
where followee in
(
select
distinct f1.follower
from
follow f1
left join
follow f2
on f1.follower = f2.followee
where f2.followee is not null
)
group by followee
order by follower asc
