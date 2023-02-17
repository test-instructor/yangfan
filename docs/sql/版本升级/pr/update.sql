
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
-- Records of api_py_pkg
-- ----------------------------
INSERT INTO hrp.api_py_pkg (created_at,updated_at,deleted_at,name,version,is_uninstall) VALUES
	 ('2023-02-07 20:05:30','2023-02-07 20:05:30',NULL,'funppy','0.5.0',0),
	 ('2023-02-07 20:06:30','2023-02-07 20:06:30',NULL,'grpcio','1.51.1',0),
	 ('2023-02-07 20:06:50','2023-02-07 20:06:50',NULL,'grpcio-tools','1.51.1',0),
	 ('2023-02-07 20:07:11','2023-02-07 20:07:11',NULL,'protobuf','4.21.12',0),
	 ('2023-02-07 20:07:21','2023-02-07 20:07:21',NULL,'pip','23.0',0);


-- ----------------------------
-- INSERT sys_base_menus
-- ----------------------------
INSERT INTO sys_base_menus (created_at,updated_at,deleted_at,menu_level,parent_id,`path`,name,hidden,component,sort,keep_alive,default_menu,title,icon,close_tab) VALUES
	 ('2023-01-29 16:21:37.647','2023-01-29 16:21:37.647',NULL,0,'33','HrpPyPkg','HrpPyPkg',0,'view/py_pkg/py_pkg.vue',0,0,0,'Python包管理','tools',0);

-- ----------------------------
-- INSERT sys_apis
-- ----------------------------
INSERT INTO sys_apis (created_at,updated_at,deleted_at,`path`,description,api_group,`method`) VALUES
	 ('2023-01-16 14:10:04.217','2023-01-16 14:10:04.217',NULL,'/case/:project/pyPkg/pyPkgList','python包','pypkg','GET'),
	 ('2023-01-16 15:18:44.127','2023-01-16 15:18:44.127',NULL,'/case/:project/pyPkg/installPyPkg','python包安装','pypkg','POST'),
	 ('2023-01-16 21:12:15.574','2023-01-16 21:12:15.574',NULL,'/case/:project/pyPkg/uninstallPyPkg','卸载python包','pypkg','POST'),
	 ('2023-01-16 21:45:20.167','2023-01-30 19:36:41.127',NULL,'/case/:project/pyPkg/updatePyPkg','升级python包','pypkg','POST'),
	 ('2023-02-05 21:29:05.056','2023-02-06 20:50:29.303',NULL,'/case/:project/pyPkg/getPkgVersionList','获取包版本','pypkg','POST');



-- ----------------------------
-- INSERT casbin_rule
-- ----------------------------
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES
	 ('p','888','/case/:project/pyPkg/pyPkgList','GET','','',''),
	 ('p','888','/case/:project/pyPkg/installPyPkg','POST','','',''),
	 ('p','888','/case/:project/pyPkg/uninstallPyPkg','POST','','',''),
	 ('p','888','/case/:project/pyPkg/updatePyPkg','POST','','',''),
	 ('p','888','/case/:project/pyPkg/getPkgVersionList','POST','','','');
------------------------------
INSERT INTO casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES
	 ('p','666','/case/:project/pyPkg/pyPkgList','GET','','',''),
	 ('p','666','/case/:project/pyPkg/installPyPkg','POST','','',''),
	 ('p','666','/case/:project/pyPkg/uninstallPyPkg','POST','','',''),
	 ('p','666','/case/:project/pyPkg/updatePyPkg','POST','','',''),
	 ('p','666','/case/:project/pyPkg/getPkgVersionList','POST','','','');