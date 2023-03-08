package orm

import (
	"cilicili/config"
)

func GetComments(videoId uint, page int) (comments []*Comment, subComments map[uint][]*Comment, ok bool) {
	var commentIds []uint
	var sbs []*Comment
	subComments = make(map[uint][]*Comment)
	result := DB.Where("video_id = ? AND reply_id = 0", videoId).Limit(config.Conf.PageSize).Offset(config.Conf.PageSize * (page - 1)).Find(&comments)
	if result.Error != nil {
		return
	}
	for _, val := range comments {
		commentIds = append(commentIds, val.ID)
	}
	result = DB.Where("reply_id in (?)", commentIds).Find(&sbs)
	for _, val := range sbs {
		if subComments[val.ReplyId] == nil {
			subComments[val.ReplyId] = []*Comment{val}
		} else {
			subComments[val.ReplyId] = append(subComments[val.ReplyId], val)
		}
	}
	return comments, subComments, result.Error == nil
}
