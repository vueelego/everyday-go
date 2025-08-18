```sql  
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL, --加密存储
    avatar VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE projects (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    owner_id BIGINT NOT NULL,  -- 创建者
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);


CREATE TABLE folders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    owner_id BIGINT NOT NULL,
    project_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
);


CREATE TABLE requests (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    folder_id BIGINT,          -- 所属文件夹（可选）
    name VARCHAR(100) NOT NULL,
    method VARCHAR(16) NOT NULL,
    url TEXT NOT NULL,
    headers JSON,              -- 存储请求头（如 {"Content-Type": "application/json"}）
    body TEXT,                 -- 请求体
    query_params JSON,         -- 查询参数（如 {"page": 1}）
    auth_type ENUM('none', 'basic', 'bearer', 'oauth2') DEFAULT 'none',
    auth_config JSON,          -- 认证配置（如 token）
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (folder_id) REFERENCES folders(id)
);



CREATE TABLE environments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    project_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    variables JSON,  -- 存储变量（如 {"base_url": "https://api.example.com"}）
    is_global BOOLEAN DEFAULT FALSE,  -- 是否全局环境
    FOREIGN KEY (project_id) REFERENCES projects(id)
);


```


设计一些短链服务
shortener

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE entries (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255),
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_entries_short_code ON entries(short_code);
CREATE INDEX idx_entries_user_id ON entries(user_id);




CREATE TABLE clicks (
    id SERIAL PRIMARY KEY,
    url_id INTEGER NOT NULL REFERENCES entries(id) ON DELETE CASCADE,
    referrer TEXT,
    ip_address VARCHAR(45),
    user_agent TEXT,
    device_type VARCHAR(20), -- 'desktop', 'mobile', 'tablet', 'bot'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_clicks_url_id ON clicks(url_id);
CREATE INDEX idx_clicks_created_at ON clicks(created_at);


CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    owner_id BIGINT NOT NULL, 
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE url_tags (
    url_id INTEGER NOT NULL REFERENCES entries(id) ON DELETE CASCADE,
    tag_id INTEGER NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (url_id, tag_id)
);



CREATE TABLE api_quotas (
    user_id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    monthly_limit INTEGER DEFAULT 1000,
    used_this_month INTEGER DEFAULT 0,
    last_reset_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


```