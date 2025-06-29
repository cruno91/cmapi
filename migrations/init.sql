-- 1. Entity Types (e.g., content, media, taxonomy)
CREATE TABLE entity_types (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);

-- 2. Bundles per entity type (e.g., article, video, tag)
CREATE TABLE bundles (
    id UUID PRIMARY KEY,
    entity_type_id UUID REFERENCES entity_types(id) ON DELETE CASCADE,
    name TEXT NOT NULL, -- machine name, e.g., "article"
    label TEXT NOT NULL,
    table_name TEXT NOT NULL UNIQUE -- e.g., "content__article"
);

-- 3. Field definitions per bundle
CREATE TABLE field_definitions (
    id UUID PRIMARY KEY,
    bundle_id UUID REFERENCES bundles(id) ON DELETE CASCADE,
    name TEXT NOT NULL,         -- field_machine_name, e.g., "title"
    type TEXT NOT NULL,         -- e.g., "text", "int", "reference"
    required BOOLEAN NOT NULL DEFAULT false,
    settings JSONB              -- optional field-type specific config
);
