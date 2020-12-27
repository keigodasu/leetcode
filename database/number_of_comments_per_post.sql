with t as (
  select
    sub_id parent_id
  from
    submissions
  where
    parent_id is null
)
select
  t.parent_id post_id,
  count(distinct s.sub_id) number_of_comments
from
  t
  left outer join submissions s using(parent_id)
group by
  t.parent_id;
select
  s.sub_Id post_id,
  (
    select
      count(distinct s1.sub_id)
    from
      submissions s1
    where
      s1.parent_id = s.sub_id
  ) number_of_comments
from
  submissions s
where
  s.parent_id is null
group by
  s.sub_id
order by
  post_id;

