UPDATE
    salary
SET
    sex = case
        WHEN sex = 'm' THEN 'f'
        WHEN sex = 'f' THEN 'm'
END;
