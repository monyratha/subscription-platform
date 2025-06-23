-- Insert default admin user (email & hashed password are examples), password: admin123
INSERT INTO users (email, password, name, role, status)
VALUES ('admin@example.com', '$2a$12$.luPiIjPBOakhKse.5isaOKLTULi1qNj8v.Uwx.FlUGHEjrMSIJvO', 'Admin', 'admin', 1);

-- Seed Casbin policies
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'admin', '/admin/*', 'GET'),
       ('p', 'admin', '/admin/*', 'POST'),
       ('p', 'admin', '/admin/*', 'PUT'),
       ('p', 'admin', '/admin/*', 'DELETE'),
       ('p', 'user', '/user/profile', 'GET'),
       ('p', 'user', '/user/profile', 'PUT');

-- Example setting
INSERT INTO settings (setting_key, setting_value, description)
VALUES ('jwt_expiry_hours', '24', 'JWT token expiry duration in hours');
