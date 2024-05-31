CREATE TYPE one_two as ENUM ('1', '2');
CREATE TYPE sex as ENUM('M', 'F', 'X');

CREATE TABLE IF NOT EXISTS crime_codes (
    crime_committed_id numeric,
    crime_committed_description text
);

CREATE TABLE IF NOT EXISTS weapon_codes (
    weapon_used_id numeric,
    weapon_used_description text
);

CREATE TABLE IF NOT EXISTS crime_status_codes (
    crime_status_id text,
    crime_staus_description text
);

CREATE TABLE IF NOT EXISTS premis_codes (
    premis_id numeric,
    premis_description text
);

CREATE TABLE IF NOT EXISTS crimes (
    id uuid gen_random_uuid (),
    dr_no text,
    date_reported timestamp,
    date_occurred timestamp,
    time_occurred time,
    area text,
    area_name text,
    report_district_number text,
    part_1_2 one_two,
    crime_committed numeric,
    crime_committed_description text,
    mocodes text,
    victim_age numeric,
    victim_sex sex,
    victime_descent char,
    premis text,
    premis_description text,
    weapon_used_code numeric,
    weapon_used_description text,
    case_status_id text,
    case_status_description text,
    crime_committed_1 numeric,
    crime_committed_2 numeric,
    crime_committed_3 numeric,
    crime_committed_4 numeric,
    crime_location text,
    crime_cross_street text,
    LAT double precision,
    LON double precision
);


-- Original encodings
-- DR_NO,
-- Date Rptd ,
-- DATE OCC,
-- TIME OCC,
-- AREA,
-- AREA NAME,
-- Rpt Dist No,
-- Part 1-2,
-- Crm Cd,
-- Crm Cd Desc,
-- Mocodes,
-- Vict Age,
-- Vict Sex,
-- Vict Descent,
-- Premis Cd,
-- Premis Desc,
-- Weapon Used Cd,
-- Weapon Desc,
-- Status,
-- Status Desc,
-- Crm Cd 1,
-- Crm Cd 2,
-- Crm Cd 3,
-- Crm Cd 4,
-- LOCATION,
-- Cross Street,
-- LAT,
-- LON