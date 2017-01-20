DROP TABLE IF EXISTS app_user;
CREATE TABLE app_user
(
  id          BIGINT PRIMARY KEY NOT NULL,
  username    VARCHAR(20),
  password    VARCHAR(100),
  type        INTEGER  DEFAULT 1,
  nickname    VARCHAR(30)        NOT NULL,
  name        VARCHAR(20),
  email       VARCHAR(50),
  mcc         SMALLINT,
  mobile      BIGINT,
  birthday    DATE,
  id_type     INTEGER,
  id_number   VARCHAR(20),
  avatar_id   VARCHAR(33),
  avatar_url  VARCHAR(50),
  wx_openid   VARCHAR(30),
  paterner_id BIGINT,
  crt         TIMESTAMP WITH TIME ZONE,
  lut         TIMESTAMP WITH TIME ZONE,
  status      SMALLINT DEFAULT 1
);
COMMENT ON TABLE app_user IS '合作伙伴用户表';
COMMENT ON COLUMN app_user.id IS '主键';
COMMENT ON COLUMN app_user.username IS '用户名（字母、数字、下划线、横线），可用于登录';
COMMENT ON COLUMN app_user.password IS '密码';
COMMENT ON COLUMN app_user.type IS '类型, 1:员工, 2:渠道';
COMMENT ON COLUMN app_user.nickname IS '昵称';
COMMENT ON COLUMN app_user.name IS '姓名';
COMMENT ON COLUMN app_user.email IS '邮箱';
COMMENT ON COLUMN app_user.mcc IS '手机号码的国家区号';
COMMENT ON COLUMN app_user.mobile IS '手机号码';
COMMENT ON COLUMN app_user.birthday IS '出生日期';
COMMENT ON COLUMN app_user.id_type IS '证件类型';
COMMENT ON COLUMN app_user.id_number IS '证件号码';
COMMENT ON COLUMN app_user.avatar_id IS '头像id';
COMMENT ON COLUMN app_user.avatar_url IS '原始头像id';
COMMENT ON COLUMN app_user.wx_openid IS '微信openid';
COMMENT ON COLUMN app_user.paterner_id IS '合作伙伴ID';
COMMENT ON COLUMN app_user.crt IS '创建时间';
COMMENT ON COLUMN app_user.lut IS '最后更新时间';
COMMENT ON COLUMN app_user.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS customer;
CREATE TABLE customer
(
  id           BIGINT PRIMARY KEY NOT NULL,
  name         VARCHAR(40),
  first_name   VARCHAR(20),
  last_name    VARCHAR(20),
  gender       INTEGER,
  mobile       VARCHAR(20),
  email        VARCHAR(50),
  birthday     VARCHAR(15),
  id_type      INTEGER,
  id_number    VARCHAR(20),
  wx_openid    VARCHAR(30),
  member_grade INTEGER,
  owner_id     BIGINT,
  last_feed_id BIGINT,
  crt          TIMESTAMP WITH TIME ZONE,
  lut          TIMESTAMP WITH TIME ZONE,
  status       SMALLINT DEFAULT 1
);
CREATE UNIQUE INDEX customer_mobile_key ON customer (mobile);

COMMENT ON TABLE customer IS '客户表';
COMMENT ON COLUMN customer.id IS '主键';
COMMENT ON COLUMN customer.name IS '客户姓名';
COMMENT ON COLUMN customer.first_name IS '名';
COMMENT ON COLUMN customer.last_name IS '姓';
COMMENT ON COLUMN customer.gender IS '性别';
COMMENT ON COLUMN customer.mobile IS '客户手机号码';
COMMENT ON COLUMN customer.email IS '客户email';
COMMENT ON COLUMN customer.birthday IS '出生日期';
COMMENT ON COLUMN customer.id_type IS '证件类型';
COMMENT ON COLUMN customer.id_number IS '证件号码';
COMMENT ON COLUMN customer.wx_openid IS '微信openid';
COMMENT ON COLUMN customer.member_grade IS '会员等级';
COMMENT ON COLUMN customer.owner_id IS '客户所有者id';
COMMENT ON COLUMN customer.last_feed_id IS '最新动态id';
COMMENT ON COLUMN customer.crt IS '创建时间';
COMMENT ON COLUMN customer.lut IS '最后更新时间';
COMMENT ON COLUMN customer.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS feed;
CREATE TABLE feed
(
  id          BIGINT PRIMARY KEY NOT NULL,
  customer_id BIGINT,
  creator_id  BIGINT,
  title       VARCHAR(50),
  content     TEXT,
  img_ids     BIGINT [],
  crt         TIMESTAMP WITH TIME ZONE,
  lut         TIMESTAMP WITH TIME ZONE,
  status      SMALLINT DEFAULT 1
);
COMMENT ON TABLE feed IS '表';
COMMENT ON COLUMN feed.id IS '主键';
COMMENT ON COLUMN feed.customer_id IS '客户id';
COMMENT ON COLUMN feed.creator_id IS '创建人id';
COMMENT ON COLUMN feed.title IS '标题';
COMMENT ON COLUMN feed.content IS '内容';
COMMENT ON COLUMN feed.img_ids IS '图片';
COMMENT ON COLUMN feed.crt IS '创建时间';
COMMENT ON COLUMN feed.lut IS '最后更新时间';
COMMENT ON COLUMN feed.status IS '状态，0:删除, 1:正常';

DROP TABLE IF EXISTS partner;
CREATE TABLE partner
(
  id             BIGINT PRIMARY KEY NOT NULL,
  name           VARCHAR(40),
  relation       INTEGER [],
  brand_id       BIGINT,
  website        VARCHAR(40),
  logo_url       VARCHAR(50),
  category       INTEGER [],
  brief          VARCHAR(200),
  description_id BIGINT,
  crt            TIMESTAMP WITH TIME ZONE,
  lut            TIMESTAMP WITH TIME ZONE,
  status         SMALLINT DEFAULT 1
);
COMMENT ON TABLE partner IS '合作伙伴表';
COMMENT ON COLUMN partner.id IS '主键';
COMMENT ON COLUMN partner.name IS '名称';
COMMENT ON COLUMN partner.relation IS '关系类型';
COMMENT ON COLUMN partner.brand_id IS '品牌id';
COMMENT ON COLUMN partner.website IS '官网地址';
COMMENT ON COLUMN partner.logo_url IS 'logo地址';
COMMENT ON COLUMN partner.category IS '所属分类';
COMMENT ON COLUMN partner.brief IS '简介';
COMMENT ON COLUMN partner.description_id IS '详情文章id';
COMMENT ON COLUMN partner.crt IS '创建时间';
COMMENT ON COLUMN partner.lut IS '最后更新时间';
COMMENT ON COLUMN partner.status IS '状态，0:删除, 1:正常, 2:冻结 默认为 1 ';

DROP TABLE IF EXISTS brand;
CREATE TABLE brand
(
  id             BIGINT PRIMARY KEY NOT NULL,
  name           VARCHAR(40),
  website        VARCHAR(40),
  logo_url       VARCHAR(50),
  category_id    INTEGER [],
  brief          VARCHAR(200),
  description_id BIGINT,
  authorized     BOOLEAN,
  crt            TIMESTAMP WITH TIME ZONE,
  lut            TIMESTAMP WITH TIME ZONE,
  status         SMALLINT DEFAULT 1
);
COMMENT ON TABLE brand IS '品牌表';
COMMENT ON COLUMN brand.id IS '主键';
COMMENT ON COLUMN brand.name IS '名称';
COMMENT ON COLUMN brand.website IS '官网地址';
COMMENT ON COLUMN brand.logo_url IS 'logo地址';
COMMENT ON COLUMN brand.category_id IS '所属分类';
COMMENT ON COLUMN brand.brief IS '简介';
COMMENT ON COLUMN brand.description_id IS '详情文章id';
COMMENT ON COLUMN brand.authorized IS '是否已取得商标认证';
COMMENT ON COLUMN brand.crt IS '创建时间';
COMMENT ON COLUMN brand.lut IS '最后更新时间';
COMMENT ON COLUMN brand.status IS '状态，0:删除, 1:正常, 2:冻结 默认为 1 ';

DROP TABLE IF EXISTS partner_brand;
CREATE TABLE partner_brand
(
  id              BIGINT PRIMARY KEY NOT NULL,
  partner_id      BIGINT             NOT NULL,
  brand_id        BIGINT             NOT NULL,
  relation        INTEGER,
  relation_detail JSONB,
  crt             TIMESTAMP WITH TIME ZONE,
  lut             TIMESTAMP WITH TIME ZONE,
  status          SMALLINT DEFAULT 1
);
COMMENT ON TABLE partner_brand IS '合作伙伴品牌关系表';
COMMENT ON COLUMN partner_brand.id IS '主键';
COMMENT ON COLUMN partner_brand.partner_id IS '合作伙伴id';
COMMENT ON COLUMN partner_brand.brand_id IS '品牌主键';
COMMENT ON COLUMN partner_brand.relation IS '关系类型, 1:所有者, 2:分支机构, 3:分销商';
COMMENT ON COLUMN partner_brand.relation_detail IS '关系详述';
COMMENT ON COLUMN partner_brand.crt IS '创建时间';
COMMENT ON COLUMN partner_brand.lut IS '最后更新时间';
COMMENT ON COLUMN partner_brand.status IS '状态，0:删除, 1:正常';

DROP TABLE IF EXISTS spu;
CREATE TABLE spu
(
  id          BIGINT PRIMARY KEY NOT NULL,
  name        VARCHAR(100),
  owner_id    BIGINT,
  cate_ids    BIGINT [],
  description TEXT,
  crt         TIMESTAMP WITH TIME ZONE,
  lut         TIMESTAMP WITH TIME ZONE,
  status      SMALLINT DEFAULT 1
);
COMMENT ON TABLE spu IS '';
COMMENT ON COLUMN spu.id IS '主键';
COMMENT ON COLUMN spu.name IS '名称';
COMMENT ON COLUMN spu.owner_id IS '所有者id';
COMMENT ON COLUMN spu.cate_ids IS '所属分类id';
COMMENT ON COLUMN spu.description IS '产品详情';
COMMENT ON COLUMN spu.crt IS '创建时间';
COMMENT ON COLUMN spu.lut IS '最后更新时间';
COMMENT ON COLUMN spu.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS item;
CREATE TABLE item
(
  id             BIGINT PRIMARY KEY NOT NULL,
  name           VARCHAR(100),
  spu_id         BIGINT,
  partner_id     BIGINT,
  cate_ids       BIGINT [],
  description    TEXT,
  img_ids        BIGINT [],
  leading_sku_id BIGINT,
  sale_status    SMALLINT DEFAULT 0,
  publish_time   TIMESTAMP WITH TIME ZONE,
  crt            TIMESTAMP WITH TIME ZONE,
  lut            TIMESTAMP WITH TIME ZONE,
  status         SMALLINT DEFAULT 1
);
COMMENT ON TABLE item IS '商品表';
COMMENT ON COLUMN item.id IS '主键';
COMMENT ON COLUMN item.name IS '产品名称';
COMMENT ON COLUMN item.spu_id IS 'spu主键';
COMMENT ON COLUMN item.partner_id IS '合作伙伴id';
COMMENT ON COLUMN item.cate_ids IS '所属分类id';
COMMENT ON COLUMN item.img_ids IS '图片id';
COMMENT ON COLUMN item.description IS '产品详情';
COMMENT ON COLUMN item.leading_sku_id IS '主打sku id';
COMMENT ON COLUMN item.sale_status IS '销售状态, 0: 未上架, 1: 上架, -1: 下架';
COMMENT ON COLUMN item.publish_time IS '发布(上架)时间';
COMMENT ON COLUMN item.crt IS '创建时间';
COMMENT ON COLUMN item.lut IS '最后更新时间';
COMMENT ON COLUMN item.status IS '状态，0:删除, 1:正常, 默认为 1 ';

DROP TABLE IF EXISTS sku;
CREATE TABLE sku
(
  id           BIGINT PRIMARY KEY NOT NULL,
  spu_id       BIGINT,
  item_id      BIGINT,
  partner_id   BIGINT,
  name         VARCHAR(100),
  description  TEXT,
  sale_status  SMALLINT DEFAULT 0,
  stock_number INTEGER  DEFAULT 0,
  pricing_type SMALLINT DEFAULT 0,
  market_price REAL,
  sale_price   REAL,
  cover_id     BIGINT,
  cover_url    VARCHAR(100),
  publish_time TIMESTAMP WITH TIME ZONE,
  crt          TIMESTAMP WITH TIME ZONE,
  lut          TIMESTAMP WITH TIME ZONE,
  status       SMALLINT DEFAULT 1
);
COMMENT ON TABLE sku IS '';
COMMENT ON COLUMN sku.id IS '主键';
COMMENT ON COLUMN sku.spu_id IS 'spu主键';
COMMENT ON COLUMN sku.item_id IS '商品主键';
COMMENT ON COLUMN sku.partner_id IS '合作伙伴id';
COMMENT ON COLUMN sku.name IS '产品名称';
COMMENT ON COLUMN sku.description IS '产品详情';
COMMENT ON COLUMN sku.sale_status IS '销售状态, 0: 未上架, 1: 上架, -1: 下架';
COMMENT ON COLUMN sku.stock_number IS '库存, -1: 无限供应, >0: 实际库存';
COMMENT ON COLUMN sku.pricing_type IS '定价方式, 0: 一口价, 1: 定金, 2: 下单后报价';
COMMENT ON COLUMN sku.market_price IS '市场价';
COMMENT ON COLUMN sku.sale_price IS '现价';
COMMENT ON COLUMN sku.cover_id IS '封面id';
COMMENT ON COLUMN sku.cover_url IS '封面url';
COMMENT ON COLUMN sku.publish_time IS '发布(上架)时间';
COMMENT ON COLUMN sku.crt IS '创建时间';
COMMENT ON COLUMN sku.lut IS '最后更新时间';
COMMENT ON COLUMN sku.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS att_cate;
CREATE TABLE att_cate
(
  id          BIGINT PRIMARY KEY NOT NULL,
  name        VARCHAR(20),
  parent_id   BIGINT,
  path        BIGINT [],
  owner_id    BIGINT,
  description TEXT,
  crt         TIMESTAMP WITH TIME ZONE,
  lut         TIMESTAMP WITH TIME ZONE,
  status      SMALLINT DEFAULT 1
);
COMMENT ON TABLE att_cate IS '属性分类表';
COMMENT ON COLUMN att_cate.id IS '主键';
COMMENT ON COLUMN att_cate.name IS '名称';
COMMENT ON COLUMN att_cate.parent_id IS '名称';
COMMENT ON COLUMN att_cate.path IS '父子关系路径';
COMMENT ON COLUMN att_cate.owner_id IS '所有者id';
COMMENT ON COLUMN att_cate.description IS '描述';
COMMENT ON COLUMN att_cate.crt IS '创建时间';
COMMENT ON COLUMN att_cate.lut IS '最后更新时间';
COMMENT ON COLUMN att_cate.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS attribute;
CREATE TABLE attribute
(
  id          BIGINT PRIMARY KEY NOT NULL,
  name        VARCHAR(20),
  value_unit  SMALLINT,
  owner_id    BIGINT,
  description TEXT,
  crt         TIMESTAMP WITH TIME ZONE,
  lut         TIMESTAMP WITH TIME ZONE,
  status      SMALLINT DEFAULT 1
);
COMMENT ON TABLE attribute IS '属性表';
COMMENT ON COLUMN attribute.id IS '主键';
COMMENT ON COLUMN attribute.name IS '名称';
COMMENT ON COLUMN attribute.owner_id IS '所有者id(平台或partner)';
COMMENT ON COLUMN attribute.description IS '属性描述';
COMMENT ON COLUMN attribute.crt IS '创建时间';
COMMENT ON COLUMN attribute.lut IS '最后更新时间';
COMMENT ON COLUMN attribute.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS model_att;
CREATE TABLE model_att
(
  id               BIGINT PRIMARY KEY NOT NULL,
  model_type       SMALLINT,
  model_id         BIGINT,
  att_id           BIGINT,
  valued           BOOL,
  sale_att         BOOL,
  required         BOOL,
  value_method     SMALLINT,
  value_data_type  SMALLINT,
  att_text_value   VARCHAR(200),
  att_int_value    INTEGER,
  att_option_value INTEGER,
  priority         INTEGER,
  alias            VARCHAR(20),
  description      TEXT,
  crt              TIMESTAMP WITH TIME ZONE,
  lut              TIMESTAMP WITH TIME ZONE,
  status           SMALLINT DEFAULT 1
);
COMMENT ON TABLE model_att IS '模型(spu,item,sku)属性表';
COMMENT ON COLUMN model_att.id IS '主键';
COMMENT ON COLUMN model_att.model_type IS '模型类型, 1: spu, 2: item, 3: sku';
COMMENT ON COLUMN model_att.model_id IS '模型id';
COMMENT ON COLUMN model_att.att_id IS '属性id';
COMMENT ON COLUMN model_att.valued IS '值是否已确定';
COMMENT ON COLUMN model_att.sale_att IS '是否为销售属性（下单时或下单后确定值）';
COMMENT ON COLUMN model_att.required IS '是否必输项';
COMMENT ON COLUMN model_att.value_method IS '输入方式类型, 0: 单选, 1: 文本框, 2: 多行文本';
COMMENT ON COLUMN model_att.value_data_type IS '值类型, 0: 字典值(整型), 1: 整型, 2: 浮点型, 3, 字符型';
COMMENT ON COLUMN model_att.att_text_value IS '属性字符值';
COMMENT ON COLUMN model_att.att_int_value IS '属性字符值';
COMMENT ON COLUMN model_att.att_option_value IS '属性字符值';
COMMENT ON COLUMN model_att.priority IS '显示优先级,优先级越大显示越靠前';
COMMENT ON COLUMN model_att.alias IS '别名/key';
COMMENT ON COLUMN model_att.description IS '属性描述';
COMMENT ON COLUMN model_att.crt IS '创建时间';
COMMENT ON COLUMN model_att.lut IS '最后更新时间';
COMMENT ON COLUMN model_att.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS model_att_option;
CREATE TABLE model_att_option
(
  id         BIGINT PRIMARY KEY NOT NULL,
  model_type SMALLINT,
  model_id   BIGINT,
  att_id     BIGINT,
  value      INTEGER,
  title      VARCHAR(50),
  image_url  VARCHAR(100),
  image_id   BIGINT,
  crt        TIMESTAMP WITH TIME ZONE,
  lut        TIMESTAMP WITH TIME ZONE,
  status     SMALLINT DEFAULT 1
);
COMMENT ON TABLE model_att_option IS '属性选项表';
COMMENT ON COLUMN model_att_option.id IS '主键';
COMMENT ON COLUMN model_att_option.att_id IS '属性主键';
COMMENT ON COLUMN model_att_option.value IS '选项字典值';
COMMENT ON COLUMN model_att_option.title IS '标题';
COMMENT ON COLUMN model_att_option.image_url IS '图片链接';
COMMENT ON COLUMN model_att_option.crt IS '创建时间';
COMMENT ON COLUMN model_att_option.lut IS '最后更新时间';
COMMENT ON COLUMN model_att_option.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS orders;
CREATE TABLE orders
(
  id              BIGINT PRIMARY KEY NOT NULL,
  customer_id     BIGINT             NOT NULL,
  creator_id      BIGINT,
  total_number    INTEGER,
  total_price     REAL,
  paid_amount     REAL,
  ship_status     SMALLINT,
  done            BOOL,
  has_sub_order   BOOL,
  is_sub_order    BOOL,
  parent_order_id BIGINT,
  partner_id      BIGINT,
  crt             TIMESTAMP WITH TIME ZONE,
  lut             TIMESTAMP WITH TIME ZONE,
  status          SMALLINT DEFAULT 1
);
COMMENT ON TABLE orders IS '订单表';
COMMENT ON COLUMN orders.id IS '主键';
COMMENT ON COLUMN orders.customer_id IS '客户id';
COMMENT ON COLUMN orders.creator_id IS '创建人id,客户本人/partner user';
COMMENT ON COLUMN orders.total_number IS '订单包含的产品总个数';
COMMENT ON COLUMN orders.total_price IS '订单总金额';
COMMENT ON COLUMN orders.paid_amount IS '已支付金额';
COMMENT ON COLUMN orders.ship_status IS '发货状态';
COMMENT ON COLUMN orders.done IS '是否已完成';
COMMENT ON COLUMN orders.has_sub_order IS '是否有子订单';
COMMENT ON COLUMN orders.parent_order_id IS '父订单id';
COMMENT ON COLUMN orders.partner_id IS '合作伙伴id';
COMMENT ON COLUMN orders.crt IS '创建时间';
COMMENT ON COLUMN orders.lut IS '最后更新时间';
COMMENT ON COLUMN orders.status IS '状态，0:删除, 1:正常, 2:冻结, 默认为 1 ';

DROP TABLE IF EXISTS order_item;
CREATE TABLE order_item
(
  id              BIGINT PRIMARY KEY NOT NULL,
  parent_order_id BIGINT,
  order_id        BIGINT,
  item_id         BIGINT,
  sku_id          BIGINT,
  number          INTEGER,
  sale_price      REAL,
  price           REAL,
  comment         VARCHAR(100),
  crt             TIMESTAMP WITH TIME ZONE,
  lut             TIMESTAMP WITH TIME ZONE,
  status          SMALLINT DEFAULT 1
);
COMMENT ON TABLE order_item IS '表';
COMMENT ON COLUMN order_item.id IS '主键';
COMMENT ON COLUMN order_item.parent_order_id IS '父订单id';
COMMENT ON COLUMN order_item.order_id IS '订单id';
COMMENT ON COLUMN order_item.item_id IS '商品id';
COMMENT ON COLUMN order_item.sku_id IS 'sku id';
COMMENT ON COLUMN order_item.number IS '数量';
COMMENT ON COLUMN order_item.sale_price IS '标价';
COMMENT ON COLUMN order_item.price IS '实际成交价';
COMMENT ON COLUMN order_item.comment IS '备注';
COMMENT ON COLUMN order_item.crt IS '创建时间';
COMMENT ON COLUMN order_item.lut IS '最后更新时间';
COMMENT ON COLUMN order_item.status IS '状态，0:删除, 1:正常';

DROP TABLE IF EXISTS order_item_att_value;
CREATE TABLE order_item_att_value
(
  id               BIGINT PRIMARY KEY NOT NULL,
  order_id         BIGINT,
  item_id          BIGINT,
  att_id           BIGINT,
  att_text_value   VARCHAR(50),
  att_int_value    INTEGER,
  att_float_value  REAL,
  att_option_value INTEGER,
  crt              TIMESTAMP WITH TIME ZONE,
  lut              TIMESTAMP WITH TIME ZONE,
  status           SMALLINT DEFAULT 1
);
COMMENT ON TABLE order_item_att_value IS '表';
COMMENT ON COLUMN order_item_att_value.id IS '主键';
COMMENT ON COLUMN order_item_att_value.order_id IS '订单id';
COMMENT ON COLUMN order_item_att_value.item_id IS '商品id';
COMMENT ON COLUMN order_item_att_value.att_id IS '属性id';
COMMENT ON COLUMN order_item_att_value.att_text_value IS '属性字符值';
COMMENT ON COLUMN order_item_att_value.att_int_value IS '属性整型值';
COMMENT ON COLUMN order_item_att_value.att_float_value IS '属性浮点型值';
COMMENT ON COLUMN order_item_att_value.att_option_value IS '属性字典值';
COMMENT ON COLUMN order_item_att_value.crt IS '创建时间';
COMMENT ON COLUMN order_item_att_value.lut IS '最后更新时间';
COMMENT ON COLUMN order_item_att_value.status IS '状态，0:删除, 1:正常';