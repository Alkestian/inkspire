IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Books')
BEGIN
  -- Create a table
  CREATE TABLE Books (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    title NVARCHAR(255) NOT NULL,
    author NVARCHAR(255) NOT NULL,
    series NVARCHAR(255),
    pages INT,
    status NVARCHAR(20) CHECK (status IN ('read', 'to-be-read')),
    updatedAt DATETIME2 DEFAULT SYSUTCDATETIME()
  );
END;
-- Insert sample data
INSERT INTO Books (title, author, series, pages, status)
VALUES ('The Lusty Argonian Maid', 'Sandon Branderson', 'Wheel of Time', 1690, 'read');
