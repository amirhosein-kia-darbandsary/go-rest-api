CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- ID as a UUID
    slug VARCHAR(255) NOT NULL,                      -- Slug as a string (adjust length as needed)
    body TEXT NOT NULL,                              -- Body as a text field
    author VARCHAR(255) NOT NULL                     -- Author as a string (adjust length as needed)
);
