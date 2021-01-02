select
a.visited_on,
sum(b.day_sum) amount,
round(avg(b.day_sum), 2) average_amount
from
(select visited_on, sum(amount) as day_sum from customer group by visited_on) a,
(select visited_on, sum(amount) as day_sum from customer group by visited_on) b
where datediff(a.visited_on, b.visited_on) between 0 and 6
group by a.visited_on
having count(b.visited_on) = 7
