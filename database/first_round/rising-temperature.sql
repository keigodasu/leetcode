select
w2.id
from
Weather w1
join
Weather w2
where
w1.recordDate = DATE_SUB(w2.recordDate, INTERVAL 1 DAY)
and
w1.Temperature < w2.Temperature
