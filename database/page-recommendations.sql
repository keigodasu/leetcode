select
distinct page_id recommended_page
from
likes
where 
user_id in 
(
select 
    user2_id
from
   Friendship
where
    user1_id = 1
union
select 
    user1_id
from
   Friendship
where
    user2_id = 1
)
and
page_id not in
(
select
    page_id
from
    likes
where
    user_id = 1
);

select
distinct page_id recommended_page
from
(
select
    case 
        when user1_id = 1 then user2_id
        when user2_id = 1 then user1_id
    end user_id
from friendship
) a
join
likes
using(user_id)
where
page_id not in (select page_id from likes where user_id = 1);

with friends_of_user1 as (
  select
    case
      when user1_id = 1 then user2_id
      when user2_id = 1 then user1_id
    end user_id
  from
    friendship
)
select
  distinct page_id recommended_page
from
  friends_of_user1
  join likes using(user_id)
where
  page_id not in (
    select
      page_id
    from
      likes
    where
      user_id = 1
  );

