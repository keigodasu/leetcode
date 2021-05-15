with first_logined as (
  select
    player_id,
    min(event_date) first_logined_date
  from
    activity
  group by
    player_id
)
select
  round(
    count(*) / (
      select
        count(distinct player_id)
      from
        activity
    ),
    2
  ) as fraction
from
  first_logined f
  join activity a using(player_id)
where
  date_add(f.first_logined_date, interval 1 day) = a.event_date;
