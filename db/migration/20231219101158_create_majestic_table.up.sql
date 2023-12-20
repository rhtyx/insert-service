CREATE TABLE IF NOT EXISTS "majestics" (
  "id" SERIAL PRIMARY KEY,
  "global_rank" INT,
  "tld_rank" INT,
  "domain" VARCHAR,
  "tld" VARCHAR,
  "ref_subnets" INT,
  "ref_ips" INT,
  "idn_domain" VARCHAR,
  "idn_tld" VARCHAR,
  "prev_global_rank" INT,
  "prev_tld_rank" INT,
  "prev_ref_subnets" INT,
  "prev_ref_ips" INT
);