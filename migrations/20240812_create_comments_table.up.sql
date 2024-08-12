CREATE TABLE comments (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing integer ID
    slug VARCHAR(255) NOT NULL,  -- Slug as a string (adjust length as needed)
    body TEXT NOT NULL,  -- Body as a text field
    author VARCHAR(255) NOT NULL  -- Author as a string (adjust length as needed)
);
