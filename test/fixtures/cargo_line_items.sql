insert into cargo_line_items(
    uuid4,
    cargo_item_id,
    shipping_marks,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
)
VALUES (
  UNHEX(REPLACE('aab30eb6-009b-411c-985c-527ce008b35a','-','')),
  8,
  'shipping marks',
  'active',
'auth0|66fd06d0bfea78a82bb42459',
'auth0|66fd06d0bfea78a82bb42459',
'2020-03-07 12:12:12.000',
'2020-04-07 12:12:12.000'
 );
