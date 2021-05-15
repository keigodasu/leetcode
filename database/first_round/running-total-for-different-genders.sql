/* patern of using only self join */
select
s1.gender,
s1.day,
sum(s2.score_points) total
from
scores s1
join 
scores s2
on s1.day >= s2.day and s1.gender = s2.gender
group by s1.day,  s1.score_points
order by gender, day;

/* patern of using window fucntion */
select
gender, 
day,
sum(score_points) over(partition by gender order by day) total
from
scores;
