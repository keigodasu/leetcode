with staging as (
    select
            user_id,
                    spend_date,
                            GROUP_CONCAT(platform) as platform,
                                    sum(amount) as amount
                                        FROM Spending a
                                            group by 1,2
                                            ),
                                            results as (
                                            select 
                                                spend_date, 
                                                    case when platform != 'mobile' and platform != 'desktop' then 'both' else platform end as platform,
                                                        sum(amount) as total_amount,
                                                            count(distinct user_id) as total_users
                                                            from staging
                                                            group by 1,2
                                                            ),
                                                            table_rows as (
                                                                SELECT distinct spend_date, 'mobile' as platform FROM Spending
                                                                    UNION ALL
                                                                        SELECT distinct spend_date, 'desktop' as platform FROM Spending
                                                                            UNION ALL
                                                                                SELECT distinct spend_date, 'both' as platform FROM Spending
                                                                                )
                                                                                select 
                                                                                    a.spend_date, 
                                                                                        a.platform, 
                                                                                            COALESCE(b.total_amount,0) as total_amount, 
                                                                                                COALESCE(b.total_users,0) as total_users 
                                                                                                from table_rows a 
                                                                                                left join results b 
                                                                                                on a.spend_date = b.spend_date and a.platform = b.platform
