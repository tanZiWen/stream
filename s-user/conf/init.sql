DROP TABLE IF EXISTS app_user;
CREATE TABLE app_user (
  id        BIGINT PRIMARY KEY,
  username varchar(100),
  password  VARCHAR(100),
  nickname  VARCHAR(30)  NOT NULL,
  email     VARCHAR(50),
  mcc       SMALLINT,
  mobile    BIGINT,
  language  int DEFAULT (1),
  avatar_id varchar(33),
  avatar_url varchar(50),
  crt       TIMESTAMP WITH TIME ZONE,
  lut       TIMESTAMP WITH TIME ZONE,
  status    SMALLINT DEFAULT (1)
);

COMMENT ON TABLE app_user IS '用户表';
COMMENT ON COLUMN app_user.id IS '主键';
COMMENT ON COLUMN app_user.username IS '用户名（字母、数字、下划线、横线），可用于登录';
COMMENT ON COLUMN app_user.password IS '密码';
COMMENT ON COLUMN app_user.nickname IS '昵称';
COMMENT ON COLUMN app_user.email IS '邮箱';
COMMENT ON COLUMN app_user.mcc IS '手机号码的国家区号';
COMMENT ON COLUMN app_user.mobile IS '手机号码';
COMMENT ON COLUMN app_user.language IS '用户语言';
COMMENT ON COLUMN app_user.avatar_id IS '头像id';
COMMENT ON COLUMN app_user.avatar_url IS '原始头像id';
COMMENT ON COLUMN app_user.crt IS '创建时间';
COMMENT ON COLUMN app_user.lut IS '最后更新时间';
COMMENT ON COLUMN app_user.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';