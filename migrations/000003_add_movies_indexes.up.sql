CREATE INDEX IF NOT EXISTS cameras_title_idx ON cameras USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS cameras_manufacturer_idx ON cameras USING GIN (to_tsvector('simple', manufacturer));
CREATE INDEX IF NOT EXISTS cameras_models_idx ON cameras USING GIN (to_tsvector('simple', model));