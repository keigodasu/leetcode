select
  max(num) num
from
  my_numbers n
where
  not exists (
    select
      inn.num
    from
      my_numbers inn
    where
      inn.num = n.num
    group by
      num
    having
      count(1) > 1
  )

