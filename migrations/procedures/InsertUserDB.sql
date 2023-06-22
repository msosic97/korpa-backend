CREATE DEFINER=`root`@`localhost` PROCEDURE `InsertUserDB`(
  IN username VARCHAR(255),
  IN password VARCHAR(255),
  IN email VARCHAR(255),
  IN firstname VARCHAR(255),
  IN lastname VARCHAR(255),
  IN phone VARCHAR(255),
  IN created_at BIGINT
)
BEGIN
  IF NOT EXISTS(SELECT ID FROM users WHERE email = email) THEN
    INSERT INTO users (username, password, email, firstname, lastname, phone, created_at)
    VALUES (username, password, email, firstname, lastname, phone, created_at);
  END IF;
END