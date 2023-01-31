
DROP TABLE IF EXISTS `api_py_pkg`;
create table if not exists api_py_pkg
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)          null,
    updated_at   datetime(3)          null,
    deleted_at   datetime(3)          null,
    name         varchar(255)         null comment '包名称',
    version      varchar(255)         null comment '包版本',
    is_uninstall tinyint(1) default 1 null comment '是否可以卸载'
);

create index idx_api_py_pkg_deleted_at
    on api_py_pkg (deleted_at);

-- ----------------------------
-- Records of api_case_relationships
-- ----------------------------
INSERT INTO api_py_pkg (id, created_at, updated_at, deleted_at, name, version, is_uninstall) VALUES (1, '2023-01-30 13:34:08.000', '2023-01-30 13:34:10.000', null, 'requests', '2.28.2', 1);
INSERT INTO api_py_pkg (id, created_at, updated_at, deleted_at, name, version, is_uninstall) VALUES (2, '2023-01-31 12:50:17.297', '2023-01-31 12:50:17.297', null, 'Faker', '16.6.1', 1);