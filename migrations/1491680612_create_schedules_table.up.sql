CREATE TABLE schedules (
       id serial,
       period integer,
       submissions_num integer,
       subreddit varchar(20),
       slack_chat varchar(128),
       last_posted_at timestamp
);

CREATE INDEX last_posted_index ON schedules(last_posted_at);
