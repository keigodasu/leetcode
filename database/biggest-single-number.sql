select
  max(num) as num
from
  my_numbers
where
  exists (
    select
      num
    from
      my_numbers
    group by
      num
    having
      count(num) = 1
  );
