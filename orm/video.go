package orm

import (
	"gorm.io/gorm"
	"strconv"
)

func CreateVideo(userId uint, videoType, videoUrl, videoName string) (v *Video, ok bool, msg string) {
	v = &Video{UserId: userId, Type: videoType, VideoUrl: videoUrl, VideoName: videoName}
	result := DB.Create(v)
	if result.Error != nil {
		return nil, false, "username has been used"
	}
	return v, true, "ok"
}

func CreateComment(videoId, userId uint, replyId, text string) (c *Comment, ok bool, msg string) {
	if replyId == "" {
		c = &Comment{VideoId: videoId, UserId: userId, Text: text}
	} else {
		replyId_, _ := strconv.Atoi(replyId)
		c = &Comment{VideoId: videoId, UserId: userId, ReplyId: uint(replyId_), Text: text}
	}
	result := DB.Create(c)
	if result.Error != nil {
		return nil, false, "db error"
	}
	return c, true, "ok"
}

func CreateGood(videoId, userId uint) (g *Good, ok bool, msg string) {
	g = &Good{VideoId: videoId, UserId: userId}
	result := DB.Create(g)
	if result.Error != nil {
		return nil, false, "db error"
	}
	return g, true, "ok"
}

func CreateCollect(videoId, userId uint) (c *Collect, ok bool, msg string) {
	c = &Collect{VideoId: videoId, UserId: userId}
	result := DB.Create(c)
	if result.Error != nil {
		return nil, false, "db error"
	}
	return c, true, "ok"
}

func CreateBulletChat(videoId, userId uint, text, videoTime string) (b *BulletChat, ok bool, msg string) {
	b = &BulletChat{VideoId: videoId, VideoTime: videoTime, UserId: userId, Text: text}
	result := DB.Create(b)
	if result.Error != nil {
		return nil, false, "db error"
	}
	return b, true, "ok"
}

func GetVideos(conditions map[string]any) (videos []*Video, ok bool) {
	d := DB
	for k, v := range conditions {
		if k == "video_name" {
			vv, _ := v.(string)
			d = d.Where(k+" like ?", "%"+vv+"%")
		} else {

			d = d.Where(k+" = ?", v)
		}
	}
	result := d.Find(&videos)
	return videos, result.Error == nil
}

func DelVideo(videoId uint) bool {
	result := DB.Where("id = ?", videoId).Delete(&Video{Model: gorm.Model{ID: videoId}})
	return result.Error == nil
}

func DelComment(commentId uint) bool {
	result := DB.Where("id = ?", commentId).Delete(&Comment{})
	return result.Error == nil
}
