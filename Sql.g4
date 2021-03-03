grammar Sql;

Id:[a-zA-Z_-]+;
Int:[0-9]+;
NotNull: [Nn][Oo][Tt] ' '* [Nn][Uu][Ll][Ll];
WS:[\r\n \t]+ ->skip;
statement
    : '"' STRING '"'
    |'\'' STRING '\''
    ;
tableName :'`' Id '`';
defaultStatement: 'DEFAULT' statement;

comment: 'COMMENT' STRING?;

FieldType
    :'char'
    |'varchar'
    |'int'
    |'bigint'
    |'tinyint'
    |'datetime'
    ;

length: '(' Int ')';

field
    :Id
    | '`' Id '`'
    ;

STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   | '\'' (ESC | SAFECODEPOINT)* '\''
   ;


fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
   : [0-9a-fA-F]
   ;
fragment SAFECODEPOINT
   : [\u0000-\u001F]
   ;
fragment INT
   : [0-9]*
   ;

fragment EXP
   : [Ee] [+\-]? INT
   ;

fieldName: '`' Id '`';

fieldStatement: field fileType=Id length? NotNull? defaultStatement? comment?;
table: 'CREATE' 'TABLE' tableName '('
    fieldStatement (',' fieldStatement)*
')';


//CREATE TABLE `op_bike` (
  //  `id` int(11) NOT NULL AUTO_INCREMENT,
  //  `name` varchar(10) NOT NULL COMMENT '车辆名称',
  //  `bike_version` tinyint(4) NOT NULL DEFAULT '1' COMMENT '车型号:1一代车,2二代车,3三代车...',
  //  `bike_model` varchar(63) NOT NULL DEFAULT '' COMMENT '车型',
  //  `factory_id` bigint(20) DEFAULT '0' COMMENT '厂商id',
  //  `batch_sn` char(4) DEFAULT '' COMMENT '批次号',
  //  `bt_token` varchar(127) DEFAULT '' COMMENT 'app连接车上蓝牙的密钥',
  //  `bt_mac` varchar(31) DEFAULT '' COMMENT '车上蓝牙的mac地址',
  //  `product_time` datetime NOT NULL COMMENT '车辆出厂时间',
  //  `stock_time` datetime NOT NULL COMMENT '入库时间',
  //  `last_launch_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后一次投放市场时间',
  //  `launch_city_id` int(11) NOT NULL DEFAULT '0' COMMENT '投放城市：0 表示无',
  //  `last_recycle_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后一次回收的时间',
  //  `last_recycle_reason` varchar(511) NOT NULL DEFAULT '' COMMENT '回收原因',
  //  `return_depot_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '返厂时间',
  //  `return_depot_reason` varchar(511) NOT NULL DEFAULT '' COMMENT '返厂原因',
  //  `scrap_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '报废时间',
  //  `scrap_reason` varchar(511) NOT NULL DEFAULT '' COMMENT '报废原因',
  //  `status` int(11) NOT NULL DEFAULT '0' COMMENT '车辆状态：1 入库，2 投放 ，3 回收， 4，报废， 5返厂（废弃），7 已返厂（废弃）， 8 资产被盗， 9资产丢失 ，10调拨入库',
  //  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '绑定的sim卡手机号',
  //  `create_time` datetime NOT NULL,
  //  `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  //  PRIMARY KEY (`id`),
  //  UNIQUE KEY `idx_sn` (`sn`),
  //  KEY `idx_lcid` (`launch_city_id`),
  //  KEY `idx_create_time` (`create_time`)
  //) ENGINE=InnoDB AUTO_INCREMENT=27807 DEFAULT CHARSET=utf8mb4;