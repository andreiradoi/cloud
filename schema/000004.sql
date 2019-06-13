DROP TYPE located_record CASCADE;
DROP TABLE fieldkit.dsl_summaries CASCADE;
DROP TABLE fieldkit.dsl_spatial_clusters CASCADE;
DROP TABLE fieldkit.dsl_temporal_geometries CASCADE;
DROP TABLE fieldkit.dsl_temporal_clusters CASCADE;

CREATE TYPE located_record AS (
  id INTEGER,
  timestamp timestamp,
  location geometry
);

CREATE TABLE fieldkit.dsl_summaries (
  device_id varchar(64) NOT NULL,
  updated_at timestamp NOT NULL,
  number_of_features integer NOT NULL,
  last_feature_id integer NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  envelope geometry NOT NULL,
  centroid geometry(POINT, 4326) NOT NULL,
  radius decimal NOT NULL,
  PRIMARY KEY (device_id)
);

CREATE TABLE fieldkit.dsl_spatial_clusters (
  device_id varchar(64) NOT NULL,
  cluster_id integer NOT NULL,
  updated_at timestamp NOT NULL,
  number_of_features integer NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  envelope geometry NOT NULL,
  circle geometry NOT NULL,
  centroid geometry(POINT, 4326) NOT NULL,
  radius decimal NOT NULL,
  PRIMARY KEY (device_id, cluster_id)
);

CREATE TABLE fieldkit.dsl_temporal_geometries (
  device_id varchar(64) NOT NULL,
  cluster_id integer NOT NULL,
  updated_at timestamp NOT NULL,
  geometry geometry(LINESTRING, 4326) NOT NULL,
  PRIMARY KEY (device_id, cluster_id)
);

CREATE TABLE fieldkit.dsl_temporal_clusters (
  device_id varchar(64) NOT NULL,
  cluster_id integer NOT NULL,
  updated_at timestamp NOT NULL,
  number_of_features integer NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  envelope geometry NOT NULL,
  centroid geometry(POINT, 4326) NOT NULL,
  radius decimal NOT NULL,
  PRIMARY KEY (device_id, cluster_id)
);


CREATE OR REPLACE VIEW fieldkit.device_stream_location_sanitized AS
WITH
  sanitized AS (
    SELECT
      *
    FROM fieldkit.device_stream_location AS dsl
    WHERE NOT (
            ST_XMax(location) >  180 OR ST_YMax(location) >  90 OR
            ST_XMin(location) < -180 OR ST_YMin(location) < -90
          )
  )
  SELECT s.device_id, s.timestamp, s.location
  FROM sanitized AS s;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_clustered_identical(desired_device_id varchar)
RETURNS TABLE (
	"device_id" varchar,
	"location" geometry,
	"size" BIGINT,
	"min_timestamp" timestamp,
	"max_timestamp" timestamp,
	"copy" BIGINT,
	"spatial_cluster_id" integer
) AS
'
BEGIN
RETURN QUERY
WITH
  with_identical_clustering AS (
    SELECT
      d.device_id AS device_id,
      d.location,
      COUNT(*) AS actual_size,
      MIN(d.timestamp) AS min_timestamp,
      MAX(d.timestamp) AS max_timestamp,
      LEAST(CAST(11 AS BIGINT), COUNT(*)) AS capped_size
    FROM fieldkit.device_stream_location_sanitized d
    WHERE d.device_id IN (desired_device_id) AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0
    GROUP BY d.device_id, d.location
  ),
  with_cluster_ids AS (
    SELECT
      s.device_id,
      s.location,
      s.actual_size,
      s.min_timestamp,
      s.max_timestamp,
      n AS n,
      ST_ClusterDBSCAN(ST_Transform(s.location, 2877), eps := 50, minPoints := 10) OVER () AS spatial_cluster_id
    FROM with_identical_clustering s, generate_series(0, s.capped_size - 1) AS x(n)
  )
  SELECT *
  FROM with_cluster_ids s
  WHERE s.spatial_cluster_id IS NOT NULL;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_tracks(desired_device_id varchar)
RETURNS TABLE (
	"device_id" varchar,
	"timestamp" timestamp,
	"min_timestamp" timestamp,
	"max_timestamp" timestamp,
	"time_difference" float,
	"actual_size" bigint,
	"temporal_cluster_id" BIGINT,
	"location" geometry
) AS
'
BEGIN
RETURN QUERY
WITH
source AS (
	SELECT
    to_timestamp(AVG(extract(epoch from d.timestamp)))::timestamp AS timestamp,
		MIN(d.timestamp) AS min_timestamp,
		MAX(d.timestamp) AS max_timestamp,
		MAX(d.timestamp) - MIN(d.timestamp) AS min_max_diff,
		d.location,
    (SELECT bool_or(ST_DWithin(d.location::geography, spatial.envelope::geography, 50)) FROM fieldkit.fk_dsl_spatial_clusters(desired_device_id) spatial WHERE spatial.device_id = desired_device_id) AS in_spatial,
    COUNT(d.location) AS actual_size
  FROM fieldkit.device_stream_location_sanitized d
  WHERE d.device_id IN (desired_device_id) AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0
  GROUP BY d.location
),
with_timestamp_differences AS (
	SELECT
		*,
  			                              LAG(s.timestamp) OVER (ORDER BY s.timestamp) AS previous_timestamp,
		EXTRACT(epoch FROM (s.timestamp - LAG(s.timestamp) OVER (ORDER BY s.timestamp))) AS time_difference
	FROM source s
  WHERE NOT s.in_spatial
	ORDER BY s.timestamp
),
with_temporal_clustering AS (
	SELECT
		*,
		CASE WHEN s.time_difference > 600
			OR s.time_difference IS NULL THEN true
			ELSE NULL
		END AS new_temporal_cluster
	FROM with_timestamp_differences s
),
with_assigned_temporal_clustering AS (
	SELECT
		*,
		COUNT(new_temporal_cluster/* OR spatial_cluster_change*/) OVER (
			ORDER BY s.timestamp
			ROWS UNBOUNDED PRECEDING
		) AS temporal_cluster_id
	FROM with_temporal_clustering s
)
SELECT
  desired_device_id AS device_id,
  s.timestamp,
  s.min_timestamp,
  s.max_timestamp,
  s.time_difference,
  s.actual_size,
  s.temporal_cluster_id,
  s.location
FROM with_assigned_temporal_clustering s;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_spatial_clusters(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_spatial_clusters
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_device_id::varchar(64) AS device_id,
  f.spatial_cluster_id,
  NOW()::timestamp,
  SUM(CASE copy WHEN 0 THEN f.size ELSE 0 END)::integer,
  MIN(min_timestamp),
  MAX(max_timestamp),
  ST_Envelope(ST_Collect(f.location)) AS envelope,
  ST_MinimumBoundingCircle(ST_Collect(f.location)) AS circle,
  ST_Centroid(ST_Collect(f.location))::geometry(POINT, 4326) AS centroid,
  SQRT(ST_Area(ST_MinimumBoundingCircle(ST_Collect(f.location))::geography) / pi())::numeric AS radius
FROM fieldkit.fk_dsl_clustered_identical(desired_device_id) AS f
GROUP BY spatial_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_temporal_clusters(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_temporal_clusters
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_device_id::varchar(64) AS device_id,
  f.temporal_cluster_id::integer,
  NOW()::timestamp,
  SUM(actual_size)::integer,
  MIN(min_timestamp),
  MAX(max_timestamp),
  ST_Envelope(ST_Collect(f.location)) AS envelope,
  ST_Centroid(ST_Collect(f.location))::geometry(POINT, 4326) AS centroid,
  SQRT(ST_Area(ST_MinimumBoundingCircle(ST_Collect(f.location))::geography) / pi())::numeric AS radius
FROM fieldkit.fk_dsl_tracks(desired_device_id) AS f
GROUP BY temporal_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_temporal_geometries(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_temporal_geometries
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_device_id::varchar(64) AS device_id,
  f.temporal_cluster_id::integer,
  NOW()::timestamp,
  ST_LineFromMultiPoint(ST_Collect(f.location))::geometry(LineString, 4326)
FROM fieldkit.fk_dsl_tracks(desired_device_id) AS f
GROUP BY temporal_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_dsl_summary(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_summaries
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_device_id::varchar(64) AS device_id,
  NOW()::timestamp AS updated_at,
	(SELECT COUNT(d.id)::integer AS number_of_features FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT d.Id::integer AS last_feature_id FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0 ORDER BY d.timestamp DESC LIMIT 1),
	(SELECT MIN(d.timestamp) AS start_time FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT MAX(d.timestamp) AS end_time FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT ST_Envelope(ST_Collect(d.location)) AS envelope FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT ST_Centroid(ST_Collect(d.location))::geometry(POINT, 4326) AS centroid FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT Sqrt(ST_Area(ST_MinimumBoundingCircle(ST_Collect(d.location))::geography))::numeric AS radius FROM fieldkit.device_stream_location AS d WHERE d.device_id = desired_device_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0);
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_dsl_summary(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_summaries
AS
'
BEGIN
  RETURN QUERY
  INSERT INTO fieldkit.dsl_summaries
  SELECT * FROM fieldkit.fk_dsl_summary(desired_device_id)
  ON CONFLICT (device_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    last_feature_id = excluded.last_feature_id,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    envelope = excluded.envelope,
    centroid = excluded.centroid,
    radius = excluded.radius
  RETURNING *;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_dsl_spatial_clusters(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_spatial_clusters
AS
'
BEGIN
  DELETE FROM fieldkit.dsl_spatial_clusters WHERE (device_id = desired_device_id);
  RETURN QUERY
  INSERT INTO fieldkit.dsl_spatial_clusters
  SELECT * FROM fieldkit.fk_dsl_spatial_clusters(desired_device_id)
  ON CONFLICT (device_id, cluster_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    envelope = excluded.envelope,
    circle = excluded.circle,
    centroid = excluded.centroid,
    radius = excluded.radius
  RETURNING *;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_dsl_temporal_clusters(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_temporal_clusters
AS
'
BEGIN
  DELETE FROM fieldkit.dsl_temporal_clusters c WHERE (c.device_id = desired_device_id);
  RETURN QUERY
  INSERT INTO fieldkit.dsl_temporal_clusters
  SELECT * FROM fieldkit.fk_dsl_temporal_clusters(desired_device_id)
  ON CONFLICT (device_id, cluster_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    envelope = excluded.envelope,
    centroid = excluded.centroid,
    radius = excluded.radius
  RETURNING *;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_dsl_temporal_geometries(desired_device_id varchar)
RETURNS SETOF fieldkit.dsl_temporal_geometries
AS
'
BEGIN
  DELETE FROM fieldkit.dsl_temporal_geometries WHERE (device_id = desired_device_id);
  RETURN QUERY
  INSERT INTO fieldkit.dsl_temporal_geometries
  SELECT * FROM fieldkit.fk_dsl_temporal_geometries(desired_device_id)
  ON CONFLICT (device_id, cluster_id) DO UPDATE SET
     updated_at = excluded.updated_at,
     geometry = excluded.geometry
  RETURNING *;
END
' LANGUAGE plpgsql;

  /*
    SELECT COUNT(*) FROM fieldkit.device_stream_location WHERE device_id = '0004a30b00232b9b';
    SELECT ST_AsText(location) FROM fieldkit.device_stream_location WHERE device_id = '0004a30b00232b9b' LIMIT 10;

    SELECT COUNT(*) FROM fieldkit.device_stream_location_sanitized WHERE device_id = '0004a30b00232b9b';
    SELECT ST_AsText(location) FROM fieldkit.device_stream_location_sanitized WHERE device_id = '0004a30b00232b9b' LIMIT 10;
   */

  /*
    SELECT * FROM fieldkit.fk_dsl_temporal_clusters('0004a30b00232b9b');

    SELECT * FROM fieldkit.fk_dsl_spatial_clusters('0004a30b00232b9b');

    SELECT ST_Transform(location, 2877) FROM fieldkit.device_stream_location;

    SELECT ST_Transform(location, 2950) FROM fieldkit.device_stream_location;


    SELECT ST_AsText(location)
    FROM fieldkit.device_stream_location WHERE
    ST_XMax(location) >  180 OR ST_YMax(location) >  90 OR
    ST_XMin(location) < -180 OR ST_YMin(location) < -90;
   */