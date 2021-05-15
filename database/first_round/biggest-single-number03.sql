select
max(num) as num
from my_numbers m
where exists 
(select
num
from
    my_numbers mm
where  mm.num = m.num
group by num
having count(num) = 1);

