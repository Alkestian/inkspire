-- Ensure Authors table exists
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Authors')
BEGIN
  CREATE TABLE Authors (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    name NVARCHAR(255) UNIQUE NOT NULL,
    birthDate DATE NULL,
    nationality NVARCHAR(100) NULL
  );
END;
GO

-- Ensure Series table exists
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Series')
BEGIN
  CREATE TABLE Series (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    name NVARCHAR(255) UNIQUE NOT NULL
  );
END;
GO

-- Ensure Books table exists
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Books')
BEGIN
  CREATE TABLE Books (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    title NVARCHAR(255) NOT NULL,
    authorId UNIQUEIDENTIFIER NOT NULL,
    seriesId UNIQUEIDENTIFIER NULL,
    pages INT NULL,
    status NVARCHAR(20) CHECK (status IN ('read', 'to-be-read')) NULL,
    updatedAt DATETIME2 DEFAULT SYSUTCDATETIME(),
    FOREIGN KEY (authorId) REFERENCES Authors(id) ON DELETE SET NULL,
    FOREIGN KEY (seriesId) REFERENCES Series(id) ON DELETE SET NULL
  );
END;
GO

-- Ensure AuthorSeries table exists (many-to-many)
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'AuthorSeries')
BEGIN
  CREATE TABLE AuthorSeries (
    authorId UNIQUEIDENTIFIER NOT NULL,
    seriesId UNIQUEIDENTIFIER NOT NULL,
    PRIMARY KEY (authorId, seriesId),
    FOREIGN KEY (authorId) REFERENCES Authors(id) ON DELETE CASCADE,
    FOREIGN KEY (seriesId) REFERENCES Series(id) ON DELETE CASCADE
  );
END;
GO

-- Ensure Genres table exists
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Genres')
BEGIN
  CREATE TABLE Genres (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    name NVARCHAR(100) UNIQUE NOT NULL
  );
END;
GO

-- Ensure BookGenres table exists (many-to-many)
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'BookGenres')
BEGIN
  CREATE TABLE BookGenres (
    bookId UNIQUEIDENTIFIER NOT NULL,
    genreId UNIQUEIDENTIFIER NOT NULL,
    PRIMARY KEY (bookId, genreId),
    FOREIGN KEY (bookId) REFERENCES Books(id) ON DELETE CASCADE,
    FOREIGN KEY (genreId) REFERENCES Genres(id) ON DELETE CASCADE
  );
END;
GO

-- Ensure Reviews table exists
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'Reviews')
BEGIN
  CREATE TABLE Reviews (
    id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
    bookId UNIQUEIDENTIFIER NOT NULL,
    rating INT CHECK (rating BETWEEN 1 AND 5) NULL,
    reviewText NVARCHAR(1000) NULL,
    createdAt DATETIME2 DEFAULT SYSUTCDATETIME(),
    FOREIGN KEY (bookId) REFERENCES Books(id) ON DELETE CASCADE
  );
END;
GO

-- Insert sample authors
INSERT INTO Authors (name) VALUES ('Sandon Branderson');
GO

-- Insert sample series
INSERT INTO Series (name) VALUES ('Wheel of Time');
GO

-- Insert sample book with authorId and seriesId
INSERT INTO Books (title, authorId, seriesId, pages, status)
VALUES (
    'The Lusty Argonian Maid',
    (SELECT TOP 1 id FROM Authors WHERE name = 'Sandon Branderson'),
    (SELECT TOP 1 id FROM Series WHERE name = 'Wheel of Time'),
    1690, 'read'
);
GO

-- Insert sample genre
INSERT INTO Genres (name) VALUES ('Fantasy');
GO

-- Insert into BookGenres using subqueries
INSERT INTO BookGenres (bookId, genreId)
VALUES (
    (SELECT TOP 1 id FROM Books WHERE title = 'The Lusty Argonian Maid'),
    (SELECT TOP 1 id FROM Genres WHERE name = 'Fantasy')
);
GO

-- Insert sample review
INSERT INTO Reviews (bookId, rating, reviewText)
VALUES (
    (SELECT TOP 1 id FROM Books WHERE title = 'The Lusty Argonian Maid'),
    5,
    'An absolute masterpiece!'
);
GO
