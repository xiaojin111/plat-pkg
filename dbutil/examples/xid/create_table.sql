DROP TABLE IF EXISTS `event`;
CREATE TABLE `event` (
  `event_id` BINARY(20) NOT NULL COMMENT 'Event ID',
  `type` VARCHAR(255) NOT NULL COMMENT 'Type',
  `ref_id` BINARY(20) NOT NULL COMMENT 'Ref ID',
  `created_at` TIMESTAMP NOT NULL COMMENT '数据记录创建时间',
  `updated_at` TIMESTAMP NOT NULL COMMENT '数据记录更新时间',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '数据记录伪删除时间',
  PRIMARY KEY (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Event';
