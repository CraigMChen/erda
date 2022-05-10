ALTER table dice_release
    ADD COLUMN `opus_id` VARCHAR(36) NOT NULL DEFAULT '' COMMENT 'erda_gallery_opus.id',
    ADD COLUMN `opus_version_id` VARCHAR(36) NOT NULL DEFAULT '' COMMENT 'erda_gallery_opus_version.id';
