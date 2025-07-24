-- InnoSync Database Initialization Script for PostgreSQL (Mobile App Version)

DROP TABLE IF EXISTS message CASCADE;
DROP TABLE IF EXISTS group_chat_member CASCADE;
DROP TABLE IF EXISTS group_chat CASCADE;
DROP TABLE IF EXISTS role_application CASCADE;
DROP TABLE IF EXISTS invitation CASCADE;
DROP TABLE IF EXISTS project_role CASCADE;
DROP TABLE IF EXISTS project CASCADE;
DROP TABLE IF EXISTS user_profile_technology CASCADE;
DROP TABLE IF EXISTS work_experience CASCADE;
DROP TABLE IF EXISTS user_profile CASCADE;
DROP TABLE IF EXISTS technology CASCADE;
DROP TABLE IF EXISTS refresh_token CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS application_status_enum CASCADE;
DROP TYPE IF EXISTS invitation_status_enum CASCADE;
DROP TYPE IF EXISTS expertise_level_enum CASCADE;
DROP TYPE IF EXISTS education_enum CASCADE;

CREATE TYPE education_enum AS ENUM ('NO_DEGREE', 'BACHELOR', 'MASTER', 'PHD');
CREATE TYPE expertise_level_enum AS ENUM ('ENTRY', 'JUNIOR', 'MID', 'SENIOR', 'RESEARCHER');
CREATE TYPE invitation_status_enum AS ENUM ('INVITED', 'ACCEPTED', 'DECLINED', 'REVOKED');
CREATE TYPE application_status_enum AS ENUM ('PENDING', 'UNDER_REVIEW', 'ACCEPTED', 'REJECTED', 'WITHDRAWN');


-- USER table (renamed to users to avoid PostgreSQL reserved word)
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       full_name VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       last_login TIMESTAMP
);

-- REFRESH_TOKEN table
CREATE TABLE refresh_token (
                               id BIGSERIAL PRIMARY KEY,
                               user_id BIGINT NOT NULL,
                               token VARCHAR(500) NOT NULL,
                               expiry_date TIMESTAMP NOT NULL,
                               device_id VARCHAR(255),  -- Added for mobile device tracking
                               FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- TECHNOLOGY table
CREATE TABLE technology (
                            id BIGSERIAL PRIMARY KEY,
                            name VARCHAR(255) UNIQUE NOT NULL,
                            category VARCHAR(100)  -- Added for better organization in mobile app
);

-- USER_PROFILE table
CREATE TABLE user_profile (
                              id BIGSERIAL PRIMARY KEY,
                              user_id BIGINT UNIQUE NOT NULL,
                              telegram VARCHAR(255),
                              github VARCHAR(255),
                              bio TEXT,
                              position VARCHAR(255),
                              education education_enum,
                              expertise TEXT,
                              expertise_level expertise_level_enum,
                              resume_url VARCHAR(500),  -- Changed to URL for mobile access
                              avatar_url VARCHAR(500),  -- Added for profile pictures
                              FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- WORK_EXPERIENCE table
CREATE TABLE work_experience (
                                 id BIGSERIAL PRIMARY KEY,
                                 user_profile_id BIGINT NOT NULL,
                                 start_date DATE NOT NULL,
                                 end_date DATE,  -- NULL means current position
                                 position VARCHAR(255) NOT NULL,
                                 company VARCHAR(255) NOT NULL,
                                 description TEXT,
                                 FOREIGN KEY (user_profile_id) REFERENCES user_profile(id) ON DELETE CASCADE
);

-- USER_PROFILE_TECHNOLOGY junction table
CREATE TABLE user_profile_technology (
                                         user_profile_id BIGINT NOT NULL,
                                         technology_id BIGINT NOT NULL,
                                         PRIMARY KEY (user_profile_id, technology_id),
                                         FOREIGN KEY (user_profile_id) REFERENCES user_profile(id) ON DELETE CASCADE,
                                         FOREIGN KEY (technology_id) REFERENCES technology(id) ON DELETE CASCADE
);

-- Project & Recruitment Tables

-- PROJECT table
CREATE TABLE project (
                         id BIGSERIAL PRIMARY KEY,
                         recruiter_id BIGINT NOT NULL,
                         title VARCHAR(255) NOT NULL,
                         description TEXT,
                         team_size INTEGER NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         is_active BOOLEAN DEFAULT TRUE,  -- Added for soft delete functionality
                         FOREIGN KEY (recruiter_id) REFERENCES users(id) ON DELETE CASCADE
);

-- PROJECT_ROLE table
CREATE TABLE project_role (
                              id BIGSERIAL PRIMARY KEY,
                              project_id BIGINT NOT NULL,
                              role_name VARCHAR(255) NOT NULL,
                              expertise_level expertise_level_enum,
                              technologies TEXT,  -- Simplified for mobile app (comma-separated values)
                              is_open BOOLEAN DEFAULT TRUE,  -- Added to mark if role is still available
                              FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);

-- INVITATION table
CREATE TABLE invitation (
                            id BIGSERIAL PRIMARY KEY,
                            project_role_id BIGINT NOT NULL,
                            user_id BIGINT NOT NULL,
                            invitation_status VARCHAR(255) DEFAULT 'INVITED',
                            sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            responded_at TIMESTAMP,
                            message TEXT,  -- Added for personal invitation message
                            FOREIGN KEY (project_role_id) REFERENCES project_role(id) ON DELETE CASCADE,
                            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ROLE_APPLICATION table
CREATE TABLE role_application (
                                  id BIGSERIAL PRIMARY KEY,
                                  user_id BIGINT NOT NULL,
                                  project_role_id BIGINT NOT NULL,
                                  status application_status_enum DEFAULT 'PENDING',
                                  applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  cover_letter TEXT,  -- Added for application details
                                  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                  FOREIGN KEY (project_role_id) REFERENCES project_role(id) ON DELETE CASCADE
);

-- Collaboration & Communication Tables

-- GROUP_CHAT table
CREATE TABLE group_chat (
                            id BIGSERIAL PRIMARY KEY,
                            project_id BIGINT NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            chat_name VARCHAR(255),  -- Added for custom chat names
                            is_active BOOLEAN DEFAULT TRUE,
                            FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);

-- GROUP_CHAT_MEMBER table
CREATE TABLE group_chat_member (
                                   group_chat_id BIGINT NOT NULL,
                                   user_id BIGINT NOT NULL,
                                   joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   PRIMARY KEY (group_chat_id, user_id),
                                   FOREIGN KEY (group_chat_id) REFERENCES group_chat(id) ON DELETE CASCADE,
                                   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- MESSAGE table
CREATE TABLE message (
                         id BIGSERIAL PRIMARY KEY,
                         sender_id BIGINT NOT NULL,
                         group_chat_id BIGINT NOT NULL,  -- Made mandatory for group-only chats
                         content TEXT NOT NULL,
                         sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         is_read BOOLEAN DEFAULT FALSE,  -- Added for read receipts
                         FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
                         FOREIGN KEY (group_chat_id) REFERENCES group_chat(id) ON DELETE CASCADE
);

-- Create indexes for performance optimization
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_refresh_token_user_id ON refresh_token(user_id);
CREATE INDEX idx_refresh_token_device ON refresh_token(device_id);
CREATE INDEX idx_user_profile_user_id ON user_profile(user_id);
CREATE INDEX idx_work_experience_user_profile_id ON work_experience(user_profile_id);
CREATE INDEX idx_project_recruiter_id ON project(recruiter_id);
CREATE INDEX idx_project_active ON project(is_active);
CREATE INDEX idx_project_role_project_id ON project_role(project_id);
CREATE INDEX idx_project_role_open ON project_role(is_open);
CREATE INDEX idx_invitation_project_role_id ON invitation(project_role_id);
CREATE INDEX idx_invitation_user_id ON invitation(user_id);
-- Переименовать индекс
DROP INDEX IF EXISTS idx_invitation_status;
CREATE INDEX idx_invitation_invitation_status ON invitation(invitation_status);
CREATE INDEX idx_role_application_user_id ON role_application(user_id);
CREATE INDEX idx_role_application_project_role_id ON role_application(project_role_id);
CREATE INDEX idx_role_application_status ON role_application(status);
CREATE INDEX idx_group_chat_project_id ON group_chat(project_id);
CREATE INDEX idx_message_sender_id ON message(sender_id);
CREATE INDEX idx_message_group_chat_id ON message(group_chat_id);
CREATE INDEX idx_message_sent_at ON message(sent_at);

-- Insert sample technologies with categories
INSERT INTO technology (name, category) VALUES
                                            ('Java', 'Backend'),
                                            ('Python', 'Backend'),
                                            ('JavaScript', 'Frontend'),
                                            ('TypeScript', 'Frontend'),
                                            ('React', 'Frontend'),
                                            ('Spring Boot', 'Backend'),
                                            ('Node.js', 'Backend'),
                                            ('PostgreSQL', 'Database'),
                                            ('MongoDB', 'Database'),
                                            ('Docker', 'DevOps'),
                                            ('Kubernetes', 'DevOps'),
                                            ('AWS', 'Cloud'),
                                            ('Flutter', 'Mobile'),
                                            ('Swift', 'Mobile'),
                                            ('Kotlin', 'Mobile'),
                                            ('TensorFlow', 'AI/ML'),
                                            ('Data Science', 'AI/ML');

-- Add trigger to update updated_at timestamp on role_application
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_role_application_updated_at
    BEFORE UPDATE ON role_application
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments to tables for documentation
COMMENT ON TABLE users IS 'Stores basic authentication data for all users';
COMMENT ON TABLE refresh_token IS 'Stores JWT refresh tokens linked to users and devices';
COMMENT ON TABLE user_profile IS 'Stores detailed user profile info beyond authentication';
COMMENT ON TABLE work_experience IS 'Records professional experiences for user profiles';
COMMENT ON TABLE technology IS 'Catalog of technologies used for tagging user skills and project needs';
COMMENT ON TABLE user_profile_technology IS 'Join table for many-to-many relationship between user profiles and technologies';
COMMENT ON TABLE project IS 'Stores projects created by recruiters';
COMMENT ON TABLE project_role IS 'Defines required roles and expectations per project';
COMMENT ON TABLE invitation IS 'Tracks project invitations sent from recruiters to recruitees';
COMMENT ON TABLE role_application IS 'Tracks user applications to project roles';
COMMENT ON TABLE group_chat IS 'Stores group chat info for project teams';
COMMENT ON TABLE group_chat_member IS 'Maps users to group chats';
COMMENT ON TABLE message IS 'Stores messages for group chats';

-- Add notification settings table for mobile app
CREATE TABLE notification_settings (
                                       user_id BIGINT PRIMARY KEY,
                                       new_message BOOLEAN DEFAULT TRUE,
                                       application_update BOOLEAN DEFAULT TRUE,
                                       invitation_received BOOLEAN DEFAULT TRUE,
                                       project_update BOOLEAN DEFAULT TRUE,
                                       FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);