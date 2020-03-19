
-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (1, 0, 'TOP', '', NULL, 1, 1, 1, 'TOP', '', 'none', '2020-03-10 14:47:13', '2020-03-10 14:47:13', NULL);
INSERT INTO `sys_menu` VALUES (2, 1, '系统管理', '', NULL, 1, 1, 1, 'Sys', 'lock', 'none', '2020-03-10 14:47:14', '2020-03-10 14:47:14', NULL);
INSERT INTO `sys_menu` VALUES (3, 2, '图标管理', '/icon', NULL, 10, 2, 1, 'Icon', 'icon', 'none', '2020-03-10 14:47:15', '2020-03-10 14:47:15', NULL);
INSERT INTO `sys_menu` VALUES (4, 2, '菜单管理', '/menu', NULL, 20, 2, 1, 'Menu', 'documentation', 'none', '2020-03-10 14:47:16', '2020-03-10 14:47:16', NULL);
INSERT INTO `sys_menu` VALUES (5, 4, '新增', '/menu/create', NULL, 1, 3, 1, 'MenuAdd', '', 'add', '2020-03-10 14:47:17', '2020-03-10 14:47:17', NULL);
INSERT INTO `sys_menu` VALUES (6, 4, '删除', '/menu/delete', NULL, 2, 3, 1, 'MenuDel', '', 'del', '2020-03-10 14:47:18', '2020-03-10 14:47:18', NULL);
INSERT INTO `sys_menu` VALUES (7, 4, '查看', '/menu/detail', NULL, 3, 3, 1, 'MenuView', '', 'view', '2020-03-10 14:47:19', '2020-03-10 14:47:19', NULL);
INSERT INTO `sys_menu` VALUES (8, 4, '编辑', '/menu/update', NULL, 4, 3, 1, 'MenuUpdate', '', 'update', '2020-03-10 14:47:20', '2020-03-10 14:47:20', NULL);
INSERT INTO `sys_menu` VALUES (9, 4, '分页api', '/menu/list', NULL, 5, 3, 1, 'MenuList', '', 'list', '2020-03-10 14:47:21', '2020-03-10 14:47:21', NULL);
INSERT INTO `sys_menu` VALUES (10, 2, '角色管理', '/role', NULL, 30, 2, 1, 'Role', 'tree', 'none', '2020-03-10 14:47:22', '2020-03-10 14:47:22', NULL);
INSERT INTO `sys_menu` VALUES (11, 10, '新增', '/role/create', NULL, 1, 3, 1, 'RoleAdd', '', 'add', '2020-03-10 14:47:23', '2020-03-10 14:47:23', NULL);
INSERT INTO `sys_menu` VALUES (12, 10, '删除', '/role/delete', NULL, 2, 3, 1, 'RoleDel', '', 'del', '2020-03-10 14:47:24', '2020-03-10 14:47:24', NULL);
INSERT INTO `sys_menu` VALUES (13, 10, '查看', '/role/detail', NULL, 3, 3, 1, 'RoleView', '', 'view', '2020-03-10 14:47:25', '2020-03-10 14:47:25', NULL);
INSERT INTO `sys_menu` VALUES (14, 10, '编辑', '/role/update', NULL, 4, 3, 1, 'RoleUpdate', '', 'update', '2020-03-10 14:47:26', '2020-03-10 14:47:26', NULL);
INSERT INTO `sys_menu` VALUES (15, 10, '分页api', '/role/list', NULL, 5, 3, 1, 'RoleList', '', 'list', '2020-03-10 14:47:27', '2020-03-10 14:47:27', NULL);
INSERT INTO `sys_menu` VALUES (16, 10, '分配角色菜单', '/role/setrole', NULL, 6, 3, 1, 'RoleSetrolemenu', '', 'setrolemenu', '2020-03-10 14:47:28', '2020-03-10 14:47:28', NULL);
INSERT INTO `sys_menu` VALUES (17, 2, '后台用户管理', '/admins', NULL, 40, 2, 1, 'Admins', 'user', 'none', '2020-03-10 14:47:29', '2020-03-10 14:47:29', NULL);
INSERT INTO `sys_menu` VALUES (18, 17, '新增', '/admins/create', NULL, 1, 3, 1, 'AdminsAdd', '', 'add', '2020-03-10 14:47:30', '2020-03-10 14:47:30', NULL);
INSERT INTO `sys_menu` VALUES (19, 17, '删除', '/admins/delete', NULL, 2, 3, 1, 'AdminsDel', '', 'del', '2020-03-10 14:47:31', '2020-03-10 14:47:31', NULL);
INSERT INTO `sys_menu` VALUES (20, 17, '查看', '/admins/detail', NULL, 3, 3, 1, 'AdminsView', '', 'view', '2020-03-10 14:47:32', '2020-03-10 14:47:32', NULL);
INSERT INTO `sys_menu` VALUES (21, 17, '编辑', '/admins/update', NULL, 4, 3, 1, 'AdminsUpdate', '', 'update', '2020-03-10 14:47:33', '2020-03-10 14:47:33', NULL);
INSERT INTO `sys_menu` VALUES (22, 17, '分页api', '/admins/list', NULL, 5, 3, 1, 'AdminsList', '', 'list', '2020-03-10 14:47:34', '2020-03-10 14:47:34', NULL);
INSERT INTO `sys_menu` VALUES (23, 17, '分配角色', '/admins/setrole', NULL, 6, 3, 1, 'AdminsSetrole', '', 'setadminrole', '2020-03-10 14:47:35', '2020-03-10 14:47:35', NULL);
INSERT INTO `sys_menu` VALUES (24, 2, 'test', '/test', NULL, 50, 2, 1, 'Test', 'list', 'none', '2020-03-10 14:47:36', '2020-03-10 14:47:36', NULL);
INSERT INTO `sys_menu` VALUES (25, 24, '新增', '/test/create', NULL, 1, 3, 1, 'TestAdd', '', 'add', '2020-03-10 14:47:37', '2020-03-10 14:47:37', NULL);
INSERT INTO `sys_menu` VALUES (26, 24, '删除', '/test/delete', NULL, 2, 3, 1, 'TestDel', '', 'del', '2020-03-10 14:47:38', '2020-03-10 14:47:38', NULL);
INSERT INTO `sys_menu` VALUES (27, 24, '查看', '/test/detail', NULL, 3, 3, 1, 'TestView', '', 'view', '2020-03-10 14:47:39', '2020-03-10 14:47:39', NULL);
INSERT INTO `sys_menu` VALUES (28, 24, '编辑', '/test/update', NULL, 4, 3, 1, 'TestUpdate', '', 'update', '2020-03-10 14:47:40', '2020-03-10 14:47:40', NULL);
INSERT INTO `sys_menu` VALUES (29, 24, '分页api', '/test/list', NULL, 5, 3, 1, 'TestList', '', 'list', '2020-03-10 14:47:41', '2020-03-10 14:47:41', NULL);
COMMIT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, 0, 'TOP', NULL, 1, 1, '2020-03-10 14:47:13', '2020-03-10 14:47:18', NULL);
INSERT INTO `sys_role` VALUES (2, 1, '超级管理员', NULL, 2, 1, '2020-03-10 14:47:21', '2020-03-10 14:47:24', NULL);
COMMIT;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` VALUES (1, 2, 1, '2020-03-10 14:47:13', '2020-03-10 14:47:13');
INSERT INTO `sys_role_menu` VALUES (2, 2, 2, '2020-03-10 14:47:14', '2020-03-10 14:47:14');
INSERT INTO `sys_role_menu` VALUES (3, 2, 3, '2020-03-10 14:47:15', '2020-03-10 14:47:15');
INSERT INTO `sys_role_menu` VALUES (4, 2, 4, '2020-03-10 14:47:16', '2020-03-10 14:47:16');
INSERT INTO `sys_role_menu` VALUES (5, 2, 5, '2020-03-10 14:47:17', '2020-03-10 14:47:17');
INSERT INTO `sys_role_menu` VALUES (6, 2, 6, '2020-03-10 14:47:18', '2020-03-10 14:47:18');
INSERT INTO `sys_role_menu` VALUES (7, 2, 7, '2020-03-10 14:47:19', '2020-03-10 14:47:19');
INSERT INTO `sys_role_menu` VALUES (8, 2, 8, '2020-03-10 14:47:20', '2020-03-10 14:47:20');
INSERT INTO `sys_role_menu` VALUES (9, 2, 9, '2020-03-10 14:47:21', '2020-03-10 14:47:21');
INSERT INTO `sys_role_menu` VALUES (10, 2, 10, '2020-03-10 14:47:22', '2020-03-10 14:47:22');
INSERT INTO `sys_role_menu` VALUES (11, 2, 11, '2020-03-10 14:47:23', '2020-03-10 14:47:23');
INSERT INTO `sys_role_menu` VALUES (12, 2, 12, '2020-03-10 14:47:24', '2020-03-10 14:47:24');
INSERT INTO `sys_role_menu` VALUES (13, 2, 13, '2020-03-10 14:47:25', '2020-03-10 14:47:25');
INSERT INTO `sys_role_menu` VALUES (14, 2, 14, '2020-03-10 14:47:26', '2020-03-10 14:47:26');
INSERT INTO `sys_role_menu` VALUES (15, 2, 15, '2020-03-10 14:47:27', '2020-03-10 14:47:27');
INSERT INTO `sys_role_menu` VALUES (16, 2, 16, '2020-03-10 14:47:28', '2020-03-10 14:47:28');
INSERT INTO `sys_role_menu` VALUES (17, 2, 17, '2020-03-10 14:47:29', '2020-03-10 14:47:29');
INSERT INTO `sys_role_menu` VALUES (18, 2, 18, '2020-03-10 14:47:30', '2020-03-10 14:47:30');
INSERT INTO `sys_role_menu` VALUES (19, 2, 19, '2020-03-10 14:47:31', '2020-03-10 14:47:31');
INSERT INTO `sys_role_menu` VALUES (20, 2, 20, '2020-03-10 14:47:32', '2020-03-10 14:47:32');
INSERT INTO `sys_role_menu` VALUES (21, 2, 21, '2020-03-10 14:47:33', '2020-03-10 14:47:33');
INSERT INTO `sys_role_menu` VALUES (22, 2, 22, '2020-03-10 14:47:34', '2020-03-10 14:47:34');
INSERT INTO `sys_role_menu` VALUES (23, 2, 23, '2020-03-10 14:47:35', '2020-03-10 14:47:35');
INSERT INTO `sys_role_menu` VALUES (24, 2, 24, '2020-03-10 14:47:36', '2020-03-10 14:47:36');
INSERT INTO `sys_role_menu` VALUES (25, 2, 25, '2020-03-10 14:47:37', '2020-03-10 14:47:37');
INSERT INTO `sys_role_menu` VALUES (26, 2, 26, '2020-03-10 14:47:38', '2020-03-10 14:47:38');
INSERT INTO `sys_role_menu` VALUES (27, 2, 27, '2020-03-10 14:47:39', '2020-03-10 14:47:39');
INSERT INTO `sys_role_menu` VALUES (28, 2, 28, '2020-03-10 14:47:40', '2020-03-10 14:47:40');
INSERT INTO `sys_role_menu` VALUES (29, 2, 29, '2020-03-10 14:47:41', '2020-03-10 14:47:41');
COMMIT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'ce0d6685-c15f-4126-a5b4-890bc9d2356d', '超级管理员', NULL, NULL, 1, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES (1, 1, 1, '2020-03-10 15:09:29', '2020-03-10 15:09:33');
INSERT INTO `sys_user_role` VALUES (2, 1, 2, '2020-03-10 15:09:36', '2020-03-10 15:09:39');
COMMIT;